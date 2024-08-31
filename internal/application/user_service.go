package application

import (
	"errors"

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

func (s *UserService) Login(email, password string) (*models.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email and Password are required")
	}

	user, err := s.repo.ValidateCredentials(email, password)
	if err != nil {
		return nil, errors.New("invalid credentials, please check your email and password")
	}

	return user, nil
}
