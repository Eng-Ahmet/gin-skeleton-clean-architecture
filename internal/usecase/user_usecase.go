// package usecase
package usecase

import (
	"errors"
	"hwai-api/internal/delivery/dto"
	"hwai-api/internal/domain"
	"hwai-api/internal/repository"
	"hwai-api/pkg/utils"
	"time"
)

// UserUsecase is a struct that defines the usecase of the user
// it contains the user repository
type UserUsecase struct {
	userRepo repository.UserRepository // user repository
}

// NewUserUsecase creates a new user usecase with the necessary dependencies
// it takes a user repository as a parameter and returns a new user usecase
func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// CreateUser creates a new user with the given name and email
// it takes the name and email as parameters and returns the created user or an error
func (u *UserUsecase) CreateUser(data dto.CreateUserRequest) (*domain.User, error) {
	// check if the email is already used
	existingUser, err := u.userRepo.GetUserByEmail(data.Email)

	// if the user exists and len(existingUser) > 0, return an error
	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// Hash password
	getHashPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := &domain.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Password:  getHashPassword,
		Email:     data.Email,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	// Save user in the repository
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID gets a user by the given ID
// it takes the ID as a parameter and returns the user or an error
func (u *UserUsecase) GetUserByID(id uint) (*domain.User, error) {
	// get the user by the given ID
	// get the user by the given ID using the user repository
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}
