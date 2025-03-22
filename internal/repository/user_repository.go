// package repository
package repository

import (
	domain "hwai-api/internal/domain" // import the domain package
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB // DB is the database connection
}

// NewUserRepository  creates a new user repository with the necessary dependencies
// it takes a database connection as a parameter and returns a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser takes a user as a parameter and creates a new user
// it returns an error if the operation fails
func (r *UserRepository) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

// GetUserByID gets a user by the given ID
// it takes the ID as a parameter and returns the user or an error
func (r *UserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	result := r.DB.First(&user, id) // find the user by ID
	return &user, result.Error
}

// GetUserByEmail gets a user by the given email
// it takes the email as a parameter and returns the user or an error
func (r *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.DB.Where("email = ?", email).First(&user) // find the user by email
	return &user, result.Error
}