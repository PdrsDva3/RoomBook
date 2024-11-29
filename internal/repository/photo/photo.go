package photo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"roombook_backend/internal/models"
	"roombook_backend/internal/repository"
	"roombook_backend/pkg/cerr"
)

type Repo struct {
	db *sqlx.DB
}

func InitPhotoRepository(db *sqlx.DB) repository.PhotoRepo {
	return Repo{db: db}
}

func (r Repo) Add(ctx context.Context, photos []models.PhotoAdd) error {
	var id, count int
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}

	row := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM photos WHERE hotel_id = $1", photos[0].HotelID)
	if err = row.Scan(&count); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Scan, err).Error()
	}

	if count == 0 {
		id = 0
	} else {
		row = r.db.QueryRowContext(ctx, "SELECT MAX(list_id) FROM photos WHERE hotel_id = $1", photos[0].HotelID)
		if err = row.Scan(&id); err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return cerr.Err(cerr.Rollback, rbErr).Error()
			}
			return cerr.Err(cerr.Scan, err).Error()
		}
	}

	// Вставляем фотографии
	for _, photo := range photos {
		id++ // Увеличиваем list_id для каждой фотографии
		_, err = transaction.ExecContext(ctx, `INSERT INTO photos (list_id, hotel_id, name, photo) VALUES ($1, $2, $3, $4)`,
			id, photo.HotelID, photo.Name, photo.Photo)
		if err != nil {
			if rbErr := transaction.Rollback(); rbErr != nil {
				return cerr.Err(cerr.Rollback, rbErr).Error()
			}
			return cerr.Err(cerr.Transaction, err).Error()
		}
	}

	// Завершаем транзакцию
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
	rows, err := r.db.QueryContext(ctx, `SELECT id, list_id, hotel_id, name, photo from photos WHERE hotel_id = $1 ORDER BY list_id;`, hotelID)
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
		row := r.db.QueryRowContext(ctx, `SELECT id, list_id, hotel_id, name, photo from photos WHERE id = $1 ORDER BY list_id;`, 0)
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
