package main

import (
	"fmt"

	"github.com/Dids/xcodebuild-remote/commander"
)

// Version is set dynamically when building
var Version = "0.0.0"

// Main entrypoint
func main() {
	// Create the individual commands
	commander.Create("help", "display help", false, false, func(value string) { generateHelp() })
	commander.Create("version", "display version number", false, false, func(value string) { fmt.Println("xcodebuild-remote version", Version) })
	commander.Create("url", "repository url", true, true, func(value string) { fmt.Println("url called with value:", value) })

	// Parse the arguments
	commander.ParseArgs()

	// Print a list of unused arguments
	fmt.Println("Unused arguments:", commander.ExtraArgs)

	/*// TODO: Check if "xcodebuild" exists and bail out if it doesn't

	// TODO: Implement the 4 different git repo url formats (https://, git://, git+ssh://, user/repo for GitHub)

	// Create the xcodebuild command
	xcodebuild := exec.Command("xcodebuild")

	// Redirect all output to terminal
	xcodebuild.Stdout = os.Stdout
	xcodebuild.Stderr = os.Stderr

	// TODO: Pass any leftover arguments to xcodebuild

	// TODO: Re-enable this when it makes more sense
	// Finally run the command
	// xcodebuild.Run()*/
}

// Generates a detailed list of available commands
func generateHelp() {
	fmt.Println()
	fmt.Println("Example usage:")
	fmt.Println("  xcodebuild-remote --url [repository url]")
	fmt.Println()
	fmt.Println("Available arguments:")
	outputString := "  %-*s  %s\n"
	padding := 30
	for _, command := range commander.Commands {
		argumentString := command.Argument + "," + command.ShortArgument
		if command.NeedsValue {
			argumentString += " [arg]"
			fmt.Printf(outputString, padding, argumentString, command.Description)
		} else {
			fmt.Printf(outputString, padding, argumentString, command.Description)
		}
	}
	fmt.Println()
}
