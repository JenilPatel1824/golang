package xyz

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is empty")
	}

	message := fmt.Sprintf("hello " + name)

	return message, nil
}
