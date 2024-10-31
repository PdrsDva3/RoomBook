package admin

import (
	"backend_roombook/internal/models"
	"backend_roombook/internal/repository"
	"backend_roombook/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func InitAdminRepository(db *sqlx.DB) repository.AdminRepo {
	return Repo{db: db}
}

func (repo Repo) Create(ctx context.Context, admin models.AdminCreate) (int, error) {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO admin (name, email, phone, hashed_password, photo) VALUES ($1, $2, $3, $4, $5) returning id;`,
		admin.Name, admin.Email, admin.Phone, admin.PWD, admin.Photo)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Scan, err).Error()
	}
	if err := transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Commit, err).Error()
	}
	return id, nil
}

func (repo Repo) Get(ctx context.Context, id int) (*models.Admin, error) {
	var admin models.Admin
	row := repo.db.QueryRowContext(ctx, `SELECT name, email, phone, photo from admin WHERE id = $1;`, id)
	err := row.Scan(&admin.Name, &admin.Email, &admin.Phone, &admin.Photo)
	if err != nil {
		return nil, cerr.Err(cerr.Scan, err).Error()
	}
	admin.ID = id
	return &admin, nil
}

func (repo Repo) GetPWDbyEmail(ctx context.Context, admin string) (int, string, error) {
	var pwd string
	var id int
	rows := repo.db.QueryRowContext(ctx, `SELECT id, hashed_password from admin WHERE email = $1;`, admin)
	err := rows.Scan(&id, &pwd)
	if err != nil {
		return 0, "", cerr.Err(cerr.Scan, err).Error()
	}
	return id, pwd, nil
}

func (repo Repo) ChangePWD(ctx context.Context, admin models.AdminChangePWD) (int, error) {
	transaction, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE admin SET hashed_password=$2 WHERE id=$1;`, admin.ID, admin.PWD)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Rows, err).Error()
	}

	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.NoOneRow, err).Error()
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Commit, err).Error()
	}
	return admin.ID, nil
}

func (repo Repo) Delete(ctx context.Context, id int) error {
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM admin WHERE id=$1;`, id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err := result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}
	if count != 1 {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.NoOneRow, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}
