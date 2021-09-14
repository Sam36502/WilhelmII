package wilhelm

import "fmt"

// The player entered an invalid command
type InvalidCommandError string

func (e InvalidCommandError) Error() string {
	return fmt.Sprintf("Invalid Command '%v'.", string(e))
}

// The command didn't get the arguments it required
type MissingArgumentsError string

func (e MissingArgumentsError) Error() string {
	return string(e)
}

// The arguments received were invalid
type InvalidArgumentsError string

func (e InvalidArgumentsError) Error() string {
	return string(e)
}

// The player entered an invalid direction
type InvalidDirectionError string

func (e InvalidDirectionError) Error() string {
	return fmt.Sprintf("Invalid Direction '%v'.", string(e))
}
