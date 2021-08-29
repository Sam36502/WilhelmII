package main

import (
	"fmt"
	"io/ioutil"
	"src/src/wilhelm"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mgutz/ansi"
)

const (
	OPTIONS_FILE   = "data/options.properties"
	ENGINE_VERSION = "1.0"
)

func init() {
	wilhelm.LoadOptions(OPTIONS_FILE)
}

func main() {

	/// LOADING GAME ///

	// TODO: Make not dogshit (Add ANSI colours and nice formatting)
	// Show main menu message
	fmt.Printf(
		"\n   WILHELM II\n"+
			"  ------------\n\n"+
			"  Version %v\n"+
			"  By Bismarck\n\n",
		ENGINE_VERSION,
	)

	// Find all game files in the 'games' folder
	gameDir := wilhelm.Options.GetString(wilhelm.OPT_GAMES_DIR, "games")
	files, err := ioutil.ReadDir(gameDir)
	if err != nil {
		wilhelm.LogMsg("Failed to read games directory.", wilhelm.LOG_FATAL)
		return
	}

	if len(files) == 0 {
		fmt.Println("  " + ansi.Color("No Games found in '"+gameDir+"' directory.\n", "red"))
		return
	}

	fileNames := make([]string, 0, 10)
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	file := ""
	prompt := &survey.Select{
		Message: "Available Games:",
		Options: fileNames,
	}
	survey.AskOne(prompt, &file)

	// Load Game
	game := wilhelm.LoadGame(gameDir + "/" + file)
	player := wilhelm.NewPlayer(game)

	///	MAIN GAMEPLAY LOOP ///
	for !game.IsFinished() {

		// Describe Room
		wilhelm.ClearScreen()

		currRoom := game.GetRoom(player.GetCoords())
		wilhelm.DisplayText(currRoom.GetDescription())

		// Get User Input
		command := ""
		prompt := &survey.Input{
			Message: " What will you do?\n    > ",
			Suggest: wilhelm.CommandSuggest,
		}
		survey.AskOne(prompt, &command)

		// Execute User's Wishes
		game.ExecuteCommand(command, []string{})

	}

	/// ENDING THE GAME ///

	wilhelm.ClearScreen()
	fmt.Println("  Goodbye!")
	time.Sleep(3 * time.Second)

}
