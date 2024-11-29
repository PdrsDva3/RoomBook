package models

type AdminBase struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"hotel_photo"`
}

type Admin struct {
	ID int `json:"id"`
	AdminBase
}

type AdminCreate struct {
	AdminBase
	PWD string `json:"password"`
}

type AdminLogin struct {
	Email string `json:"email"`
	PWD   string `json:"password"`
}

type AdminChangePWD struct {
	ID     int    `json:"id"`
	NewPWD string `json:"password"`
}
