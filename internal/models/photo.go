package models

type PhotoAdd struct {
	HotelID int    `json:"hotel_id"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
}

type PhotoBase struct {
	ListID int `json:"list_id"`
	PhotoAdd
}

type Photo struct {
	ID int `json:"id"`
	PhotoBase
}
