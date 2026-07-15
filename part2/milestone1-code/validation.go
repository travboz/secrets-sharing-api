package main

import "errors"

var (
	ErrActionNotSpecified    = errors.New("Action must be specified")
	ErrUrlNotSpecified       = errors.New("Must specify the URL for the secret sharing api")
	ErrCreateDataOptionEmpty = errors.New("'create' command requires that --data option not be empty")
	ErrCreateIdNotEmpty      = errors.New("'create' command requires empty --id option")
	ErrViewIdEmpty           = errors.New("'view' command requires that --id option not be empty")
	ErrViewDataNotEmpty      = errors.New("'view' command requires empty --data option")
)

func validateUrlAndAction(c ClientConfig) error {
	if c.Action == "" {
		return ErrActionNotSpecified
	}
	if c.URL == "" {
		return ErrUrlNotSpecified
	}

	return nil
}

func validateCreateArgs(c ClientConfig) error {
	if c.Data == "" || c.Data == " " {
		return ErrCreateDataOptionEmpty
	}
	if c.Id != "" {
		return ErrCreateIdNotEmpty
	}

	return validateUrlAndAction(c)
}

func validateViewArgs(c ClientConfig) error {
	if c.Id == "" || c.Id == " " {
		return ErrViewIdEmpty
	}

	if c.Data != "" {
		return ErrViewDataNotEmpty
	}

	return validateUrlAndAction(c)
}
