package main

type Room struct {
	Coords      `json:"coords"`
	DoorNames   []string `json:"doors"`
	ItemNames   []string `json:"items"`
	Description []string `json:"description"`

	Items []*Item
	Doors []*Door
}
