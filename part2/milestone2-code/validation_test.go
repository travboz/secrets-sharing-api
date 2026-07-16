package main

import "testing"

type testConfig struct {
	name string
	c    ClientConfig
	err  error
}

func TestValidateUrlAndAction(t *testing.T) {
	testUrl := "http://www.example.com"
	testAction := "action"

	tests := []testConfig{
		{
			name: "Empty client config",
			c:    ClientConfig{},
			err:  ErrActionNotSpecified,
		},
		{
			name: "Action only",
			c:    ClientConfig{Action: testAction},
			err:  ErrUrlNotSpecified,
		},
		{
			name: "Invalid url endpoint",
			c:    ClientConfig{Action: testAction, URL: "blahblah"},
			err:  ErrInvalidURL,
		},
		{
			name: "(happy path) Action and URL provided is valid",
			c:    ClientConfig{Action: testAction, URL: testUrl},
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateUrlAndAction(tc.c)

			if tc.err != nil && err != tc.err {
				t.Errorf("want error to be: %q, got %q\n", tc.err, err)
			}

			if tc.err == nil && err != nil {
				t.Errorf("want nil error, but got: %q\n", err)
			}
		})
	}
}

func TestValidateCreateArgs(t *testing.T) {
	testUrl := "http://www.example.com"
	testAction := "action"

	tests := []testConfig{
		{
			// Uncertain about this test - it actually tests whether
			// 'validateUrlAndAction' is working.
			name: "Empty client config",
			c:    ClientConfig{},
			err:  ErrActionNotSpecified,
		},
		{
			name: "Empty data option fails",
			c:    ClientConfig{Action: testAction, URL: testUrl, Data: ""},
			err:  ErrCreateDataOptionEmpty,
		},
		{
			name: "Id must be empty",
			c:    ClientConfig{Action: testAction, URL: testUrl, Data: "data", Id: "id"},
			err:  ErrCreateIdNotEmpty,
		},
		{
			name: "(happy path) Create with non-empty data and empty id is successful",
			c:    ClientConfig{Action: testAction, URL: testUrl, Data: "data"},
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateCreateArgs(tc.c)

			if tc.err != nil && err != tc.err {
				t.Errorf("want error to be: %q, got %q\n", tc.err, err)
			}

			if tc.err == nil && err != nil {
				t.Errorf("want nil error, but got: %q\n", err)
			}
		})
	}
}

func TestValidateViewArgs(t *testing.T) {
	testUrl := "http://www.example.com"
	testAction := "view"

	tests := []testConfig{
		{
			// Uncertain about this test - it actually tests whether
			// 'validateUrlAndAction' is working.
			name: "Empty client config",
			c:    ClientConfig{},
			err:  ErrActionNotSpecified,
		},
		{
			name: "Empty id option fails",
			c:    ClientConfig{Action: testAction, URL: testUrl, Id: ""},
			err:  ErrViewIdEmpty,
		},
		{
			name: "Data must be empty",
			c:    ClientConfig{Action: testAction, URL: testUrl, Data: "data", Id: "id"},
			err:  ErrViewDataNotEmpty,
		},
		{
			name: "(happy path) View with id and empty data succeeds",
			c:    ClientConfig{Action: testAction, URL: testUrl, Id: "id"},
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateViewArgs(tc.c)

			if tc.err != nil && err != tc.err {
				t.Errorf("want error to be: %q, got %q\n", tc.err, err)
			}

			if tc.err == nil && err != nil {
				t.Errorf("want nil error, but got: %q\n", err)
			}
		})
	}
}

func TestValidateConfig(t *testing.T) {
	testUrl := "http://www.example.com"

	tests := []testConfig{
		{
			// Uncertain about this test - it actually tests whether
			// 'validateUrlAndAction' is working.
			name: "Empty client config fails",
			c:    ClientConfig{},
			err:  ErrInvalidAction,
		},
		{
			name: "Invalid action fails",
			c:    ClientConfig{Action: "invalid"},
			err:  ErrInvalidAction,
		},
		{
			name: "(happy path) Create action and valid options succeed",
			c:    ClientConfig{Action: "create", URL: testUrl, Data: "data", Id: ""},
			err:  nil,
		},
		{
			name: "(happy path) View action and valid options succeed",
			c:    ClientConfig{Action: "view", URL: testUrl, Data: "", Id: "id"},
			err:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateConfig(tc.c)

			if tc.err != nil && err != tc.err {
				t.Errorf("want error to be: %q, got %q\n", tc.err, err)
			}

			if tc.err == nil && err != nil {
				t.Errorf("want nil error, but got: %q\n", err)
			}
		})
	}
}
