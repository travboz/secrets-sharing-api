package main

import "errors"

var (
	ErrActionNotSpecified = errors.New("action must be specified")
	ErrInvalidAction      = errors.New("invalid action")

	ErrUrlNotSpecified = errors.New("must specify the URL for the secret sharing api")
	ErrInvalidURL      = errors.New("invalid url specified")

	ErrCreateDataOptionEmpty = errors.New("'create' command requires that --data option not be empty")
	ErrCreateIdNotEmpty      = errors.New("'create' command requires empty --id option")

	ErrViewIdEmpty      = errors.New("'view' command requires that --id option not be empty")
	ErrViewDataNotEmpty = errors.New("'view' command requires empty --data option")

	ErrInvalidPosArgSpecified = errors.New("no positional arguments are allowed")
)
