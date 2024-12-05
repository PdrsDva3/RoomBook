package user

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/internal/service"
	"RoomBook/pkg/auth"
	"RoomBook/pkg/cerr"
	"RoomBook/pkg/database/cached"
	"RoomBook/pkg/log"
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type ServUser struct {
	Repo    repository.UserRepo
	log     *log.Logs
	jwt     auth.JWTUtil
	session cached.Session
}

func InitUserService(userRepo repository.UserRepo, log *log.Logs, jwt auth.JWTUtil, session cached.Session) service.UserServ {
	return ServUser{Repo: userRepo, log: log, jwt: jwt, session: session}
}

func (s ServUser) Registration(ctx context.Context, user models.UserCreate) (*models.JWTPair, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	newUser := models.UserCreate{
		UserBase: user.UserBase,
		Password: string(hashedPassword),
	}

	id, err := s.Repo.Create(ctx, newUser)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	token := s.jwt.CreateToken(id)
	if token == "" {
		return nil, errors.New("token is empty")
	}

	uuidBytes, err := uuid.NewV4()
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	sd := cached.SessionData{
		RefreshToken:   uuidBytes.String(),
		LoginTimeStamp: time.Now(),
	}

	err = s.session.Set(ctx, id, sd)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}

	res := models.JWTPair{
		JWTToken: token,
		Refresh:  uuidBytes.String(),
	}

	s.log.Info(fmt.Sprintf("create user %v", id))

	return &res, nil
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

func (s ServUser) RefreshToken(ctx context.Context, refreshToken string) (*models.JWTPair, error) {
	
}
