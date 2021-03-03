package json

import (
	"errors"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getenv(data []byte) ([]byte, error) {
	// omit leading and closing quotes
	data = data[1 : len(data)-1]

	if !strings.HasPrefix(string(data), "env[") || !strings.HasSuffix(string(data), "]") {
		return data, errors.New("Env key is not represented within data bytes array")
	}

	return data[4 : len(data)-1], nil
}
