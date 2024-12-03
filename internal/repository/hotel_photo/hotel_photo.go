package hotel_photo

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func InitPhotoRepository(db *sqlx.DB) repository.PhotoRepo {
	return Repo{db: db}
}

func (r Repo) Add(ctx context.Context, photos []models.PhotoAdd) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	for _, photo := range photos {
		_, err = transaction.ExecContext(ctx, `INSERT INTO photo_hotels (hotel_id, name, photo) VALUES ($1, $2, $3)`,
			photo.HotelID, photo.Name, photo.Photo)
		if err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return cerr.Err(cerr.Rollback, rbErr).Error()
			}
			return cerr.Err(cerr.Transaction, err).Error()
		}
	}

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}

	return nil
}

func (r Repo) Get(ctx context.Context, hotelID int) (*[]models.Photo, error) {
	var photos []models.Photo
	rows, err := r.db.QueryContext(ctx, `SELECT id, hotel_id, name, photo from photo_hotels WHERE hotel_id = $1;`, hotelID)
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
	if len(photos) == 0 {
		row := r.db.QueryRowContext(ctx, `SELECT id, hotel_id, name from photo_hotels WHERE id = $1;`, 0)
		var photo models.Photo
		err = row.Scan(&photo.ID, &photo.ListID, &photo.HotelID, &photo.Name, &photo.Photo)
		if err != nil {
			return nil, cerr.Err(cerr.InvalidCount, err).Error()
		}
		photos = append(photos, photo)
	}
	return &photos, nil
}

func (r Repo) Delete(ctx context.Context, ids []int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	for _, id := range ids {
		result, err := transaction.ExecContext(ctx, `DELETE FROM photo_hotels WHERE id=$1;`, id)
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
