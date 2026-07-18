package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

const programName = "secret-share"

func main() {
	// Check if length of args is valid, we expect: cli subcommand
	if len(os.Args) < 2 {
		printUsage(os.Stderr)
		os.Exit(1)
	}

	// Check for if help has been called with no subcommand.
	if isGlobalHelp(os.Args) {
		printUsage(os.Stdout)
		os.Exit(0)
	}

	// We have at least 2 args, so determine if a subcommand has been called.
	result, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		// Check if subcommand help has been called - so no need to call as fs.Usage() was called.
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = validateConfig(result.Config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	output, err := performAction(result.Config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, output)
}
