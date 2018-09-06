package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Version is set dynamically when building
var Version = "0.0.0"

// TODO: Utilize this instead of hardcoding the arguments
type cliArgs struct {
	url string
}

func main() {
	// Get the command line arguments
	args := os.Args[1:]
	fmt.Println("args:", args)

	// Validate that at least a single argument exists
	if len(args) == 0 || (getArgument("--url") == "" && getArgument("-u") == "") {
		fmt.Println("Missing required argument: --url")
		os.Exit(1)
	}

	// TODO: Somehow check if --help/-h exists and parse it

	// TODO: Somehow check if --version/-v exists and parse it

	// TODO: Check if --url/-u exists and parse it
	url := getArgument("--url")
	if url != "" {
		fmt.Println("Got url:", url)
	}

	// FIXME: Don't do this, but instead scan the args array for the argument and its value
	urlArg := os.Args[0]
	fmt.Println("urlArg:", urlArg)

	// TODO: Check if "xcodebuild" exists and bail out if it doesn't

	// TODO: Implement the 4 different git repo url formats (https://, git://, git+ssh://, user/repo for GitHub)

	// Create the xcodebuild command
	xcodebuild := exec.Command("xcodebuild")

	// Redirect all output to terminal
	xcodebuild.Stdout = os.Stdout
	xcodebuild.Stderr = os.Stderr

	// TODO: Pass any leftover arguments to xcodebuild

	// TODO: Re-enable this when it makes more sense
	// Finally run the command
	// xcodebuild.Run()
}

func getArgument(argument string) string {
	for i, arg := range os.Args[1:] {
		// Check if the argument name matches
		if arg == argument {
			fmt.Println("Found argument:", argument)

			// Check if the argument has a value
			if len(os.Args[1:]) > i+1 {
				value := os.Args[1:][i+1]
				fmt.Println("Found value:", value)
				return value
			}
			fmt.Println("Value missing for argument:", argument)
		}
	}
	return ""
}
