package main

import (
	"fmt"
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mgutz/ansi"
)

const (
	OPTIONS_FILE   = "data/options.properties"
	ENGINE_VERSION = "1.0"
)

func init() {
	LoadOptions(OPTIONS_FILE)
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
	gameDir := Options.GetString(OPT_GAMES_DIR, "games")
	files, err := ioutil.ReadDir(gameDir)
	if err != nil {
		LogMsg("Failed to read games directory.", LOG_FATAL)
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
	game := LoadGame(gameDir + "/" + file)

	///	MAIN GAMEPLAY LOOP ///
	fmt.Println("Ending: " + game.endingIndex["The Content Ending"].Name)

	/// ENDING THE GAME ///

}
