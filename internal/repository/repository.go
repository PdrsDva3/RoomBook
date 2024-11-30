package repository

import (
	"RoomBook/internal/models"
	"context"
)

type UserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.UserGet, error)
	GetPWDbyEmail(ctx context.Context, user string) (int, string, error)
	BookRoom(ctx context.Context, data models.UserBookRoom) error
	AddPhoto(ctx context.Context, adding models.AddPhoto) error
	Delete(ctx context.Context, id int) error
}

type UserChangeRepo interface {
	ChangePWD(ctx context.Context, user models.UserChangePWD) error
	ChangeEmail(ctx context.Context, user models.UserChangeEmail) error
	ChangePhone(ctx context.Context, user models.UserChangePhone) error
	ChangeUserData(ctx context.Context, user models.UserChange) error
}

type AdminRepo interface {
	Create(ctx context.Context, admin models.AdminCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Admin, error)
	GetPWDbyEmail(ctx context.Context, admin string) (int, string, error)
	ChangePWD(ctx context.Context, admin models.AdminChangePWD) (int, error)
	Delete(ctx context.Context, id int) error
}
type HotelRepo interface {
	Create(ctx context.Context, hotel models.HotelCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Hotel, error)
	Delete(ctx context.Context, id int) error
}

type PhotoRepo interface {
	Add(ctx context.Context, photos []models.PhotoAdd) error
	Get(ctx context.Context, hotelID int) (*[]models.Photo, error)
	Delete(ctx context.Context, id []int) error
}
