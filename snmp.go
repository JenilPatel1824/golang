package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
)

var oids = map[string]string{
	"system.name":        "1.3.6.1.2.1.1.5.0",
	"system.description": "1.3.6.1.2.1.1.1.0",
	"system.location":    "1.3.6.1.2.1.1.6.0",
	"system.objectId":    "1.3.6.1.2.1.1.2.0",
	"system.uptime":      "1.3.6.1.2.1.1.3.0",
}

func main() {
	target := "172.16.14.6" // Replace with your SNMP device IP
	community := "public"
	port := 161

	snmp := &gosnmp.GoSNMP{
		Target:    target,
		Port:      uint16(port),
		Community: community,
		Version:   gosnmp.Version2c,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   2,
	}

	err := snmp.Connect()
	if err != nil {
		log.Fatalf("SNMP connection failed: %v", err)
	}
	defer snmp.Conn.Close()

	snmpData := make(map[string]string)
	for key, oid := range oids {
		result, err := snmp.Get([]string{oid})
		if err != nil {
			log.Printf("Failed to fetch %s: %v", key, err)
			continue
		}

		for _, variable := range result.Variables {
			switch v := variable.Value.(type) {
			case []byte:
				snmpData[key] = string(v) // Convert byte slice to string
			case uint32: // Handling Timeticks (System Uptime)
				if key == "system.uptime" {
					snmpData[key] = parseTimeticks(v)
				} else {
					snmpData[key] = fmt.Sprintf("%v", v)
				}
			default:
				snmpData[key] = fmt.Sprintf("%v", v)
			}
		}
	}

	interfaceData := make(map[string]map[string]string)
	baseOID := "1.3.6.1.2.1.2.2.1"
	interfaceFields := map[string]string{
		"interface.index":                 "1",
		"interface.name":                  "2",
		"interface.operational.status":    "8",
		"interface.admin.status":          "7",
		"interface.description":           "2",
		"interface.sent.error.packet":     "20",
		"interface.received.error.packet": "14",
		"interface.sent.octets":           "16",
		"interface.received.octets":       "10",
		"interface.speed":                 "5",
		"interface.physical.address":      "6",
		"interface.in.packets":            "11",
		"interface.out.packets":           "17",
	}

	for field, index := range interfaceFields {
		result, err := snmp.BulkWalkAll(baseOID + "." + index)
		if err != nil {
			log.Printf("Failed to fetch %s: %v", field, err)
			continue
		}

		for _, variable := range result {
			oidParts := variable.Name[len(baseOID)+1:] // Extract interface index
			ifaceIndex := oidParts

			if _, exists := interfaceData[ifaceIndex]; !exists {
				interfaceData[ifaceIndex] = make(map[string]string)
			}

			// Convert byte slices to strings
			switch v := variable.Value.(type) {
			case []byte:
				interfaceData[ifaceIndex][field] = string(v)
			default:
				interfaceData[ifaceIndex][field] = fmt.Sprintf("%v", v)
			}
		}
	}

	finalData := map[string]interface{}{
		"snmp":           snmpData,
		"snmp.interface": interfaceData,
	}

	jsonData, err := json.MarshalIndent(finalData, "", "    ")
	if err != nil {
		log.Fatalf("JSON encoding failed: %v", err)
	}

	fmt.Println(string(jsonData))
}

// Function to convert SNMP Timeticks to human-readable format
func parseTimeticks(ticks uint32) string {
	seconds := int(ticks) / 100
	days := seconds / 86400
	hours := (seconds % 86400) / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	return fmt.Sprintf("%d days, %02d:%02d:%02d", days, hours, minutes, secs)
}
