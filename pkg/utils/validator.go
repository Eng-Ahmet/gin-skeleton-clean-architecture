package utils

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
	// Example: Password must be at least 8 characters long
	return len(password) >= 8
}
