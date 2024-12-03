package models

type HotelCreate struct {
	Name    string   `json:"name"`
	Stars   int      `json:"stars"`
	Address string   `json:"address"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Links   []string `json:"links"`
	Lat     string   `json:"lat"`
	Lon     string   `json:"lon"`
}

type HotelBase struct {
	Name    string   `json:"name"`
	Stars   int      `json:"stars"`
	Rating  float64  `json:"rating"`
	Address string   `json:"address"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Links   []string `json:"links"`
	Photo   []Photo  `json:"hotel_photo"`
	Lat     string   `json:"lat"`
	Lon     string   `json:"lon"`
	Tags    []Tag    `json:"tags"`
}

type Hotel struct {
	ID int `json:"id"`
	HotelBase
}

type HotelAddPhoto struct {
	IDHotel int     `json:"id_hotel"`
	Photo   []Photo `json:"hotel_photo"`
}

type HotelDelPhoto struct {
	IDHotel int     `json:"id_hotel"`
	Photo   []Photo `json:"hotel_photo"`
}
type HotelAddLink struct {
	IDHotel int    `json:"id_hotel"`
	Link    string `json:"link"`
}

type HotelDelLink struct {
	IDHotel int    `json:"id_hotel"`
	Link    string `json:"link"`
}

type HotelChange struct {
	IDHotel int      `json:"id_hotel"`
	Name    string   `json:"name"`
	Stars   int      `json:"stars"`
	Address string   `json:"address"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Links   []string `json:"links"`
}

type HotelAdmin struct {
	IDHotel int `json:"id_hotel"`
	IDAdmin int `json:"id_admin"`
}

type HotelRating struct {
	IDHotel int     `json:"id_hotel"`
	IDUser  int     `json:"id_user"`
	Rating  float32 `json:"rating"`
}
