// snmp/snmp.go
package snmpinfo

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"golang/snmp/config"
	//"golang/snmp/model"
	"log"
	"sync"
)

// Fetch SNMP system data
func FetchSNMPSystemData(g *gosnmp.GoSNMP) (map[string]string, error) {
	resultChan := make(chan map[string]string)
	var snmpData = make(map[string]string)
	var wg sync.WaitGroup

	for oid := range config.SNMPOids {
		wg.Add(1)
		go fetchSNMPData(oid, g, resultChan, &wg)
	}

	for range config.SNMPOids {
		data := <-resultChan
		for key, value := range data {
			snmpData[key] = value
		}
	}

	wg.Wait()
	return snmpData, nil
}

func fetchSNMPData(oid string, g *gosnmp.GoSNMP, resultChan chan<- map[string]string, wg *sync.WaitGroup) {
	result, err := g.Get([]string{oid})
	if err != nil {
		log.Printf("Error fetching SNMP data for OID %s: %s", oid, err)
		resultChan <- map[string]string{"error": fmt.Sprintf("Error fetching OID %s: %s", oid, err)}
		wg.Done()
		return
	}

	for _, variable := range result.Variables {
		var value string
		switch v := variable.Value.(type) {
		case []byte:
			value = string(v) // Convert byte array to string
		default:
			value = fmt.Sprintf("%v", v)
		}
		resultChan <- map[string]string{config.SNMPOids[oid]: value}
	}
	wg.Done()
}
