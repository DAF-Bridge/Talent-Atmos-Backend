package utils

import (
	"strings"
)

func SeparateName(fullName string) (string, string) {
	// Trim any extra spaces and split the name by space
	nameParts := strings.Fields(strings.TrimSpace(fullName))

	// If there are no parts or only one part, return the full name as the first name and an empty last name
	if len(nameParts) < 1 {
		return fullName, ""
	}
	if len(nameParts) == 1 {
		return nameParts[0], ""
	}

	// Assume that everything after the first part is the last name
	firstName := nameParts[0]
	lastName := strings.Join(nameParts[1:], " ")

	return firstName, lastName
}