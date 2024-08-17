package application

import (
	"github.com/Uno-count/Task-Management/internal/domain/models"
	"github.com/Uno-count/Task-Management/internal/domain/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}
