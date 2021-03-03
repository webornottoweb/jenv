package json

import (
	"os"
)

// JEnvString extends base string type with specific JSON unmarshalling
type JEnvString string

// UnmarshalJSON implements unmarshalling with env implementation
func (val *JEnvString) UnmarshalJSON(data []byte) error {
	env, err := getenv(data)

	if err != nil {
		*val = JEnvString(data)
		return nil
	}

	*val = JEnvString(os.Getenv(string(env)))
	return nil
}
