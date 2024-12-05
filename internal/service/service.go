package service

import (
	"RoomBook/internal/models"
	"context"
)

type UserServ interface {
	Registration(ctx context.Context, user models.UserCreate) (int, error)
	Login(ctx context.Context, user models.UserLogin) (int, error)
	GetMe(ctx context.Context, id int) (*models.UserGet, error)
	Get(ctx context.Context, id int) (*models.UserGet, error)
	Delete(ctx context.Context, id int) error
	AddPhoto(ctx context.Context, adding models.AddPhoto) error
	BookRoom(ctx context.Context, data models.UserBookRoom) error
}

type UserChangeServ interface {
	PWD(ctx context.Context, user models.UserChangePWD) error
	Email(ctx context.Context, user models.UserChangeEmail) error
	Phone(ctx context.Context, user models.UserChangePhone) error
	UserData(ctx context.Context, user models.UserChange) error
}

type AdminServ interface {
	Create(ctx context.Context, admin models.AdminCreate) (int, error)
	Get(ctx context.Context, id int) (*models.Admin, error)
	Login(ctx context.Context, admin models.AdminLogin) (int, error)
	Delete(ctx context.Context, id int) error
}

type AdminChangeServ interface {
	PWD(ctx context.Context, admin models.AdminChangePWD) error
	Email(ctx context.Context, admin models.AdminChangeEmail) error
	Phone(ctx context.Context, admin models.AdminChangePhone) error
	AdminData(ctx context.Context, admin models.AdminChange) error
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
	Get(ctx context.Context, hotelID int) (*models.PhotoWithIDHotel, error)
	Delete(ctx context.Context, id []int) error
}
