package main

type Item struct {
	Name        string   `json:"name"`
	Aliases     []string `json:"aliases"`
	Description []string `json:"description"`
}
