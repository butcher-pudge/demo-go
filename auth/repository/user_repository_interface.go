package repository

import (
	"context"
	"go-learning-demo/models"
)

type IUserRepository interface {
	CreateUser(context context.Context, user *models.User) error
	GetUser(context context.Context, username, password string) (*models.User, error)
}
