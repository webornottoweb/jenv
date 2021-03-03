package json

import (
	"os"
	"strconv"
)

// JEnvInt extends base int (platform dependent) type with specific JSON unmarshalling
type JEnvInt int64

// UnmarshalJSON implements unmarshalling with env implementation
func (val *JEnvInt) UnmarshalJSON(data []byte) error {
	env, err := getenv(data)

	if err != nil {
		res, err := strconv.Atoi(string(env))

		if err != nil {
			return err
		}

		*val = JEnvInt(res)

		return nil

	}

	res, err := strconv.Atoi(os.Getenv(string(env)))

	*val = JEnvInt(res)

	return nil
}
