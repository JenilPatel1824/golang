package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gosnmp/gosnmp"
)

// Struct to store SNMP system information

var oids = map[string]string{
	"1.3.6.1.2.1.1.5.0": "system.name",        // sysName
	"1.3.6.1.2.1.1.1.0": "system.description", // sysDescr
	"1.3.6.1.2.1.1.6.0": "system.location",    // sysLocation
	"1.3.6.1.2.1.1.2.0": "system.objectId",    // sysObjectID
	"1.3.6.1.2.1.1.3.0": "system.uptime",      // sysUpTime
	"1.3.6.1.2.1.2.1.0": "system.interfaces",  // ifNumber
}

type SNMPSystemInfo struct {
	SystemName        string `json:"system.name"`
	SystemDescription string `json:"system.description"`
	SystemLocation    string `json:"system.location"`
	SystemObjectID    string `json:"system.objectId"`
	SystemUptime      string `json:"system.uptime"`
	SystemInterfaces  string `json:"system.interfaces"`
}

func fetchSNMPData(oid string, g *gosnmp.GoSNMP, resultChan chan<- map[string]string, wg *sync.WaitGroup) {
	// Perform SNMP GET request for a single OID
	result, err := g.Get([]string{oid})
	if err != nil {
		log.Printf("Error fetching SNMP data for OID %s: %s", oid, err)
		resultChan <- map[string]string{"error": fmt.Sprintf("Error fetching OID %s: %s", oid, err)}
		wg.Done()
		return
	}

	// Send the result to the channel
	for _, variable := range result.Variables {
		var value string
		// Check if the value is a byte array and convert it to a string
		switch v := variable.Value.(type) {
		case []byte:
			value = string(v) // Convert byte array to string
		default:
			value = fmt.Sprintf("%v", v)
		}
		// Map the value to the corresponding key
		resultChan <- map[string]string{oids[oid]: value}
	}
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	startTime := time.Now()
	fmt.Println("Execution Started at:", startTime.Format(time.RFC3339))

	// Initialize the GoSNMP object
	g := &gosnmp.GoSNMP{
		Target:    "172.16.14.6", // SNMP target device IP
		Community: "public",      // SNMP community string
		Port:      161,           // SNMP port (default is 161)
		Version:   gosnmp.Version2c,
		Timeout:   2 * time.Second, // SNMP version 2c
	}

	// Connect to the SNMP target
	err := g.Connect()
	if err != nil {
		log.Fatalf("Error connecting to SNMP target: %s", err)
	}
	defer g.Conn.Close()

	// Create a channel to collect results
	resultChan := make(chan map[string]string)

	// OIDs to fetch, mapped to their corresponding keys

	// Launch Goroutines to fetch SNMP data concurrently
	for oid := range oids {
		wg.Add(1)
		go fetchSNMPData(oid, g, resultChan, &wg)
	}

	// Store the SNMP data in a map
	snmpData := make(map[string]string)
	for range oids {
		data := <-resultChan
		for key, value := range data {
			snmpData[key] = value
		}
	}

	// Print the collected SNMP data
	for key, value := range snmpData {
		fmt.Printf("%s: %s\n", key, value)
	}

	wg.Wait()

	jsonData, err := json.MarshalIndent(snmpData, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling map:", err)
		return
	}

	//fmt.Println("json data: ", jsonData)

	// Convert JSON string to struct
	var snmpcore SNMPSystemInfo
	err = json.Unmarshal(jsonData, &snmpcore)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Print result
	fmt.Printf("%+v\n", snmpcore)

	// Record end time
	endTime := time.Now()
	fmt.Println("Execution Ended at:", endTime.Format(time.RFC3339))

	// Calculate total execution time
	duration := endTime.Sub(startTime)
	fmt.Printf("Total Execution Time: %v\n", duration)
}
