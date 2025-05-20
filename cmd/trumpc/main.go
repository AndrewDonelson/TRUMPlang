// file: cmd/trumpc/main.go
// description: Main entry point for the TRUMP language compiler and interpreter

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AndrewDonelson/trumplang/internal/cmd"
)

func main() {
	// Define command-line flags
	buildCmd := flag.NewFlagSet("build", flag.ExitOnError)
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	inspectCmd := flag.NewFlagSet("inspect", flag.ExitOnError)

	// Add verbosity flags
	buildVerbose := buildCmd.Bool("verbose", false, "Enable verbose output")
	runVerbose := runCmd.Bool("verbose", false, "Enable verbose output")
	buildNoFakeNews := buildCmd.Bool("no-fake-news", false, "Suppress warnings")

	// Check for correct number of arguments
	if len(os.Args) < 2 {
		cmd.PrintUsage()
		os.Exit(1)
	}

	// Parse the command
	switch os.Args[1] {
	case "build":
		buildCmd.Parse(os.Args[2:])
		cmd.BuildTrump(buildCmd.Args(), *buildVerbose, *buildNoFakeNews)
	case "run":
		runCmd.Parse(os.Args[2:])
		cmd.RunTrump(runCmd.Args(), *runVerbose)
	case "create":
		createCmd.Parse(os.Args[2:])
		cmd.CreateTrump(createCmd.Args())
	case "inspect":
		inspectCmd.Parse(os.Args[2:])
		cmd.InspectTrump(inspectCmd.Args())
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		cmd.PrintUsage()
		os.Exit(1)
	}
}
