package service

import (
	"RoomBook/internal/models"
	"context"
)

type UserServ interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	Login(ctx context.Context, user models.UserLogin) (int, error)
	ChangePWD(ctx context.Context, user models.UserChangePWD) (int, error)
	Delete(ctx context.Context, id int) error
}

type AdminServ interface {
	Create(ctx context.Context, admin models.AdminCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Admin, error)
	Login(ctx context.Context, admin models.AdminLogin) (int, error)
	ChangePWD(ctx context.Context, admin models.AdminChangePWD) (int, error)
	Delete(ctx context.Context, id int) error
}

type HotelServ interface {
	Create(ctx context.Context, hotel models.HotelCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Hotel, error)
	GetAll(ctx context.Context) ([]models.Hotel, error)
	Change(ctx context.Context, hotel models.HotelChange) error
	Admin(ctx context.Context, admin models.HotelAdmin) error
	Rating(ctx context.Context, hotel models.HotelRating) error
	Delete(ctx context.Context, id int) error
}

type PhotoServ interface {
	Add(ctx context.Context, photos models.PhotoAddWithIDHotel) error
	Get(ctx context.Context, hotelID int) (*[]models.Photo, error)
	Delete(ctx context.Context, id []int) error
}
