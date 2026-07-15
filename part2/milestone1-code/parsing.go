package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

type ParseResult struct {
	Config ClientConfig
	Usage  func()
}

var (
	subcommandUsage = `
%s

Usage: %s [options]

Options:
`

	ErrInvalidPosArgSpecified = errors.New("No positional arguments are allowed")
)

func parseCreateArgs(w io.Writer, args []string) (ParseResult, error) {
	c := ClientConfig{}

	subcommandDesc := "Create a new secret using the create secret endpoint."

	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		fmt.Fprintf(w, subcommandUsage, subcommandDesc, fs.Name())
		fs.PrintDefaults()
	}

	fs.StringVar(&c.Data, "data", " ", "The plaintext secret you wish to create")
	fs.StringVar(&c.URL, "url", " ", "The url of the create endpoint of the secret sharing api")

	err := fs.Parse(args)
	if err != nil {
		return ParseResult{}, err
	}

	if fs.NArg() != 0 {
		return ParseResult{}, ErrInvalidPosArgSpecified
	}

	return ParseResult{Config: c, Usage: fs.Usage}, nil
}

func parseViewArgs(w io.Writer, args []string) (ParseResult, error) {
	c := ClientConfig{}

	subcommandDesc := "Fetch a secret using the get secret endpoint."

	fs := flag.NewFlagSet("view", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		fmt.Fprintf(w, subcommandUsage, subcommandDesc, fs.Name())
		fs.PrintDefaults()
	}

	fs.StringVar(&c.Id, "id", " ", "The hashed id of the secret you want to fetch")
	fs.StringVar(&c.URL, "url", " ", "The url of the create endpoint of the secret sharing api")

	err := fs.Parse(args)
	if err != nil {
		return ParseResult{}, err
	}

	if fs.NArg() != 0 {
		return ParseResult{}, ErrInvalidPosArgSpecified
	}

	return ParseResult{Config: c, Usage: fs.Usage}, nil
}
