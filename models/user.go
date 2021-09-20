package models

type User struct {
	UserID      int
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	OfficeID    string
	RoleID      int
	Latitude    float32
	Longitude   float32
	CurrentDate string
	AuthToken   string
}
