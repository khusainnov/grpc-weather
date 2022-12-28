package val

import (
	"fmt"
	"regexp"
)

var (
	isValidCity = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	l := len(value)
	if l < minLength || l > maxLength {
		return fmt.Errorf("must contain from %d-%d characters len", maxLength, maxLength)
	}

	return nil
}

func ValidateCity(value string) error {
	err := ValidateString(value, 2, 100)
	if err != nil {
		return err
	}
	if !isValidCity(value) {
		return fmt.Errorf("must contain only letters, digits or underscore")
	}
	return nil
}
