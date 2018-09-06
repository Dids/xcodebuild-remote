package commander

import (
	"fmt"
	"os"
)

// Command represents a single command
type Command struct {
	Name          string
	Description   string
	Argument      string
	ShortArgument string
	NeedsValue    bool
	Required      bool
	Action        func(value string)
}

// Commands is an array of all the created commands
var Commands []Command

// ExtraArgs is an array of unused arguments (no commands match them)
var ExtraArgs []string

// Create is used to create new commands
func Create(name string, description string, needsValue bool, required bool, action func(value string)) Command {
	// TODO: We should ideally handle short arg name clashes as well (bound to happen with --verbose and --version, for instance)
	// Create long and short arguments automatically from the name
	longArg := "--" + name
	shortArg := "-" + string([]rune(name)[0])
	newCommand := Command{name, description, longArg, shortArg, needsValue, required, action}
	//fmt.Println("Created a new command:", newCommand)

	// Store the command locally, so we can reference it later (eg. when parsing args)
	Commands = append(Commands, newCommand)

	// Return the newly created command
	return newCommand
}

// ParseArgs will parse the command line arguments and match them with commands
func ParseArgs() {
	// Get the command line arguments
	args := os.Args[1:]
	//fmt.Println("args:", args)

	// Loop through all the command line arguments
	for _, arg := range args {
		// Loop through all of our commands to find which args are considered extra
		isCommand := false
		for _, command := range Commands {
			if arg == command.Argument || arg == command.ShortArgument {
				isCommand = true
			}
		}
		// Store unused arguments in an array
		if !isCommand {
			ExtraArgs = append(ExtraArgs, arg)
		}
	}

	// Loop through all the command line arguments (again)
	for i, arg := range args {
		// Loop through all of our commands and execute them if necessary
		for _, command := range Commands {
			// Check if the long or short argument matches
			if arg == command.Argument || arg == command.ShortArgument {
				//fmt.Println("Argument", arg, "matches the command", command.Argument)

				// Get the value (if any)
				value := ""
				if command.NeedsValue {
					if len(args) > i+1 {
						value = args[i+1]
						//fmt.Println("Command value found:", value)
					} else {
						// Exit if a value is required but missing
						fmt.Println("Missing value for argument:", arg)
						os.Exit(1)
					}
				}

				// Execute the command function
				command.Action(value)

				// Stop processing commands
				return
			}
		}
	}

	// If we get this far, we should check for required commands
	for _, command := range Commands {
		if command.Required {
			// Exit if an argument is required, but hasn't been specified (we wouldn't be here otherwise, right?)
			fmt.Println("Missing required argument:", command.Argument)
			fmt.Println()
			fmt.Println("Use 'xcodebuild-remote --help' to display all available arguments")
			os.Exit(1)
		}
	}
}
