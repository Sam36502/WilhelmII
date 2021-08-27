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
	for _, v := range parsedGame.Rooms {
		RoomIndex[v.Coords] = v
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Rooms))+" rooms.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Rooms)) + " rooms.")
	}

	// Index Items
	for k, v := range parsedGame.Items {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := ItemNameIndex[n]
			if exists {
				LogMsg(
					"Failed to load Game file: Item name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			ItemNameIndex[n] = k
		}
		ItemIndex[k] = v
	}
	LogMsg("Indexed "+fmt.Sprint(len(parsedGame.Items))+" items.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("Indexed " + fmt.Sprint(len(parsedGame.Items)) + " items.")
	}

	// Index Doors
	for k, v := range parsedGame.Doors {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := DoorNameIndex[n]
			if exists {
				LogMsg(
					"Failed to load Game file: Door name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			DoorNameIndex[n] = k
		}
		DoorIndex[k] = v
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
