package wilhelm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	DIR_NORTH = Direction(0)
	DIR_EAST  = Direction(1)
	DIR_SOUTH = Direction(2)
	DIR_WEST  = Direction(3)
)

func parseDirection(dirName string) (Direction, error) {
	dirName = strings.ToLower(dirName)

	switch dirName {
	case "north":
	case "n":
		return DIR_NORTH, nil
	case "east":
	case "e":
		return DIR_EAST, nil
	case "south":
	case "s":
		return DIR_SOUTH, nil
	case "west":
	case "w":
		return DIR_WEST, nil
	}

	return Direction(-1), InvalidDirectionError(dirName)
}

type Direction int

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

	Rooms   []Room   `json:"rooms"`
	Items   []Item   `json:"items"`
	Doors   []Door   `json:"doors"`
	Endings []Ending `json:"endings"`

	Player Player

	roomIndex   map[Coords]*Room
	itemIndex   map[string]*Item
	doorIndex   map[string]*Door
	endingIndex map[string]*Ending

	gameFinished bool
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
	LogMsg("  Loading game file...", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Loading game file...")
	}

	bytes, err := ioutil.ReadFile(gameFile)
	if err != nil {
		LogMsg(
			"  Failed to load game file '"+gameFile+"'. Please check it is correctly formatted.",
			LOG_FATAL,
		)
	}

	// Parse JSON
	LogMsg("  Parsing game file...", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Parsing game file...")
	}
	var parsedGame Game
	err = json.Unmarshal(bytes, &parsedGame)
	if err != nil {
		LogMsg(
			"  Failed to parse game file '"+gameFile+"':\n"+err.Error(),
			LOG_FATAL,
		)
	}

	// Index all Names
	LogMsg("  Indexing object names...", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Indexing object names...")
	}

	// Index Rooms
	parsedGame.roomIndex = make(map[Coords]*Room)
	for k, v := range parsedGame.Rooms {
		parsedGame.roomIndex[v.Coords] = &parsedGame.Rooms[k]
	}
	LogMsg("  Indexed "+fmt.Sprint(len(parsedGame.Rooms))+" rooms.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Indexed " + fmt.Sprint(len(parsedGame.Rooms)) + " rooms.")
	}

	// Index Items
	parsedGame.itemIndex = make(map[string]*Item)
	for k, v := range parsedGame.Items {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := parsedGame.itemIndex[n]
			if exists {
				LogMsg(
					"  Failed to load Game file: Item name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			parsedGame.itemIndex[n] = &parsedGame.Items[k]
		}
	}
	LogMsg("  Indexed "+fmt.Sprint(len(parsedGame.Items))+" items.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Indexed " + fmt.Sprint(len(parsedGame.Items)) + " items.")
	}

	// Index Doors
	parsedGame.doorIndex = make(map[string]*Door)
	for k, v := range parsedGame.Doors {
		for _, n := range v.Names {
			// Check for duplicates
			_, exists := parsedGame.doorIndex[n]
			if exists {
				LogMsg(
					"  Failed to load Game file: Door name '"+n+"' is used multiple times",
					LOG_FATAL,
				)
				return Game{}
			}

			parsedGame.doorIndex[n] = &parsedGame.Doors[k]
		}
	}
	LogMsg("  Indexed "+fmt.Sprint(len(parsedGame.Doors))+" doors.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Indexed " + fmt.Sprint(len(parsedGame.Doors)) + " doors.")
	}

	// Index Endings
	parsedGame.endingIndex = make(map[string]*Ending)
	for k, v := range parsedGame.Endings {
		parsedGame.endingIndex[v.Name] = &parsedGame.Endings[k]
	}
	LogMsg("  Indexed "+fmt.Sprint(len(parsedGame.Endings))+" endings.", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Indexed " + fmt.Sprint(len(parsedGame.Endings)) + " endings.")
	}

	LogMsg("  Game loaded successfully!", LOG_INFO)
	if Options.GetBool(OPT_SHOW_LOAD_INFO, false) {
		fmt.Println("  Game loaded successfully!")
	}

	parsedGame.Player = Player{
		position:  parsedGame.StartCoords,
		inventory: make([]Item, 0),
	}

	parsedGame.gameFinished = false

	return parsedGame
}

func (g *Game) IsFinished() bool {
	return g.gameFinished
}
