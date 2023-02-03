package internal

import (
	"os"
)

func LoadToken(file string) (string, error) {

	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
