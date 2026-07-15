package main

import (
	"fmt"
	"io"
)

func runCreateSecret(w io.Writer, c ClientConfig) error {
	result, err := createSecret(c.URL, c.Data)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}

func runViewSecret(w io.Writer, c ClientConfig) error {
	result, err := getSecret(c.URL, c.Id)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, result)

	return nil
}
