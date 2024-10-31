package models

type HotelBase struct {
	Name    string   `json:"name"`
	Stars   int      `json:"stars"`
	Address string   `json:"address"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Links   []string `json:"links"`
	Photo   []Photo  `json:"photo"`
}

type Hotel struct {
	ID int `json:"id"`
	HotelBase
}

type HotelAddPhoto struct {
	IDHotel int     `json:"id_hotel"`
	Photo   []Photo `json:"photo"`
}

type HotelDelPhoto struct {
	IDHotel int     `json:"id_hotel"`
	Photo   []Photo `json:"photo"`
}
type HotelAddLink struct {
	IDHotel int    `json:"id_hotel"`
	Link    string `json:"link"`
}

type HotelDelLink struct {
	IDHotel int    `json:"id_hotel"`
	Link    string `json:"link"`
}
