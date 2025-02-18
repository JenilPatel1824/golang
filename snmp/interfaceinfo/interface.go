package interfaces

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"golang/snmp/config"
	"golang/snmp/model"
	"log"
)

// GetInterfaces fetches SNMP interface data from a device
func GetInterfaces(g *gosnmp.GoSNMP, numInterfaces int) ([]model.SNMPInterfaceInfo, error) {
	var interfacesData []model.SNMPInterfaceInfo

	// Loop over interface indexes
	for i := 1; i <= numInterfaces; i++ {
		interfaceInfo := model.SNMPInterfaceInfo{
			Index: fmt.Sprintf("%d", i),
		}

		// Fetch data for each interface using the corresponding OIDs
		for oid, field := range config.InterfaceOids {
			// Fetch the value for the current OID
			result, err := g.Get([]string{oid + fmt.Sprintf(".%d", i)})
			if err != nil {
				log.Printf("Error fetching data for OID %s.%d: %s", oid, i, err)
				continue
			}

			// Check if the result is valid and not nil before type assertion
			if len(result.Variables) > 0 {
				if result.Variables[0].Value != nil {
					// Assign the result to the corresponding field in the interfaceInfo struct
					switch field {
					case "Index":
						interfaceInfo.Index = fmt.Sprintf("%v", result.Variables[0].Value)
					case "Name":
						interfaceInfo.Name = fmt.Sprintf("%s", result.Variables[0].Value)
					case "Alias":
						interfaceInfo.Alias = fmt.Sprintf("%s", result.Variables[0].Value)
					case "OperationalStatus":
						interfaceInfo.OperationalStatus = fmt.Sprintf("%v", result.Variables[0].Value)
					case "AdminStatus":
						interfaceInfo.AdminStatus = fmt.Sprintf("%v", result.Variables[0].Value)
					case "Description":
						interfaceInfo.Description = fmt.Sprintf("%s", result.Variables[0].Value)
					case "SentErrorPackets":
						interfaceInfo.SentErrorPackets = fmt.Sprintf("%v", result.Variables[0].Value)
					case "ReceivedErrorPackets":
						interfaceInfo.ReceivedErrorPackets = fmt.Sprintf("%v", result.Variables[0].Value)
					case "SentOctets":
						interfaceInfo.SentOctets = fmt.Sprintf("%v", result.Variables[0].Value)
					case "ReceivedOctets":
						interfaceInfo.ReceivedOctets = fmt.Sprintf("%v", result.Variables[0].Value)
					case "Speed":
						interfaceInfo.Speed = fmt.Sprintf("%v", result.Variables[0].Value)
					case "PhysicalAddress":
						interfaceInfo.PhysicalAddress = fmt.Sprintf("%v", result.Variables[0].Value)
					case "DiscardPackets":
						interfaceInfo.DiscardPackets = fmt.Sprintf("%v", result.Variables[0].Value)
					case "InPackets":
						interfaceInfo.InPackets = fmt.Sprintf("%s", result.Variables[0].Value)
					case "OutPackets":
						interfaceInfo.OutPackets = fmt.Sprintf("%s", result.Variables[0].Value)
					}
				}
			}
		}

		// Add the interface info to the result slice
		interfacesData = append(interfacesData, interfaceInfo)
	}

	return interfacesData, nil
}
