package main

import (
	"fmt"
	"io/ioutil"
	"src/src/wilhelm"
	"strings"
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

	// Initialize Enging
	wilhelm.LoadCommands()

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
		wilhelm.DisplayText(currRoom.Description)

		// Get User Input
		commandStr := ""
		prompt := &survey.Input{
			Message: " What will you do?\n    > ",
			Suggest: wilhelm.CommandSuggest,
		}
		survey.AskOne(prompt, &commandStr)

		// Execute User's Wishes
		commandArr := strings.Split(commandStr, " ")
		var args []string = nil
		if len(commandArr) > 1 {
			args = commandArr[1:]
		}
		game.ExecuteCommand(commandArr[0], args)
		wilhelm.WaitForEnter()
	}

	/// ENDING THE GAME ///

	wilhelm.ClearScreen()
	// TODO: Make not shitty and boring
	fmt.Println("  Goodbye!")
	time.Sleep(1 * time.Second)
	wilhelm.ClearScreen()

}
