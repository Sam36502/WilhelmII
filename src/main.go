package main

import (
	"fmt"
)

const (
	OPTIONS_FILE = "src/options.properties"
)

func init() {
	LoadOptions(OPTIONS_FILE)
}

func main() {
	game := LoadGame("src/testgame.json")
	fmt.Println("Game loaded.")
	fmt.Println("  Num. Rooms: " + fmt.Sprint(len(game.Rooms)))
	for k, v := range game.Items {
		fmt.Println("[" + k + "] - " + v.Name + "; \"" + v.Description[0] + " ...\"")
	}
}
