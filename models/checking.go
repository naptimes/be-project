package models

type Attendances struct {
	UserId        int     `json:"user_id"`
	UserLongitude float32 `json:"user_longitude"`
	UserLatitude  float32 `json:"user_latitude"`
	CurrentDate   string  `json:"current_date" gorm:"column:dates"`
}
