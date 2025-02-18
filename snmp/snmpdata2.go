package main

type SNMPData struct {
	SNMP       `json:"snmp"`
	Interfaces []snmp.Interface `json:"snmp.interface"`
}
