package hotel_photo

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/internal/service"
	"RoomBook/pkg/log"
	"context"
	"fmt"
)

type Serv struct {
	Repo repository.PhotoHotelRepo
	log  *log.Logs
}

func InitPhotoService(photoRepo repository.PhotoHotelRepo, log *log.Logs) service.PhotoServ {
	return Serv{Repo: photoRepo, log: log}
}

func (s Serv) Add(ctx context.Context, photos models.PhotoAddWithIDHotel) error {
	idHotel := photos.HotelID
	newPhotos := make([]models.PhotoAdd, len(photos.Photos))
	for i, photo := range photos.Photos {
		newPhotos[i] = models.PhotoAdd{
			HotelID:                idHotel,
			PhotoAddWithoutIDHotel: photo,
		}
	}
	err := s.Repo.Add(ctx, newPhotos)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("Add hotel_photo from %d", idHotel))
	return nil
}

func (s Serv) Get(ctx context.Context, hotelID int) (*[]models.Photo, error) {
	photos, err := s.Repo.Get(ctx, hotelID)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("Get hotel_photo from %d", hotelID))
	return photos, nil
}

func (s Serv) Delete(ctx context.Context, id []int) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("Delete photos count %d", len(id)))
	return nil
}
