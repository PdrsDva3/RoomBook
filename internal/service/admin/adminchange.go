package admin

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

type ServChange struct {
	Repo repository.AdminChangeRepo
	log  *log.Logs
}

func InitAdminChangeService(adminChangeRepo repository.AdminChangeRepo, log *log.Logs) service.AdminChangeServ {
	return ServChange{Repo: adminChangeRepo, log: log}
}

func (s ServChange) PWD(ctx context.Context, admin models.AdminChangePWD) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.NewPWD), 10)
	if err != nil {
		s.log.Error(cerr.Err(cerr.Hash, err).Str())
		return cerr.Err(cerr.Hash, err).Error()
	}
	newPWD := models.AdminChangePWD{
		ID:     admin.ID,
		NewPWD: string(hash),
	}

	err = s.Repo.ChangePWD(ctx, newPWD)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("change pwd admin"))
	return nil
}

func (s ServChange) Email(ctx context.Context, admin models.AdminChangeEmail) error {
	err := s.Repo.ChangeEmail(ctx, admin)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	return nil
}

func (s ServChange) Phone(ctx context.Context, admin models.AdminChangePhone) error {
	err := s.Repo.ChangePhone(ctx, admin)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	return nil
}

func (s ServChange) AdminData(ctx context.Context, admin models.AdminChange) error {
	err := s.Repo.ChangeAdminData(ctx, admin)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	return nil
}
