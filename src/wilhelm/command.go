package wilhelm

import (
	"fmt"
	"strings"
)

type Command struct {
	name        string
	description string
	argsDesc    []string
	executeFunc func(*Game, []string) error
}

var AllCommands []Command

func LoadCommands() {
	AllCommands = []Command{
		{
			name:        "help",
			description: "Lists all commands and their requireds arguments",
			argsDesc:    []string{},
			executeFunc: helpCommand,
		},
		{
			name:        "go",
			description: "You attempt to move in one of the four cardinal directions (N, S, E, W)",
			argsDesc:    []string{"direction"},
			executeFunc: goCommand,
		},
		{
			name:        "end",
			description: "Ends the game for this session",
			argsDesc:    []string{},
			executeFunc: endCommand,
		},
	}
}

func CommandSuggest(toComplete string) []string {
	var commNames []string
	for _, c := range AllCommands {
		commNames = append(commNames, c.name)
	}

	if len(toComplete) == 0 {
		return commNames
	}

	var suggestions []string
	for _, n := range commNames {
		frac := n[:len(toComplete)]
		if toComplete == frac {
			suggestions = append(suggestions, n)
		}
	}

	return suggestions
}

// Validate command input and execute commands
func (game *Game) ExecuteCommand(command string, args []string) {
	command = strings.ToLower(command)

	commandFound := false
	for _, comm := range AllCommands {
		if comm.name == command {
			commandFound = true
			if len(args) >= len(comm.argsDesc) {

				err := comm.executeFunc(game, args)
				if err != nil {
					// TODO: Something, idrk
					fmt.Println(err.Error())
				}

			} else {
				// TODO: Display game error text
				fmt.Println("  I need more information. Check 'help' for more info.")
			}
		}
	}

	if !commandFound {
		fmt.Println("  What? I don't understand that.")
	}

}

func helpCommand(game *Game, args []string) error {

	length := 0
	entries := []string{}

	// Make left side of table
	for _, c := range AllCommands {
		entry := c.name
		for _, a := range c.argsDesc {
			entry += " <" + a + ">"
		}
		entry += "..."

		if len(entry) > length {
			length = len(entry)
		}
		entries = append(entries, entry)
	}

	// Extend dots to max length
	for i, e := range entries {
		numDots := length - len(e)
		for n := 0; n < numDots; n++ {
			entries[i] += "."
		}
	}

	// Add Descriptions
	for i, _ := range entries {
		command := AllCommands[i]
		entries[i] += command.description
		if len(entries[i]) > length {
			length = len(entries[i])
		}
	}

	fmt.Print("\n  [ Valid Commands ]\n --")
	for n := 0; n < length; n++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	for _, e := range entries {
		fmt.Println("  " + e)
	}
	fmt.Print("\n")

	return nil
}

func goCommand(game *Game, args []string) error {
	if args == nil || len(args) < 1 {
		return MissingArgumentsError("  I need to know where you want to go (north, south, east, west)")
	}
	dir, err := parseDirection(args[0])
	if err != nil {
		return InvalidArgumentsError("  What? Where's '" + args[0] + "'?\n  Please just stick to the cardinal directions.")
	}

	// Check for Doors
	game.GetRoom(game.Player.position)

	// Move Player
	game.Player.position = newCoords

	return nil
}

func endCommand(game *Game, args []string) error {
	game.gameFinished = true
	return nil
}
