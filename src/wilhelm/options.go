package wilhelm

import (
	"github.com/magiconair/properties"
)

const (
	OPT_ERR_DISPL_LVL  = "err.display_lvl" // The level at which messages should be displayed (recc. 5 for debug, 10 for production)
	OPT_SHOW_LOAD_INFO = "info.show_load"  // Whether info about the game loading process should be displayed to the user
	OPT_GAMES_DIR      = "files.game_dir"  // Where the game files are held
)

var Options *properties.Properties

func LoadOptions(optionsFile string) {
	Options = properties.MustLoadFile(optionsFile, properties.UTF8)
}
