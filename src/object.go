package main

type Room struct {
	Coords      `json:"coords"`
	DoorNames   []string `json:"doors"`
	ItemNames   []string `json:"items"`
	Description []string `json:"description"`

	Items []*Item
	Doors []*Door
}

type Item struct {
	Names       []string `json:"names"`
	Description []string `json:"description"`
}

type Door struct {
	Locked      bool                `json:"locked"`
	Names       []string            `json:"names"`
	Description []string            `json:"description"`
	Openings    map[string][]string `json:"openings"`
}
