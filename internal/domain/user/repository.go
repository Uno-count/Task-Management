package user

import "github.com/Uno-count/Task-Management/internal/domain/models"

type Repository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	ValidateCredentials(email, password string) (*models.User, error)
}
