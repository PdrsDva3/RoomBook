package service

import (
	"backend_roombook/internal/models"
	"context"
)

type UserServ interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, user models.UserLogin) (int, error)
	ChangePWD(ctx context.Context, user models.UserChangePWD) (int, error)
	Delete(ctx context.Context, id int) error
}
