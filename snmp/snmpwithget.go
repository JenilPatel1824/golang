package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
)

func main() {
	// Record start time
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

	// Perform SNMP GET request for all the required OIDs
	oids := []string{
		"1.3.6.1.2.1.1.5.0", // sysName
		"1.3.6.1.2.1.1.1.0", // sysDescr
		"1.3.6.1.2.1.1.6.0", // sysLocation
		"1.3.6.1.2.1.1.2.0", // sysObjectID
		"1.3.6.1.2.1.1.3.0", // sysUpTime
		"1.3.6.1.2.1.2.1.0", // ifNumber
	}

	result, err := g.Get(oids)
	if err != nil {
		log.Fatalf("Error fetching SNMP data: %s", err)
	}

	// Print the result
	for _, variable := range result.Variables {
		fmt.Printf("OID: %s, Value: %v\n", variable.Name, variable.Value)
	}

	// Record end time
	endTime := time.Now()
	fmt.Println("Execution Ended at:", endTime.Format(time.RFC3339))

	// Calculate total execution time
	duration := endTime.Sub(startTime)
	fmt.Printf("Total Execution Time: %v\n", duration)
}
