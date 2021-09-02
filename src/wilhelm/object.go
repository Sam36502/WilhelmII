package wilhelm

type Room struct {
	Coords      Coords   `json:"coords"`
	DoorNames   []string `json:"doors"`
	ItemNames   []string `json:"items"`
	Description []string `json:"description"`

	Items []*Item
	Doors []*Door
}

type Item struct {
	Names   []string `json:"names"`
	Context []string `json:"context"` // The text displayed when it's found
	Inspect []string `json:"inspect"` // The text displayed when inspected in the player's inventory
}

type Door struct {
	Locked      bool                `json:"locked"`
	Names       []string            `json:"names"`
	Description []string            `json:"description"`
	Openings    map[string][]string `json:"openings"`
}

type Ending struct {
	Name        string   `json:"name"`
	Description []string `json:"description"`
}
