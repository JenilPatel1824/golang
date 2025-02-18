package main

type SNMP struct {
	Name        string `json:"system.name"`
	Description string `json:"system.description"`
	Location    string `json:"system.location"`
	ObjectID    string `json:"system.objectId"`
	Uptime      string `json:"system.uptime"`
	Interfaces  string `json:"system.interfaces"` // This can be a string or an array depending on the data
}

var Hello string = "123"
