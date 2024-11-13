package user

import (
	"backend_roombook/internal/models"
	"backend_roombook/internal/repository"
	"backend_roombook/internal/service"
	"backend_roombook/pkg/cerr"
	"backend_roombook/pkg/log"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Serv struct {
	Repo repository.UserRepo
	log  *log.Logs
}

func InitUserService(userRepo repository.UserRepo, log *log.Logs) service.UserServ {
	return &Serv{Repo: userRepo, log: log}
}

func (s Serv) Create(ctx context.Context, user models.UserCreate) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PWD), 10)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	newUser := models.UserCreate{
		UserBase: user.UserBase,
		PWD:      string(hashedPassword),
	}
	id, err := s.Repo.Create(ctx, newUser)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("create user %v", id))
	return id, nil
}

func (s Serv) Get(ctx context.Context, id int) (*models.User, error) {
	user, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get user %v", id))
	return user, nil
}

func (s Serv) Login(ctx context.Context, user models.UserLogin) (int, error) {
	id, pwd, err := s.Repo.GetPWDbyEmail(ctx, user.Email)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(user.PWD))
	if err != nil {
		s.log.Error(cerr.Err(cerr.InvalidPWD, err).Str())
		return 0, cerr.Err(cerr.InvalidPWD, err).Error()
	}
	s.log.Info(fmt.Sprintf("login user %v", id))
	return id, nil
}

func (s Serv) ChangePWD(ctx context.Context, user models.UserChangePWD) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.NewPWD), 10)
	if err != nil {
		s.log.Error(cerr.Err(cerr.Hash, err).Str())
		return 0, cerr.Err(cerr.Hash, err).Error()
	}
	newPWD := models.UserChangePWD{
		ID:     user.ID,
		NewPWD: string(hash),
	}
	id, err := s.Repo.ChangePWD(ctx, newPWD)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("change pwd user %v", id))
	return id, nil
}

func (s Serv) Delete(ctx context.Context, id int) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("delete user %v", id))
	return nil
}
