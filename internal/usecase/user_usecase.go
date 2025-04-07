package usecase

import "fiper-skeleton-clean-architecture/internal/domain"

type UserUsecase struct {
	// Add repository dependency here
}

func (u *UserUsecase) CreateUser(user *domain.User) error {
	// Business logic for creating a user
	return nil
}

func (u *UserUsecase) GetUserByID(id uint) (*domain.User, error) {
	// Business logic for retrieving a user
	return &domain.User{}, nil
}
