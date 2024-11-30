package hotel

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/pkg/cerr"
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func InitHotelRepository(db *sqlx.DB) repository.HotelRepo {
	return Repo{db: db}
}

func (r Repo) Create(ctx context.Context, hotel models.HotelCreate) (int, error) {
	var id int
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	linksJson, err := json.Marshal(hotel.Links)
	if err != nil {
		return 0, cerr.Err(cerr.JSON, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO hotels (name, stars, address, email, phone, links) VALUES ($1, $2, $3, $4, $5, $6) returning id;`,
		hotel.Name, hotel.Stars, hotel.Address, hotel.Email, hotel.Phone, linksJson)

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

func (r Repo) Get(ctx context.Context, id int) (*models.Hotel, error) {
	var hotel models.Hotel
	var linksJSON []byte
	row := r.db.QueryRowContext(ctx, `SELECT name, stars, address, email, phone, links from hotels WHERE id = $1;`, id)
	err := row.Scan(&hotel.Name, &hotel.Stars, &hotel.Address, &hotel.Email, &hotel.Phone, &linksJSON)
	if err != nil {
		return nil, cerr.Err(cerr.InvalidEmail, err).Error()
	}
	if err = json.Unmarshal(linksJSON, &hotel.Links); err != nil {
		return nil, cerr.Err(cerr.JSON, err).Error()
	}
	hotel.ID = id
	rows, err := r.db.QueryContext(ctx, `SELECT id, list_id, hotel_id, name, photo from photos WHERE hotel_id = $1 ORDER BY list_id;`, id)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for rows.Next() {
		var photo models.Photo
		err = rows.Scan(&photo.ID, &photo.ListID, &photo.HotelID, &photo.Name, &photo.Photo)
		if err != nil {
			return nil, cerr.Err(cerr.InvalidCount, err).Error()
		}
		hotel.Photo = append(hotel.Photo, photo)
	}

	if len(hotel.Photo) == 0 {
		row = r.db.QueryRowContext(ctx, `SELECT id, list_id, hotel_id, name, photo from photos WHERE id = $1 ORDER BY list_id;`, 0)
		var photo models.Photo
		err = row.Scan(&photo.ID, &photo.ListID, &photo.HotelID, &photo.Name, &photo.Photo)
		if err != nil {
			return nil, cerr.Err(cerr.InvalidCount, err).Error()
		}
		hotel.Photo = append(hotel.Photo, photo)
	}
	return &hotel, nil
}

func (r Repo) Delete(ctx context.Context, id int) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `DELETE FROM hotels WHERE id=$1;`, id)
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
	result, err = transaction.ExecContext(ctx, `DELETE FROM photos WHERE hotel_id=$1;`, id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.ExecContext, err).Error()
	}
	count, err = result.RowsAffected()
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Rows, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}

	return nil
}
