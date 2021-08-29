package wilhelm

type Room struct {
	coords      Coords   `json:"coords"`
	doorNames   []string `json:"doors"`
	itemNames   []string `json:"items"`
	description []string `json:"description"`

	items []*Item
	doors []*Door
}

func (r *Room) GetDescription() []string { return r.description }

type Item struct {
	names   []string `json:"names"`
	context []string `json:"context"` // The text displayed when it's found
	inspect []string `json:"inspect"` // The text displayed when inspected in the player's inventory
}

type Door struct {
	locked      bool                `json:"locked"`
	names       []string            `json:"names"`
	description []string            `json:"description"`
	openings    map[string][]string `json:"openings"`
}

type Ending struct {
	name        string   `json:"name"`
	description []string `json:"description"`
}
