package pbclient

import (
	"errors"

	"github.com/Uno-count/Task-Management/internal/domain/models"
	pbModels "github.com/pocketbase/pocketbase/models"
)

type UserRepository struct {
	client *Client
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{client: client}
}

func (r *UserRepository) Create(user *models.User) error {
	collection, err := r.client.App.Dao().FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	record := pbModels.NewRecord(collection)
	record.Set("username", user.Username)
	record.Set("email", user.Email)
	record.SetPassword(user.Password)

	if err := r.client.App.Dao().SaveRecord(record); err != nil {
		return err
	}

	user.ID = record.Id
	user.CreatedAt = record.Created.Time()
	user.UpdatedAt = record.Updated.Time()

	return nil
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	record, err := r.client.App.Dao().FindAuthRecordByEmail("users", email)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:        record.Id,
		Username:  record.GetString("username"),
		Email:     record.Email(),
		CreatedAt: record.Created.Time(),
		UpdatedAt: record.Updated.Time(),
	}

	return user, nil
}

func (r *UserRepository) ValidateCredentials(email, password string) (*models.User, error) {
	record, err := r.client.App.Dao().FindAuthRecordByEmail("users", email)
	if err != nil {
		return nil, err
	}

	if !record.ValidatePassword(password) {
		return nil, errors.New("invalid password, please check your password")
	}

	user := &models.User{
		ID:       record.Id,
		Username: record.GetString("username"),
		Email:    record.Email(),
	}

	return user, nil
}
