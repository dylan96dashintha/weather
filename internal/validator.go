package internal

import (
	"errors"
	"regexp"
	"unicode"
)

func stringValidator(stringName string) (err error) {

	// Method 1: Using regular expression
	match, _ := regexp.MatchString("^[a-zA-Z ]+$", stringName)
	if match {
		return nil
	}

	// Method 2: Using unicode.IsLetter
	for _, char := range stringName {
		if !unicode.IsLetter(char) {
			return errors.New("invalid input")
		}
	}

	return nil
}
