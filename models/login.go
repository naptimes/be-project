package models

type Login struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	RoleId   string `json:"role_id"`
}
