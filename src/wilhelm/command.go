package wilhelm

import "fmt"

var allCommands = []string{
	"help",
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
	switch command {
	case "help":
		fmt.Println(
			"\n  [ Valid Commands ]\n" +
				"  ---------------------------------------------------------------------------------------------------\n" +
				"  help ..................... List all commands and their required arguments.\n" +
				"  go <direction> ........... You attempt to move in one of the four cardinal directions (N, S, E, W).\n" +
				"  end ...................... End the game instantly.\n",
		)
	case "go":
		if args[0] == "up" {
			fmt.Println("  You broke out of the game.\n  Congratulations; you win.\n  Now leave.")
		} else {
			fmt.Println("  Oh, shit. This isn't done yet...\n  Uhh...\n  All the doors in the room have magically sealed.\n  You are trapped.\n  Game Over.")
		}
		game.gameFinished = true
	case "end":
		game.gameFinished = true
	default:
		fmt.Println("  Eh? I don't understand. Try \"help\" for a list of commands.")
	}
}
