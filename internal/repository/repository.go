package repository

import (
	"context"
	"roombook_backend/internal/models"
)

type UserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	GetPWDbyEmail(ctx context.Context, user string) (int, string, error)
	ChangePWD(ctx context.Context, user models.UserChangePWD) (int, error)
	Delete(ctx context.Context, id int) error
}
type AdminRepo interface {
	Create(ctx context.Context, admin models.AdminCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Admin, error)
	GetPWDbyEmail(ctx context.Context, admin string) (int, string, error)
	ChangePWD(ctx context.Context, admin models.AdminChangePWD) (int, error)
	Delete(ctx context.Context, id int) error
}
type HotelRepo interface {
	Create(ctx context.Context, hotel models.HotelBase) (int, error)
	Get(ctx context.Context, id int) (*models.Hotel, error)
	Delete(ctx context.Context, id int) error
}

type PhotoRepo interface {
	Add(ctx context.Context, photos []models.PhotoAdd) error
	Get(ctx context.Context, hotelID int) (*[]models.Photo, error)
	Delete(ctx context.Context, id []int) error
}
