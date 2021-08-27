package main

type Item struct {
	Names       []string `json:"names"`
	Description []string `json:"description"`
}

var ItemIndex = make([]Item, 10, 10)
var ItemNameIndex = make(map[string]int)

func GetItem(name string) Item {
	return ItemIndex[ItemNameIndex[name]]
}
