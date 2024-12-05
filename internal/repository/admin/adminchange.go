package admin

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoChange struct {
	db *sqlx.DB
}

func InitAdminChangeRepository(db *sqlx.DB) repository.AdminChangeRepo {
	return RepoChange{db: db}
}

func (r RepoChange) ChangePWD(ctx context.Context, admin models.AdminChangePWD) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE admins SET hashed_password=$2 WHERE id=$1;`, admin.ID, admin.NewPWD)
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

func (r RepoChange) ChangeEmail(ctx context.Context, admin models.AdminChangeEmail) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE admins SET photo=$2 WHERE id=$1;`, admin.ID, admin.Email)
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

func (r RepoChange) ChangePhone(ctx context.Context, admin models.AdminChangePhone) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE admins SET photo=$2 WHERE id=$1;`, admin.ID, admin.Phone)
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

func (r RepoChange) ChangeAdminData(ctx context.Context, admin models.AdminChange) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE admins SET name=$2, photo=$3 WHERE id=$1;`, admin.ID, admin.Name, admin.Photo)
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
