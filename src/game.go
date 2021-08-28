package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	DIR_NORTH = Direction(0)
	DIR_EAST  = Direction(1)
	DIR_SOUTH = Direction(2)
	DOR_WEST  = Direction(3)
)

type Coords struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Direction int

type Game struct {
	GameName    string `json:"game_name"`
	GameCreator string `json:"game_creator"`
	GameVersion string `json:"game_version"`

	EngineVersion string `json:"engine_version"`
	StartCoords   Coords `json:"start_coords"`
	EndCoords     Coords `json:"end_coords"`

	Rooms   []Room   `json:"rooms"`
	Items   []Item   `json:"items"`
	Doors   []Door   `json:"doors"`
	Endings []Ending `json:"endings"`

	roomIndex   map[Coords]*Room
	itemIndex   map[string]*Item
	doorIndex   map[string]*Door
	endingIndex map[string]*Ending
}

func (game *Game) GetRoom(coords Coords) *Room {
	return game.roomIndex[coords]
}

func (game *Game) GetItem(name string) *Item {
	return game.itemIndex[name]
}

func (game *Game) GetDoor(name string) *Door {
	return game.doorIndex[name]
}

func (game *Game) GetEnding(name string) *Ending {
	return game.endingIndex[name]
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
	parsedGame.roomIndex = make(map[Coords]*Room)
	for _, v := range parsedGame.Rooms {
		parsedGame.roomIndex[v.Coords] = &v
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Rooms))+" rooms.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Rooms)) + " rooms.")
	}

	// Index Items
	parsedGame.itemIndex = make(map[string]*Item)
	for _, v := range parsedGame.Items {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := parsedGame.itemIndex[n]
			if exists {
				LogMsg(
					"Failed to load Game file: Item name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			parsedGame.itemIndex[n] = &v
		}
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Items))+" items.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Items)) + " items.")
	}

	// Index Doors
	parsedGame.doorIndex = make(map[string]*Door)
	for _, v := range parsedGame.Doors {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := parsedGame.doorIndex[n]
			if exists {
				LogMsg(
					"Failed to load Game file: Door name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			parsedGame.doorIndex[n] = &v
		}
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Doors))+" doors.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Doors)) + " doors.")
	}

	// Index Endings
	parsedGame.endingIndex = make(map[string]*Ending)
	for _, v := range parsedGame.Endings {
		parsedGame.endingIndex[v.Name] = &v
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Endings))+" endings.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Endings)) + " endings.")
	}

	LogMsg("Game loaded successfully!", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Game loaded successfully!")
	}

	return parsedGame
}
