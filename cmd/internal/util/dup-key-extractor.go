package util

import (
	"fmt"
	"regexp"
)

func ExtractDupKey(errorMessage string) (string, error) {
	// Define a regular expression pattern to match the "dup key" information and capture only the key name
	re := regexp.MustCompile(`dup key: { (\w+): "[^"]+" }`)

	// Find submatches in the error message
	matches := re.FindStringSubmatch(errorMessage)

	// Check if the pattern is found
	if len(matches) == 2 {
		return matches[1], nil
	}

	// Return an error if the pattern is not found
	return "", fmt.Errorf("unable to extract dup key information from error message")
}
