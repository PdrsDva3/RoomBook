package models

import (
	"github.com/guregu/null"
	"time"
)

type UserBase struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

type UserCreate struct {
	UserBase
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserGet struct {
	ID int `json:"id"`
	UserBase
	Photo null.String `json:"photo"`
}

type UserChange struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Photo   string `json:"photo"`
}

type UserChangePWD struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type UserChangeEmail struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserChangePhone struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
}

type UserBookRoom struct {
	ID          int       `json:"id"`
	IDRoom      int       `json:"id_room"`
	DayCheckIn  time.Time `json:"day_check_in"`
	DayCheckOut time.Time `json:"day_check_out"`
}

type AddPhoto struct {
	ID    int    `json:"id"`
	Photo string `json:"photo"`
}

type JWTPair struct {
	JWTToken string `json:"jwt_token"`
	Refresh  string `json:"refresh_token"`
}
