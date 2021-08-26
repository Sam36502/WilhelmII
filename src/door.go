package main

type Door struct {
	Name        string              `json:"name"`
	Locked      bool                `json:"locked"`
	Aliases     []string            `json:"aliases"`
	Description []string            `json:"description"`
	Openings    map[string][]string `json:"openings"`
}
