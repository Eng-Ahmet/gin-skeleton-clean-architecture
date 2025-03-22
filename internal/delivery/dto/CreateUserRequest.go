package dto

// CreateUserRequest is a struct that defines the request body for creating a new user
type CreateUserRequest struct {
	FirstName string `json:"first_name"` 
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}