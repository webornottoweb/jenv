package json

import (
	"os"
	"strings"
)

// JEnvString extends base string type with specific JSON unmarshalling
type JEnvString string

// UnmarshalJSON implements unmarshalling with env implementation
func (val *JEnvString) UnmarshalJSON(data []byte) error {
	// omit leading and closing quotes
	data = data[1 : len(data)-1]

	if !strings.HasPrefix(string(data), "env[") || !strings.HasSuffix(string(data), "]") {
		*val = JEnvString(data)
		return nil
	}

	env := data[4 : len(data)-1]

	*val = JEnvString(os.Getenv(string(env)))

	return nil
}
