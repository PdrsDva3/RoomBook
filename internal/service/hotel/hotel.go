package hotel

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/internal/service"
	"RoomBook/pkg/log"
	"context"
	"fmt"
)

type Serv struct {
	Repo repository.HotelRepo
	log  *log.Logs
}

func InitHotelService(hotelRepo repository.HotelRepo, log *log.Logs) service.HotelServ {
	return Serv{Repo: hotelRepo, log: log}
}

func (s Serv) GetAll(ctx context.Context) ([]models.Hotel, error) {
	hotels, err := s.Repo.GetAll(ctx)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info("Get hotels")
	return hotels, nil
}

func (s Serv) Change(ctx context.Context, hotel models.HotelChange) error {
	err := s.Repo.Change(ctx, hotel)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("Change hotel with id %d", hotel.IDHotel))
	return nil
}

func (s Serv) Admin(ctx context.Context, admin models.HotelAdmin) error {
	err := s.Repo.Admin(ctx, admin)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("Add admin hotel with HotelID %d", admin.IDHotel))
	return nil
}

func (s Serv) Rating(ctx context.Context, hotel models.HotelRating) error {
	err := s.Repo.Rating(ctx, hotel)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("Add Rating hotel with id %d", hotel.IDHotel))
	return nil
}

func (s Serv) Create(ctx context.Context, hotel models.HotelCreate) (int, error) {
	id, err := s.Repo.Create(ctx, hotel)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("Created hotel with id %d", id))
	return id, nil
}

func (s Serv) Get(ctx context.Context, id int) (*models.Hotel, error) {
	hotel, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("Get hotel with id %d", id))
	return hotel, nil
}

func (s Serv) Delete(ctx context.Context, id int) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("Delete hotel with id %d", id))
	return nil
}
