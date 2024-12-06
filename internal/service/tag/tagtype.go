package tag

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/internal/service"
	"RoomBook/pkg/log"
	"context"
)

type ServType struct {
	Repo repository.TagTypeRepo
	log  *log.Logs
}

func InitTagTypeService(typeRepo repository.TagTypeRepo, log *log.Logs) service.TagTypeServ {
	return ServType{Repo: typeRepo, log: log}
}

func (s ServType) CreateType(ctx context.Context, types models.TypeCreate) (int, error) {
	id, err := s.Repo.CreateType(ctx, types)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info("create type")
	return id, nil
}

func (s ServType) CreateTag(ctx context.Context, tag models.TagCreate) (*models.Tag, error) {
	id, err := s.Repo.CreateTag(ctx, tag)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info("create tag")
	return id, nil
}

func (s ServType) Tags(ctx context.Context) ([]models.Tag, error) {
	tags, err := s.Repo.Tags(ctx)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info("tags")
	return tags, nil
}

func (s ServType) Types(ctx context.Context) ([]models.TypeBase, error) {
	types, err := s.Repo.Types(ctx)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info("types")
	return types, nil
}

func (s ServType) TagsType(ctx context.Context, idType int) (*models.TagsType, error) {
	types, err := s.Repo.TagsType(ctx, idType)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info("tags type")
	return types, nil
}
