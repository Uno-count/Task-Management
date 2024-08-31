package user

import "github.com/Uno-count/Task-Management/internal/domain/models"

type Service interface {
	Register(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	LogIn(email, password string) (*models.User, error)
}
