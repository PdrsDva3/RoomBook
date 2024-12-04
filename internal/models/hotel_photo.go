package models

type PhotoAddWithoutIDHotel struct {
	Name  string `json:"name"`
	Photo string `json:"hotel_photo"`
}

type PhotoAddWithIDHotel struct {
	HotelID int `json:"hotel_id"`
	Photos  []PhotoAddWithoutIDHotel
}

type PhotoAdd struct {
	HotelID int `json:"hotel_id"`
	PhotoAddWithoutIDHotel
}

type PhotoBase struct {
	PhotoAdd
}

type Photo struct {
	ID int `json:"id"`
	PhotoBase
}

type PhotoDelete struct {
	ID []int `json:"id"`
}
