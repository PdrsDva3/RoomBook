package hotel

import (
	"context"
	"fmt"
	"roombook_backend/internal/models"
	"roombook_backend/internal/repository"
	"roombook_backend/internal/service"
	"roombook_backend/pkg/log"
)

type Serv struct {
	Repo repository.HotelRepo
	log  *log.Logs
}

func InitHotelService(hotelRepo repository.HotelRepo, log *log.Logs) service.HotelServ {
	return Serv{Repo: hotelRepo, log: log}
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
