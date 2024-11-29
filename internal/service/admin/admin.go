package admin

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"roombook_backend/internal/models"
	"roombook_backend/internal/repository"
	"roombook_backend/internal/service"
	"roombook_backend/pkg/cerr"
	"roombook_backend/pkg/log"
)

type Serv struct {
	Repo repository.AdminRepo
	log  *log.Logs
}

func InitAdminService(adminRepo repository.AdminRepo, log *log.Logs) service.AdminServ {
	return Serv{Repo: adminRepo, log: log}
}

func (s Serv) Create(ctx context.Context, admin models.AdminCreate) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.PWD), 10)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	newAdmin := models.AdminCreate{
		AdminBase: admin.AdminBase,
		PWD:       string(hashedPassword),
	}
	id, err := s.Repo.Create(ctx, newAdmin)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("create admin %v", id))
	return id, nil
}

func (s Serv) Get(ctx context.Context, id int) (*models.Admin, error) {
	admin, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get admin %v", id))
	return admin, nil
}

func (s Serv) Login(ctx context.Context, admin models.AdminLogin) (int, error) {
	id, pwd, err := s.Repo.GetPWDbyEmail(ctx, admin.Email)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(admin.PWD))
	if err != nil {
		s.log.Error(cerr.Err(cerr.InvalidPWD, err).Str())
		return 0, cerr.Err(cerr.InvalidPWD, err).Error()
	}
	s.log.Info(fmt.Sprintf("login admin %v", id))
	return id, nil
}

func (s Serv) ChangePWD(ctx context.Context, admin models.AdminChangePWD) (int, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.NewPWD), 10)
	if err != nil {
		s.log.Error(cerr.Err(cerr.Hash, err).Str())
		return 0, cerr.Err(cerr.Hash, err).Error()
	}
	newPWD := models.AdminChangePWD{
		ID:     admin.ID,
		NewPWD: string(hash),
	}
	id, err := s.Repo.ChangePWD(ctx, newPWD)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("change pwd admin %v", id))
	return id, nil
}

func (s Serv) Delete(ctx context.Context, id int) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("delete admin %v", id))
	return nil
}
