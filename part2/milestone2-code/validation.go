package main

import "net/url"

// validateConfig validates the values in the config given which action subcommand is picked.
func validateConfig(cfg ClientConfig) error {
	switch cfg.Action {
	case ActionCreate:
		return validateCreateArgs(cfg)
	case ActionView:
		return validateViewArgs(cfg)
	default:
		return ErrInvalidAction
	}
}

func validateUrlAndAction(c ClientConfig) error {
	if c.Action == "" {
		return ErrActionNotSpecified
	}
	if c.URL == "" {
		return ErrUrlNotSpecified
	}

	if _, err := url.ParseRequestURI(c.URL); err != nil {
		return ErrInvalidURL
	}

	return nil

	// Redundant because action is checked in validateConfig
	// switch c.Action {
	// case ActionView, ActionCreate:
	// 	return nil
	// default:
	// 	return ErrInvalidAction
	// }
}

func validateCreateArgs(c ClientConfig) error {
	if err := validateUrlAndAction(c); err != nil {
		return err
	}

	// Redundant because if action isn't 'create', this code is never actually reached
	// if c.Action != ActionCreate {
	// 	return ErrInvalidAction
	// }

	if c.Data == "" || c.Data == " " {
		return ErrCreateDataOptionEmpty
	}
	if c.Id != "" {
		return ErrCreateIdNotEmpty
	}

	return nil
}

func validateViewArgs(c ClientConfig) error {
	if err := validateUrlAndAction(c); err != nil {
		return err
	}

	// Redundant because if action isn't 'view', this code is never actually reached
	// if c.Action != ActionView {
	// 	return ErrInvalidAction
	// }

	if c.Id == "" || c.Id == " " {
		return ErrViewIdEmpty
	}
	if c.Data != "" {
		return ErrViewDataNotEmpty
	}

	return nil
}
