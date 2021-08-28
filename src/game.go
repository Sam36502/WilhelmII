package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Coords struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Game struct {
	GameName    string `json:"game_name"`
	GameCreator string `json:"game_creator"`
	GameVersion string `json:"game_version"`

	EngineVersion string `json:"engine_version"`
	StartCoords   Coords `json:"start_coords"`
	EndCoords     Coords `json:"end_coords"`

	Rooms []Room `json:"rooms"`
	Items []Item `json:"items"`
	Doors []Door `json:"doors"`

	roomIndex     map[Coords]Room
	itemIndex     []Item
	itemNameIndex map[string]int
	doorIndex     []Door
	doorNameIndex map[string]int
}

func (game *Game) GetRoom(coords Coords) Room {
	return game.roomIndex[coords]
}

func (game *Game) GetItem(name string) Item {
	return game.itemIndex[game.itemNameIndex[name]]
}

func (game *Game) GetDoor(name string) Door {
	return game.doorIndex[game.doorNameIndex[name]]
}

func LoadGame(gameFile string) Game {
	// Retrieve File Contents
	LogMsg("Loading game file...", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Loading game file...")
	}

	bytes, err := ioutil.ReadFile(gameFile)
	if err != nil {
		LogMsg(
			"Failed to load game file '"+gameFile+"'. Please check it is correctly formatted.",
			LOG_FATAL,
		)
	}

	// Parse JSON
	LogMsg("Parsing game file...", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Parsing game file...")
	}
	var parsedGame Game
	err = json.Unmarshal(bytes, &parsedGame)
	if err != nil {
		LogMsg(
			"Failed to parse game file '"+gameFile+"':\n"+err.Error(),
			LOG_FATAL,
		)
	}

	// Index all Names
	LogMsg("Indexing object names...", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexing object names...")
	}

	// Index Rooms
	parsedGame.roomIndex = make(map[Coords]Room)
	for _, v := range parsedGame.Rooms {
		parsedGame.roomIndex[v.Coords] = v
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Rooms))+" rooms.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Rooms)) + " rooms.")
	}

	// Index Items
	parsedGame.itemIndex = make([]Item, 0)
	parsedGame.itemNameIndex = make(map[string]int)
	for k, v := range parsedGame.Items {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := parsedGame.itemNameIndex[n]
			if exists {
				LogMsg(
					"Failed to load Game file: Item name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			parsedGame.itemNameIndex[n] = k
		}
		parsedGame.itemIndex = append(parsedGame.itemIndex, v)
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Items))+" items.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Items)) + " items.")
	}

	// Index Doors
	parsedGame.doorIndex = make([]Door, 0)
	parsedGame.doorNameIndex = make(map[string]int)
	for k, v := range parsedGame.Doors {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := parsedGame.doorNameIndex[n]
			if exists {
				LogMsg(
					"Failed to load Game file: Door name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			parsedGame.doorNameIndex[n] = k
		}
		parsedGame.doorIndex = append(parsedGame.doorIndex, v)
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Doors))+" doors.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Doors)) + " doors.")
	}

	LogMsg("Game loaded successfully!", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Game loaded successfully!")
	}

	return parsedGame
}
