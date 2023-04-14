package repository

import (
	"context"

	"github.com/Miguelburitica/go-rest-server/models"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	Close() error
}

var implementation UserRepository

func SetUserRepository(repo UserRepository) {
	implementation = repo
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}
