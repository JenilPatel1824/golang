package main

type Interface struct {
	Index               string `json:"interface.index"`
	Name                string `json:"interface.name"`
	Alias               string `json:"interface.alias"`
	OperationalStatus   string `json:"interface.operational.status"`
	AdminStatus         string `json:"interface.admin.status"`
	Description         string `json:"interface.description"`
	SentErrorPacket     string `json:"interface.sent.error.packet"`
	ReceivedErrorPacket string `json:"interface.received.error.packet"`
	SentOctets          string `json:"interface.sent.octets"`
	ReceivedOctets      string `json:"interface.received.octets"`
	Speed               string `json:"interface.speed"`
	PhysicalAddress     string `json:"interface.physical.address"`
	DiscardPackets      string `json:"interface.discard.packets"`
	InPackets           string `json:"interface.in.packets"`
	OutPackets          string `json:"interface.out.packets"`
}
