package main

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
