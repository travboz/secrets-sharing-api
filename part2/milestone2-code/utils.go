package main

import (
	"fmt"
	"io"
)

const (
	createDescription = "Create a new secret using the create secret endpoint."
	viewDescription   = "Fetch a secret using the get secret endpoint."

	ActionView   string = "view"
	ActionCreate string = "create"
)

var (
	usageString = `Secret Share CLI

A command-line client for interacting with the Secret Sharing web application.
Create and view secrets through the application's HTTP API.

Usage:
  %s <command> [options]

Available subcommands:
  create    Create a new secret using the create secret endpoint.
  view      Fetch a secret using the get secret endpoint.
`

	subcommandUsage = `%s

Usage: %s [options]

Options:
`
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString, programName)
}

// isGlobalHelp checks for whether the help flag was requested before any subcommand
// was given. e.g. 'cli -h' or 'cli --help' with nothing else.
func isGlobalHelp(args []string) bool {
	if len(args) != 2 {
		return false
	}

	switch args[1] {
	case "-h", "--help":
		return true
	default:
		return false
	}
}
