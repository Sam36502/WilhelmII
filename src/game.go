package main

import (
	"encoding/json"
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

	Rooms map[Coords]Room `json:"rooms"`
	Items map[string]Item `json:"items"`
	Doors map[string]Door `json:"doors"`
}

func LoadGame(gameFile string) Game {
	bytes, err := ioutil.ReadFile(gameFile)
	if err != nil {
		ErrorMsg(
			"Failed to load game file '"+gameFile+"'. Please check it is correctly formatted.",
			LOG_FATAL,
		)
	}

	var parsedGame Game
	json.Unmarshal(bytes, &parsedGame)
	return parsedGame
}
