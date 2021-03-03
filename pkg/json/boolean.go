package json

import (
	"os"
)

// JEnvBool extends base bool type with specific JSON unmarshalling
type JEnvBool bool

// UnmarshalJSON implements unmarshalling with env implementation
func (val *JEnvBool) UnmarshalJSON(data []byte) error {
	env, err := getenv(data)

	if err != nil {
		if string(env) == "true" || string(env) == "1" {
			*val = true
		} else {
			*val = false
		}

		return nil
	}

	res := os.Getenv(string(env))

	if res == "true" || res == "1" {
		*val = true
	} else {
		*val = false
	}

	return nil
}
