package main

import (
	"flag"
	"fmt"
	"io"
)

type ParseResult struct {
	Config ClientConfig
	Usage  func()
}

// parseArgs parses the arguments given to the CLI.
// How it parses the arguments depends on the subcommand action chosen.
func parseArgs(w io.Writer, args []string) (ParseResult, error) {
	subcommand := args[0]
	options := args[1:]

	switch subcommand {
	case ActionCreate:
		return parseCreateArgs(w, options)
	case ActionView:
		return parseViewArgs(w, options)
	default:
		return ParseResult{}, ErrInvalidAction
	}
}

// parseCreateArgs parses the arguments based on the subcommand being 'create'.
// It then returns the result wrapped in with the usage info of 'create'.
func parseCreateArgs(w io.Writer, args []string) (ParseResult, error) {
	c := ClientConfig{}
	result := ParseResult{}

	fs := flag.NewFlagSet("create", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		fmt.Fprintf(
			w,
			subcommandUsage,
			createDescription,
			fs.Name(),
		)
		fs.PrintDefaults()
	}
	result.Usage = fs.Usage

	fs.StringVar(&c.Data, "data", " ", "The plaintext secret you wish to create")
	fs.StringVar(&c.URL, "url", " ", "The url of the create endpoint of the secret sharing api")

	err := fs.Parse(args)
	if err != nil {
		return result, err
	}

	if fs.NArg() != 0 {
		return result, ErrInvalidPosArgSpecified
	}

	c.Action = ActionCreate
	result.Config = c

	return result, nil
}

// parseViewArgs parses the arguments based on the subcommand being 'view'.
// It then returns the result wrapped in with the usage info of 'view'.
func parseViewArgs(w io.Writer, args []string) (ParseResult, error) {
	c := ClientConfig{}
	result := ParseResult{}

	fs := flag.NewFlagSet("view", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		fmt.Fprintf(
			w,
			subcommandUsage,
			viewDescription,
			fs.Name(),
		)
		fs.PrintDefaults()
	}
	result.Usage = fs.Usage

	fs.StringVar(&c.Id, "id", " ", "The hashed id of the secret you want to fetch")
	fs.StringVar(&c.URL, "url", " ", "The url of the create endpoint of the secret sharing api")

	err := fs.Parse(args)
	if err != nil {
		return result, err
	}

	if fs.NArg() != 0 {
		return result, ErrInvalidPosArgSpecified
	}

	c.Action = ActionView
	result.Config = c

	return result, nil
}
