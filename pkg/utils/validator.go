package utils

import (
    "regexp"
)

// check if the email is valid
func IsValidEmail(email string) bool {
    regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re := regexp.MustCompile(regex)
    return re.MatchString(email)
}

// check if the password is strong
// it should be at least 8 characters long
func IsValidPassword(password string) bool {
    return len(password) >= 8
}
