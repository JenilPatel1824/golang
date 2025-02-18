package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gosnmp/gosnmp"
	"golang/snmp/interfaceinfo"
	//"golang/snmp/model"
	"golang/snmp/snmpinfo"
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

	// Fetch SNMP system data
	systemData, err := snmpinfo.FetchSNMPSystemData(g)
	if err != nil {
		log.Fatalf("Error fetching SNMP system data: %s", err)
	}

	str := systemData["system.interfaces"]

	numberOfInterface, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the integer
	fmt.Println("Converted integer:", numberOfInterface)
	// Fetch SNMP interface data (specify the number of interfaces to fetch)
	interfaceData, err := interfaces.GetInterfaces(g, numberOfInterface)
	if err != nil {
		log.Fatalf("Error fetching SNMP interface data: %s", err)
	}

	// Prepare the result map with SNMP system data and interfaces data
	result := map[string]interface{}{
		"snmp":       systemData,
		"interfaces": interfaceData,
	}

	// Print the final result as formatted JSON
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling data to JSON: %s", err)
	}

	fmt.Println(string(jsonData))

	// Record end time
	endTime := time.Now()
	fmt.Println("Execution Ended at:", endTime.Format(time.RFC3339))

	// Calculate total execution time
	duration := endTime.Sub(startTime)
	fmt.Printf("Total Execution Time: %v\n", duration)
}
