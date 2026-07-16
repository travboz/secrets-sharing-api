package main

import (
	"fmt"
	"io"
	"slices"
)

const (
	programName = "secret-share"

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
`

	subcommandUsage = `
%s

Usage: %s [options]

Options:
`
)

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString, programName)
}

// helpCalled is a quick check for whether the help flag is present in the args slice.
func helpCalled(args []string) bool {
	if slices.Contains(args, "--help") || slices.Contains(args, "-h") {
		return true
	}

	return false
}
