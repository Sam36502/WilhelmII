package main

import "fmt"

type LogLevel int

const (
	LOG_INFO    = LogLevel(0)
	LOG_WARNING = LogLevel(5)
	LOG_FATAL   = LogLevel(10)
)

const (
	OPT_ERR_DISPL_LVL = "err.displ.lvl" // The level at which messages should be displayed (recc. 5 for debug, 10 for production)
)

func ErrorMsg(msg string, lvl LogLevel) {
	errDisplLvl := Options.GetInt(OPT_ERR_DISPL_LVL, -1)
	if errDisplLvl == -1 {
		fmt.Println("[INFO] No level set for when to display error messages, defaulting to 5.")
	}

	if lvl >= LogLevel(errDisplLvl) {
		var lvlMsg string

		switch lvl {
		case LOG_INFO:
			lvlMsg = "[INFO]"
		case LOG_WARNING:
			lvlMsg = "[WARNING]"
		case LOG_FATAL:
			lvlMsg = "[ERROR]"
		}

		fmt.Println(lvlMsg + msg)
	}
}
