package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// hashPassword  hashes the password using bcrypt
// It returns the hashed password as a string and an error if any
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hashedPassword), err
}

// ComparePassword  compares the hashed password with the password
// It returns true if the password matches the hashed password, false otherwise
func ComparePassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
