package models

type User struct {
	UserID    int    `json:"user_id"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	OfficeID  string `json:"office_id"`
	RoleID    int    `json:"role_id"`
	AuthToken string `json:"auth_token"`
}
