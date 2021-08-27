package main

type Door struct {
	Locked      bool                `json:"locked"`
	Names       []string            `json:"names"`
	Description []string            `json:"description"`
	Openings    map[string][]string `json:"openings"`
}

var DoorIndex = make([]Door, 10, 10)
var DoorNameIndex = make(map[string]int)

func GetDoor(name string) Door {
	return DoorIndex[DoorNameIndex[name]]
}
