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
	if len(args) == 0 {
		fmt.Println("Missing required argument (use --help for more information)")
		os.Exit(1)
	}

	// Check if --help/-h is set
	if getArgument("--help", true) != "" || getArgument("-h", true) != "" {
		// TODO: Actually implement the help command
		fmt.Println("WARNING: --help not implemented")
		os.Exit(0)
	}

	// Check if --version/-v is set
	if getArgument("--version", true) != "" || getArgument("-v", true) != "" {
		fmt.Println("Version:", "v"+Version)
		os.Exit(0)
	}

	// Validate that --url/-u exists
	if getArgument("--url", false) == "" || getArgument("-u", false) == "" {
		fmt.Println("Missing required argument: --url")
		os.Exit(1)
	}

	// Check if --url/-u exists and parse it
	url := getArgument("--url", false)
	if url != "" {
		fmt.Println("Got url:", url)
	}

	// FIXME: How the heck do we get any leftover args? Use our custom struct and construct it with that?

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

func getArgument(argument string, skipValue bool) string {
	for i, arg := range os.Args[1:] {
		// Check if the argument name matches
		if arg == argument {
			fmt.Println("Found argument:", argument)

			// Check if we don't need a value
			if skipValue {
				return argument
			}

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
