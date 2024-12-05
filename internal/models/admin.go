package models

type AdminBase struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Photo string `json:"photo"`
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

type AdminChange struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}

type AdminChangeEmail struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type AdminChangePhone struct {
	ID    int    `json:"id"`
	Phone string `json:"phone"`
}
