package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Dids/xcodebuild-remote/commander"
	git "github.com/gogs/git-module"
)

// Version is set dynamically when building
var Version = "0.0.0"

// Main entrypoint
func main() {
	// TODO: Add support for a "branch" argument, which should be optional
	// Create the individual commands
	commander.Create("help", "display help", false, false, func(value string) { generateHelp() })
	commander.Create("version", "display version number", false, false, func(value string) { fmt.Println("xcodebuild-remote version", Version) })
	commander.Create("url", "repository url", true, true, func(value string) {
		fmt.Println("url called with value:", value)

		// TODO: Handle errors and validate repoOwner and repoValue
		// Parse the repository owner and name
		splitURL := strings.Split(value, "/")
		repoOwner := splitURL[len(splitURL)-2]
		repoName := splitURL[len(splitURL)-1]

		fmt.Println("Repository owner:", repoOwner)
		fmt.Println("Repository name:", repoName)

		// Create a temporary path based on the repository owner and name
		tempPath := getTempPath(repoOwner + "/" + repoName)

		// Clone the repository
		gitClone(tempPath, value)

		// Get any extra arguments
		extraArgs := commander.ExtraArgs
		fmt.Println("xcodebuild arguments:", extraArgs)

		// Build the project
		xcodebuild(tempPath, extraArgs)

		// TODO: Open "tempPath/build" if it exists, otherwise fallback to "tempPath"
		// Open the project directory
		exec.Command("open", tempPath).Start()
	})

	// Parse the arguments
	commander.ParseArgs()
}

// Generates a detailed list of available commands
func generateHelp() {
	fmt.Println()
	fmt.Println("Example usage:")
	fmt.Println("  xcodebuild-remote --url <repository url> [optional xcodebuild arguments]")
	fmt.Println()
	fmt.Println("Available arguments:")
	outputString := "  %-*s  %s\n"
	padding := 30
	for _, command := range commander.Commands {
		argumentString := command.Argument + "," + command.ShortArgument
		if command.NeedsValue {
			argumentString += " <arg>"
			fmt.Printf(outputString, padding, argumentString, command.Description)
		} else {
			fmt.Printf(outputString, padding, argumentString, command.Description)
		}
	}
	fmt.Println()
}

func gitClone(cwd string, url string) {
	// TODO: Validate the url with this: git ls-remote the-url-to-test
	//       It returns 0 in success, otherwise non-zero

	// TODO: Get this from the command (default to "master")
	branch := "master"

	// TODO: Implement the 4 different git repo url formats (https://, git://, git+ssh://, user/repo for GitHub)
	// TODO: Clone or checkout depending on if it already exists (clean it in-between builds though?)
	git.Clone(url, cwd, git.CloneOptions{Branch: branch, Bare: false, Quiet: false})
}

func xcodebuild(cwd string, args []string) {
	// TODO: Validate cwd

	// Create the xcodebuild command
	xcodebuild := exec.Command("xcodebuild")

	// Set the working directory
	xcodebuild.Dir = cwd

	// Redirect all output to terminal
	xcodebuild.Stdout = os.Stdout
	xcodebuild.Stderr = os.Stderr

	// Pass any leftover arguments to xcodebuild
	args = prepend(args, "xcodebuild")
	xcodebuild.Args = args

	// TODO: Remove this?
	// Print the command
	printCommand(xcodebuild)

	// Finally run xcodebuild
	xcodebuild.Run()
}

func prepend(arr []string, item string) []string {
	return append([]string{item}, arr...)
}

// TODO: Remove this?
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func getTempPath(project string) string {
	tempPath := "/tmp/" + project
	fmt.Println("Temp path:", tempPath)
	return tempPath
}
