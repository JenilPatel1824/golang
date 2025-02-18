// models/models.go
package model

// Struct for SNMP system information
type SNMPSystemInfo struct {
	SystemName        string `json:"system.name"`
	SystemDescription string `json:"system.description"`
	SystemLocation    string `json:"system.location"`
	SystemObjectID    string `json:"system.objectId"`
	SystemUptime      string `json:"system.uptime"`
	SystemInterfaces  string `json:"system.interfaces"`
}

// Struct for SNMP interface information
type SNMPInterfaceInfo struct {
	Index                string `json:"interface.index"`
	Name                 string `json:"interface.name"`
	Alias                string `json:"interface.alias"`
	OperationalStatus    string `json:"interface.operational.status"`
	AdminStatus          string `json:"interface.admin.status"`
	Description          string `json:"interface.description"`
	SentErrorPackets     string `json:"interface.sent.error.packet"`
	ReceivedErrorPackets string `json:"interface.received.error.packet"`
	SentOctets           string `json:"interface.sent.octets"`
	ReceivedOctets       string `json:"interface.received.octets"`
	Speed                string `json:"interface.speed"`
	PhysicalAddress      string `json:"interface.physical.address"`
	DiscardPackets       string `json:"interface.discard.packets"`
	InPackets            string `json:"interface.in.packets"`
	OutPackets           string `json:"interface.out.packets"`
}
