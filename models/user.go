package models

type User struct {
	UserID         int    `json:"user_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Password       []byte `json:"-"`
	OfficeID       string `json:"office_id"`
	RoleID         int    `json:"role_id"`
	AuthToken      string `json:"auth_token"`
	ApprovalStatus bool   `json:"status" gorm:"column:approval_status"`
}
