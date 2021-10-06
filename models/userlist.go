package models

type Users struct {
	UserID         int    `json:"user_id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	ApprovalStatus bool   `json:"status" gorm:"column:approval_status"`
	Absen          int    `json:"absen"`
}
