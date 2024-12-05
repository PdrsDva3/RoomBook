package user

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoUserChange struct {
	db *sqlx.DB
}

func InitUserChangeRepository(db *sqlx.DB) repository.UserChangeRepo {
	return RepoUserChange{db: db}
}

func (r RepoUserChange) ChangePWD(ctx context.Context, user models.UserChangePWD) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE users SET hashed_pwd=$1 WHERE id=$2;`, user.Password, user.ID)
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.NoOneRow, err).Error()
	}

	if err = tr.Commit(); err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r RepoUserChange) ChangeEmail(ctx context.Context, user models.UserChangeEmail) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE users SET email=$1 WHERE id=$2;`, user.Email, user.ID)
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.NoOneRow, err).Error()
	}

	if err = tr.Commit(); err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r RepoUserChange) ChangePhone(ctx context.Context, user models.UserChangePhone) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE users SET phone=$1 WHERE id=$2;`, user.Phone, user.ID)
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.NoOneRow, err).Error()
	}

	if err = tr.Commit(); err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r RepoUserChange) ChangeUserData(ctx context.Context, user models.UserChange) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE users SET name=$2, surname=$3, photo=$4 WHERE id=$1;`, user.ID, user.Name, user.Surname, user.Photo)
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.NoOneRow, err).Error()
	}

	if err = tr.Commit(); err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}
