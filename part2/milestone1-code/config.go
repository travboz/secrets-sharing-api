package main

import "fmt"

type ClientConfig struct {
	Action string // view or create
	Data   string // plaintext secret to store
	Id     string // hash of id returned after creating the secret
	URL    string // url of secret sharing api
}

func (c ClientConfig) String() string {
	return fmt.Sprintf(
		"Action: %s\nData: %s\nId: %s\nURL: %s",
		c.Action,
		c.Data,
		c.Id,
		c.URL,
	)
}
