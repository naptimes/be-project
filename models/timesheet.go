package models

type Timesheet struct {
	CurrentDate      string `json:"current_date" gorm:"column:dates"`
	CheckIn          string `json:"checkin" gorm:"column:checkin"`
	CheckOut         string `json:"checkout" gorm:"column:checkout"`
	WorkingHours     int    `json:"working_hours" gorm:"column:working_hours"`
	AttendanceStatus bool   `json:"status" gorm:"column:attendance_status"`
}
