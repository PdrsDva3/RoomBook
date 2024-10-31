package photo

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

func InitPhotoRepository(db *sqlx.DB) repository.PhotoRepo {
	return Repo{db: db}
}

func (repo Repo) Add(ctx context.Context, photos []models.PhotoAdd) error {
	var id int
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	row, err := repo.db.QueryContext(ctx, "SELECT MAX(list_id) FROM photos WHERE hotel_id = $1", photos[0].HotelID)
	if row == nil {
		id = 0
	} else {
		err = row.Scan(&id)

		if err != nil {
			return cerr.Err(cerr.Scan, err).Error()
		}
	}
	for _, photo := range photos {
		id++
		transaction.QueryRowContext(ctx, `INSERT INTO photos (list_id, hotel_id, name, photo) VALUES ($1, $2, $3)`,
			id, photo.HotelID, photo.Name, photo.Photo)
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (repo Repo) Get(ctx context.Context, id int) (*[]models.Photo, error) {
	var photos []models.Photo
	rows, err := repo.db.QueryContext(ctx, `SELECT id, list_id, hotel_id, name, photo from photos WHERE id = $1 ORDER BY list_id;`, id)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for rows.Next() {
		var photo models.Photo
		err = rows.Scan(&photo.ID, &photo.ListID, &photo.HotelID, &photo.Name, &photo.Photo)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		photos = append(photos, photo)
	}
	return &photos, nil
}

func (repo Repo) Delete(ctx context.Context, ids []int) error {
	transaction, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	for _, id := range ids {
		result, err := transaction.ExecContext(ctx, `DELETE FROM photos WHERE id=$1;`, id)
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
	}
	return nil
}
