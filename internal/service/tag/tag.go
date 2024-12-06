package tag

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/internal/service"
	"RoomBook/pkg/log"
	"context"
)

type Serv struct {
	Repo repository.TagRepo
	log  *log.Logs
}

func InitTagService(tagRepo repository.TagRepo, log *log.Logs) service.TagServ {
	return Serv{Repo: tagRepo, log: log}
}

func (s Serv) AddHotel(ctx context.Context, hotel models.TagHotel) error {
	err := s.Repo.AddHotel(ctx, hotel)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("add hotel")
	return nil
}

func (s Serv) AddRoom(ctx context.Context, room models.TagRoom) error {
	err := s.Repo.AddRoom(ctx, room)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("add room")
	return nil
}

func (s Serv) DeleteHotel(ctx context.Context, hotel models.TagHotel) error {
	err := s.Repo.DeleteHotel(ctx, hotel)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("delete hotel")
	return nil
}

func (s Serv) DeleteRoom(ctx context.Context, room models.TagRoom) error {
	err := s.Repo.DeleteRoom(ctx, room)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info("delete room")
	return nil
}
