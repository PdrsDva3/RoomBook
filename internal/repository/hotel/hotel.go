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

func (r Repo) GetAll(ctx context.Context) ([]models.Hotel, error) {
	var hotels []models.Hotel
	var hotel models.Hotel
	var linksJSON []byte
	row, err := r.db.QueryContext(ctx, `SELECT id, name, rating, stars, address, email, phone, links, lat, lon from hotels;`)
	if err != nil {
		return nil, cerr.Err(cerr.Execution, err).Error()
	}
	for row.Next() {
		err := row.Scan(&hotel.ID, &hotel.Name, &hotel.Rating, &hotel.Stars, &hotel.Address, &hotel.Email, &hotel.Phone, &linksJSON, &hotel.Lat, &hotel.Lon)
		if err != nil {
			return nil, cerr.Err(cerr.InvalidEmail, err).Error()
		}
		if err = json.Unmarshal(linksJSON, &hotel.Links); err != nil {
			return nil, cerr.Err(cerr.JSON, err).Error()
		}
		rows, err := r.db.QueryContext(ctx, `SELECT id, hotel_id, name, photo from photo_hotels WHERE hotel_id = $1;`, hotel.ID)
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
			roww := r.db.QueryRowContext(ctx, `SELECT id, hotel_id, name from photo_hotels WHERE id = $1;`, 0)
			var photo models.Photo
			err = roww.Scan(&photo.ID, &photo.ListID, &photo.HotelID, &photo.Name, &photo.Photo)
			if err != nil {
				return nil, cerr.Err(cerr.InvalidCount, err).Error()
			}
			hotel.Photo = append(hotel.Photo, photo)
		}
		hotels = append(hotels, hotel)
	}
	return hotels, nil
}

func (r Repo) Change(ctx context.Context, hotel models.HotelChange) error {
	transaction, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	linksJson, err := json.Marshal(hotel.Links)
	if err != nil {
		return cerr.Err(cerr.JSON, err).Error()
	}
	result, err := transaction.ExecContext(ctx, `UPDATE hotels SET name=$2, stars=$3, phone=$4, links=$5, address=$6, email=$7 WHERE id=$1;`, hotel.IDHotel, hotel.Name, hotel.Stars, hotel.Phone, linksJson, hotel.Address, hotel.Email)
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

func (r Repo) Admin(ctx context.Context, admin models.HotelAdmin) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	_ = transaction.QueryRowContext(ctx, `INSERT INTO admin_hotel (id_admin, id_hotel) VALUES ($1, $2);`,
		admin.IDAdmin, admin.IDHotel)

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}
	return nil
}

func (r Repo) Rating(ctx context.Context, rating models.HotelRating) error {
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return cerr.Err(cerr.Transaction, err).Error()
	}
	_ = transaction.QueryRowContext(ctx, `INSERT INTO ratings (id_user, id_hotel, rating) VALUES ($1, $2, $3);`,
		rating.IDUser, rating.IDHotel, rating.Rating)

	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return cerr.Err(cerr.Commit, err).Error()
	}

	rows, err := r.db.QueryContext(ctx, `SELECT rating from ratings WHERE id_hotel = $1;`, rating.IDHotel)
	if err != nil {
		return cerr.Err(cerr.Rows, err).Error()
	}
	cnt := 0.0
	rat := 0.0
	rating := 0.0
	for rows.Next() {
		err = rows.Scan(&photo.ID, &photo.ListID, &photo.HotelID, &photo.Name, &photo.Photo)
		if err != nil {
			return nil, cerr.Err(cerr.InvalidCount, err).Error()
		}
		hotel.Photo = append(hotel.Photo, photo)
	}
	return nil
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
	row := transaction.QueryRowContext(ctx, `INSERT INTO hotels (name, stars, address, email, phone, links, lat, lon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning id;`,
		hotel.Name, hotel.Stars, hotel.Address, hotel.Email, hotel.Phone, linksJson, hotel.Lat, hotel.Lon)

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
	row := r.db.QueryRowContext(ctx, `SELECT name, rating, stars, address, email, phone, links, lat, lon from hotels WHERE id = $1;`, id)
	err := row.Scan(&hotel.Name, &hotel.Rating, &hotel.Stars, &hotel.Address, &hotel.Email, &hotel.Phone, &linksJSON, &hotel.Lat, &hotel.Lon)
	if err != nil {
		return nil, cerr.Err(cerr.InvalidEmail, err).Error()
	}
	if err = json.Unmarshal(linksJSON, &hotel.Links); err != nil {
		return nil, cerr.Err(cerr.JSON, err).Error()
	}
	hotel.ID = id
	rows, err := r.db.QueryContext(ctx, `SELECT id, hotel_id, name, photo from photo_hotels WHERE hotel_id = $1;`, id)
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
		row = r.db.QueryRowContext(ctx, `SELECT id, hotel_id, name from photo_hotels WHERE id = $1;`, 0)
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
	result, err = transaction.ExecContext(ctx, `DELETE FROM photo_hotels WHERE hotel_id=$1;`, id)
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
