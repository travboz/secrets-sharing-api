package main

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"testing"
)

var (
	createUsageString = `Create a new secret using the create secret endpoint.

Usage: create [options]

Options:
  -data string
    	The plaintext secret you wish to create (default " ")
  -url string
    	The url of the create endpoint of the secret sharing api (default " ")
`

	viewUsageString = `Fetch a secret using the get secret endpoint.

Usage: view [options]

Options:
  -id string
    	The hashed id of the secret you want to fetch (default " ")
  -url string
    	The url of the create endpoint of the secret sharing api (default " ")
`
)

type paTestConfig struct {
	name      string
	args      []string
	c         ClientConfig
	output    string
	err       error
	parseFunc func(w io.Writer, args []string) (ParseResult, error)
}

func TestParseCreateAndViewArgs(t *testing.T) {
	defaultUrlValue := " "
	defaultDataValue := " "
	defaultIdValue := " "

	testUrl := "http://www.example.com"
	testData := "secret"
	testId := "id"

	tests := []paTestConfig{
		// parseCreateArgs TESTS
		{
			name:      "No flags mean no parse errors",
			parseFunc: parseCreateArgs,
			args:      []string{},
			c:         ClientConfig{Action: ActionCreate, URL: defaultUrlValue, Data: defaultDataValue},
			output:    "",
			err:       nil,
		},
		{
			name:      "No positional args allowed",
			parseFunc: parseCreateArgs,
			args:      []string{"hello"},
			c:         ClientConfig{},
			err:       ErrInvalidPosArgSpecified,
		},
		{
			name:      "Unknown flag",
			parseFunc: parseCreateArgs,
			args:      []string{"--hello"},
			c:         ClientConfig{},
			output: `flag provided but not defined: -hello

` + createUsageString,
			err: errors.New("flag provided but not defined: -hello"),
		},
		{
			name:      "Correct required flags with extra undefined flag means fail",
			parseFunc: parseCreateArgs,
			args:      []string{"--url", testUrl, "--data", testData, "--extra", "hello"},
			c:         ClientConfig{},
			output: `flag provided but not defined: -extra

` + createUsageString,
			err: errors.New("flag provided but not defined: -extra"),
		},
		{
			name:      "Correct required flags with positional arg",
			parseFunc: parseCreateArgs,
			args:      []string{"--url", testUrl, "--data", testId, "hello"},
			c:         ClientConfig{},
			err:       ErrInvalidPosArgSpecified,
		},
		{
			name:      "(happy path) Correct use means success",
			parseFunc: parseCreateArgs,
			args:      []string{"--url", testUrl, "--data", testData},
			c:         ClientConfig{Action: ActionCreate, URL: testUrl, Data: testData},
			err:       nil,
		},
		// parseViewArgs TESTS
		{
			name:      "No flags mean no parse errors",
			parseFunc: parseViewArgs,
			args:      []string{},
			c:         ClientConfig{Action: ActionView, URL: defaultUrlValue, Id: defaultIdValue},
			output:    "",
			err:       nil,
		},
		{
			name:      "No positional args allowed",
			parseFunc: parseViewArgs,
			args:      []string{"hello"},
			c:         ClientConfig{},
			err:       ErrInvalidPosArgSpecified,
		},
		{
			name:      "Unknown flag",
			parseFunc: parseViewArgs,
			args:      []string{"--goodbye"},
			c:         ClientConfig{},
			output: `flag provided but not defined: -goodbye

` + viewUsageString,
			err: errors.New("flag provided but not defined: -goodbye"),
		},
		{
			name:      "Correct required flags with extra undefined flag means fail",
			parseFunc: parseViewArgs,
			args:      []string{"--url", testUrl, "--id", testId, "--extra"},
			c:         ClientConfig{},
			output: `flag provided but not defined: -extra

` + viewUsageString,
			err: errors.New("flag provided but not defined: -extra"),
		},
		{
			name:      "Correct required flags with positional arg",
			parseFunc: parseViewArgs,
			args:      []string{"--url", testUrl, "--id", testId, "goodbye"},
			c:         ClientConfig{},
			err:       ErrInvalidPosArgSpecified,
		},
		{
			name:      "(happy path) Correct use means success",
			parseFunc: parseViewArgs,
			args:      []string{"--url", testUrl, "--id", testId},
			c:         ClientConfig{Action: ActionView, URL: testUrl, Id: testId},
			err:       nil,
		},
	}

	outputBuf := new(bytes.Buffer)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.parseFunc(outputBuf, tc.args)
			if tc.err == nil && err != nil {
				t.Fatalf("want nil error, got: %v\n", err)
			}
			if tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("want error to be: %v, got: %v\n", tc.err, err)
			}

			if !reflect.DeepEqual(result.Config, tc.c) {
				t.Errorf("want config to be: %v, but got: %v", tc.c, result.Config)
			}

			gotMsg := outputBuf.String()
			if len(tc.output) != 0 && gotMsg != tc.output {
				t.Errorf("want stdout message to be: %#v, got: %#v\n", tc.output, gotMsg)
			}

			outputBuf.Reset()
		})
	}
}

func TestParseArgs(t *testing.T) {
	testUrl := "http://www.example.com"
	testData := "secret"
	testId := "id"

	tests := []paTestConfig{
		{
			name:      "empty args",
			err:       ErrEmptyArgs,
			args:      []string{},
			parseFunc: parseArgs,
		},
		{
			name:      "create with correct args returns config",
			err:       nil,
			args:      []string{"create", "--url", testUrl, "--data", testData},
			parseFunc: parseArgs,
			c:         ClientConfig{Action: ActionCreate, URL: testUrl, Data: testData},
		},
		{
			name:      "view with correct args returns config",
			err:       nil,
			args:      []string{"view", "--url", testUrl, "--id", testId},
			parseFunc: parseArgs,
			c:         ClientConfig{Action: ActionView, URL: testUrl, Id: testId},
		},
		{
			name:      "unknown action gives no config and gives invalid action error",
			err:       ErrInvalidAction,
			args:      []string{"unknown"},
			parseFunc: parseArgs,
		},
		// {
		// Unsure about this test because it technically tests
		// validateViewArgs.
		// name: "wrong options for view means fail",
		// },
		// {
		// Unsure about this test because it technically tests
		// validateCreateArgs.
		// name: "wrong options for create means fail",
		// },
	}

	outputBuf := new(bytes.Buffer)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := tc.parseFunc(outputBuf, tc.args)
			if tc.err == nil && err != nil {
				t.Fatalf("want nil error, got: %v\n", err)
			}
			if tc.err != nil && err.Error() != tc.err.Error() {
				t.Fatalf("want error to be: %v, got: %v\n", tc.err, err)
			}

			if !reflect.DeepEqual(result.Config, tc.c) {
				t.Errorf("want config to be: %v, but got: %v", tc.c, result.Config)
			}

			gotMsg := outputBuf.String()
			if len(tc.output) != 0 && gotMsg != tc.output {
				t.Errorf("want stdout message to be: %#v, got: %#v\n", tc.output, gotMsg)
			}

			outputBuf.Reset()
		})
	}
}
