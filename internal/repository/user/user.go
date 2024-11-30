package user

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	db *sqlx.DB
}

func InitUserRepository(db *sqlx.DB) repository.UserRepo {
	return RepoUser{db: db}
}

func (r RepoUser) Create(ctx context.Context, user models.UserCreate) (int, error) {
	var id int
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO users (name, sur_name, email, hashed_password) VALUES ($1, $2, $3, $4) returning id;`,
		user.Name, user.Surname, user.Email, user.Password)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Scan, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Commit, err).Error()
	}
	return id, nil
}

func (r RepoUser) Get(ctx context.Context, id int) (*models.UserGet, error) {
	var user models.UserGet
	row := r.db.QueryRowContext(ctx, `SELECT name, sur_name, email from users WHERE id = $1;`, id)
	err := row.Scan(&user.Name, &user.Surname, &user.Email)
	if err != nil {
		return nil, cerr.Err(cerr.Scan, err).Error()
	}
	user.ID = id
	return &user, nil
}

func (r RepoUser) ChangePWD(ctx context.Context, user models.UserChangePWD) (int, error) {
	transaction, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE users SET hashed_password=$2 WHERE id=$1;`, user.ID, user.Password)
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
	return user.ID, nil
}

func (r RepoUser) Delete(ctx context.Context, id int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM users WHERE id=$1;`, id)
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

func (r RepoUser) AddPhoto(ctx context.Context, adding models.AddPhoto) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	result, err := tr.ExecContext(ctx, `UPDATE users SET photo=$2 WHERE id=$1;`, adding.ID, adding.Photo)
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

func (r RepoUser) BookRoom(ctx context.Context, data models.UserBookRoom) error {
	tr, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	_ = tr.QueryRowContext(ctx, `INSERT INTO users_rooms (id, id_room, day_checkin, day_checkout) VALUES ($1, $2, $3, $4)`,
		data.ID, data.IDRoom, data.DayCheckIn, data.DayCheckOut)

	if err = tr.Commit(); err != nil {
		if rbErr := tr.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r RepoUser) GetPWDbyEmail(ctx context.Context, user string) (int, string, error) {
	var pwd string
	var id int
	rows := r.db.QueryRowContext(ctx, `SELECT id,  hashed_password from users WHERE email = $1;`, user)
	err := rows.Scan(&id, &pwd)
	if err != nil {
		return 0, "", cerr.Err(cerr.Scan, err).Error()
	}
	return id, pwd, nil
}
