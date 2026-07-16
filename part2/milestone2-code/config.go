package main

import "fmt"

// ClientConfig holds the data used to determine what action to perform, and its relevant data.
type ClientConfig struct {
	Action string // view or create
	Data   string // plaintext secret to store
	Id     string // hash of id returned after creating the secret
	URL    string // url of secret sharing api
}

// Pretty printing
func (c ClientConfig) String() string {
	return fmt.Sprintf(
		"Action: %s\nData: %s\nId: %s\nURL: %s",
		c.Action,
		c.Data,
		c.Id,
		c.URL,
	)
}
