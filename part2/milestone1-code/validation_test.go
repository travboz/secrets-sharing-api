package main

import "testing"

type testConfig struct {
	name string
	c    ClientConfig
	err  error
}

func TestValidateUrlAndAction(t *testing.T) {
	tests := []testConfig{
		{
			name: "Empty client config",
			c:    ClientConfig{},
			err:  ErrActionNotSpecified,
		},
		{
			name: "Action only",
			c:    ClientConfig{Action: "test"},
			err:  ErrUrlNotSpecified,
		},
		{
			name: "Action and URL provided is valid",
			c:    ClientConfig{Action: "action", URL: "url"},
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
	tests := []testConfig{
		{
			name: "Empty client config",
			c:    ClientConfig{},
			err:  ErrCreateDataOptionEmpty,
		},
		{
			name: "Empty data option fails",
			c:    ClientConfig{Action: "action", URL: "url", Data: ""},
			err:  ErrCreateDataOptionEmpty,
		},
		{
			name: "Id must be empty",
			c:    ClientConfig{Action: "action", URL: "url", Data: "data", Id: "id"},
			err:  ErrCreateIdNotEmpty,
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
	tests := []testConfig{
		{
			name: "Empty client config",
			c:    ClientConfig{},
			err:  ErrViewIdEmpty,
		},
		{
			name: "Empty id option fails",
			c:    ClientConfig{Action: "action", URL: "url", Id: ""},
			err:  ErrViewIdEmpty,
		},
		{
			name: "Data must be empty",
			c:    ClientConfig{Action: "action", URL: "url", Data: "data", Id: "id"},
			err:  ErrViewDataNotEmpty,
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
