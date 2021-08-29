package wilhelm

var allCommands = []string{
	"help",
	"hitler",
	"hitch",
	"go",
	"end",
}

func CommandSuggest(toComplete string) []string {
	if len(toComplete) == 0 {
		return allCommands
	}

	arr := make([]string, 0)
	for _, c := range allCommands {
		if len(toComplete) <= len(c) && c[:len(toComplete)] == toComplete {
			arr = append(arr, c)
		}
	}
	return arr
}

func (game *Game) ExecuteCommand(command string, args []string) {
	if command == "end" {
		game.gameFinished = true
	}
}
