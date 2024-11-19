package utils

import (
	"encoding/json"
	"errors"
)

func ValidateJSON(input string) error {
	var js map[string]interface{}
	if err := json.Unmarshal([]byte(input), &js); err != nil {
		return errors.New("invalid JSON format")
	}
	return nil
}
