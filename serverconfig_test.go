package serverconfig

import (
	"reflect"
	"strings"
	"testing"
)

const validYaml = `
Addr: localhost:8080
ReadTimeoutMs: 123
ReadHeaderTimeoutMs: 234
WriteTimeoutMs: 345
IdleTimeoutMs: 456
MaxHeaderBytes: 567
`

const invalidYaml = `
Addr = localhost:8080
ReadTimeoutMs =  123
ReadHeaderTimeoutMs = 234
WriteTimeoutMs = 345
IdleTimeoutMs = 456
MaxHeaderBytes = 567
`

const negativeDuration = `
Addr: localhost:8080
ReadTimeoutMs: -123
ReadHeaderTimeoutMs: 234
WriteTimeoutMs: -345
IdleTimeoutMs: 456
MaxHeaderBytes: 567
`

const validEmptyFields = `
Addr: localhost:8080
`

const invalidEmptyFields = `
ReadTimeoutMs: 123
ReadHeaderTimeoutMs: 234
WriteTimeoutMs: 345
IdleTimeoutMs: 456
MaxHeaderBytes: 567
`

func TestFromReader(t *testing.T) {
	testcases := []struct {
		Name     string
		Yaml     string
		Expected *Config
		IsError  bool
	}{
		{
			Name: "valid yaml",
			Yaml: validYaml,
			Expected: &Config{
				Addr:                "localhost:8080",
				ReadTimeoutMs:       123,
				ReadHeaderTimeoutMs: 234,
				WriteTimeoutMs:      345,
				IdleTimeoutMs:       456,
				MaxHeaderBytes:      567,
			},
			IsError: false,
		},
		{
			Name:     "invalid yaml",
			Yaml:     invalidYaml,
			Expected: nil,
			IsError:  true,
		},
		{
			Name:     "negative duration",
			Yaml:     negativeDuration,
			Expected: nil,
			IsError:  true,
		},
		{
			Name: "valid empty fields",
			Yaml: validEmptyFields,
			Expected: &Config{
				Addr:                "localhost:8080",
				ReadTimeoutMs:       0,
				ReadHeaderTimeoutMs: 0,
				WriteTimeoutMs:      0,
				IdleTimeoutMs:       0,
				MaxHeaderBytes:      0,
			},
			IsError: false,
		},
		{
			Name:     "invalid empty",
			Yaml:     invalidEmptyFields,
			Expected: nil,
			IsError:  true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			gotConf, err := FromReader(strings.NewReader(tc.Yaml))

			gotIsError := err != nil
			if gotIsError != tc.IsError {
				t.Errorf("err != nil: want %v, got %v\ngotten err: %v\n", tc.IsError, gotIsError, err)
				return
			}

			if !reflect.DeepEqual(gotConf, tc.Expected) {
				t.Errorf("\ngot: %+v\nexpected: %+v\n", gotConf, tc.Expected)
				return
			}
		})
	}
}
