package user

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/internal/service"
	"RoomBook/pkg/cerr"
	"RoomBook/pkg/log"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type ServUser struct {
	Repo repository.UserRepo
	log  *log.Logs
}

func InitUserService(userRepo repository.UserRepo, log *log.Logs) service.UserServ {
	return ServUser{Repo: userRepo, log: log}
}

func (s ServUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	newUser := models.UserCreate{
		UserBase: user.UserBase,
		Password: string(hashedPassword),
	}
	id, err := s.Repo.Create(ctx, newUser)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("create user %v", id))
	return id, nil
}

func (s ServUser) GetMe(ctx context.Context, id int) (*models.UserGet, error) {
	user, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get user %v", id))
	return user, nil
}

func (s ServUser) Get(ctx context.Context, id int) (*models.UserGet, error) {
	user, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get user %v", id))
	return user, nil
}

func (s ServUser) Login(ctx context.Context, user models.UserLogin) (int, error) {
	id, pwd, err := s.Repo.GetPWDbyEmail(ctx, user.Email)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(user.Password))
	if err != nil {
		s.log.Error(cerr.Err(cerr.InvalidPWD, err).Str())
		return 0, cerr.Err(cerr.InvalidPWD, err).Error()
	}
	s.log.Info(fmt.Sprintf("login user %v", id))
	return id, nil
}

func (s ServUser) Delete(ctx context.Context, id int) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("delete user %v", id))
	return nil
}

func (s ServUser) Registration(ctx context.Context, user models.UserCreate) (int, error) {
	id, err := s.Repo.Create(ctx, user)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("registred user %v", id))
	return id, nil
}

func (s ServUser) AddPhoto(ctx context.Context, adding models.AddPhoto) error {
	err := s.Repo.AddPhoto(ctx, adding)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("add photo"))
	return nil
}

func (s ServUser) BookRoom(ctx context.Context, data models.UserBookRoom) error {
	err := s.Repo.BookRoom(ctx, data)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("room book =)"))
	return nil
}
