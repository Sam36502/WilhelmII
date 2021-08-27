package main

import (
	"fmt"
)

const (
	OPTIONS_FILE = "options.properties"
)

func init() {
	LoadOptions(OPTIONS_FILE)
}

func main() {
	game := LoadGame("testgame.json")
	fmt.Println("Game loaded.")
	fmt.Println("  Num. Rooms: " + fmt.Sprint(len(game.Rooms)))
	for k, v := range game.Items {
		fmt.Println("[" + fmt.Sprint(k) + "] - " + v.Names[0] + "; \"" + v.Description[0] + " ...\"")
	}
}
