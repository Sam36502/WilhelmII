package wilhelm

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

type LogLevel int

const (
	LOG_INFO    = LogLevel(0)
	LOG_WARNING = LogLevel(5)
	LOG_FATAL   = LogLevel(10)
)

var noLvlSetMessageDisplayed = false

/*
	Logs a message to the console based on the message level provided and the err.display_lvl option
*/
func LogMsg(msg string, lvl LogLevel) {
	errDisplLvl := Options.GetInt(OPT_ERR_DISPL_LVL, -1)
	if errDisplLvl == -1 && !noLvlSetMessageDisplayed {
		fmt.Println(ansi.Color("[INFO] No level set for when to display error messages, defaulting to 5.", "cyan"))
		noLvlSetMessageDisplayed = true
	}

	if lvl >= LogLevel(errDisplLvl) {

		switch lvl {
		case LOG_INFO:
			fmt.Println(ansi.Color("[INFO] "+msg, "cyan"))
		case LOG_WARNING:
			fmt.Println(ansi.Color("[WARNING] "+msg, "yellow"))
		case LOG_FATAL:
			fmt.Println(ansi.Color("[ERROR] "+msg, "red"))
		}
	}

	if lvl >= LOG_FATAL {
		os.Exit(1)
	}
}

/*
	Clears the terminal screen with ANSI control codes.
*/
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

/*
	Displays an object description filtering out escape sequences
*/
func DisplayText(text []string) {
	fmt.Print("\n")
	for _, l := range text {
		if l[:0] == "/" {
			// TODO: Parse Commands
		} else {
			fmt.Println("  " + l)
		}
	}
}
