package models

// for catch
type Checking struct {
	UserId        int     `json:"user_id"`
	UserLongitude float32 `json:"user_longitude"`
	UserLatitude  float32 `json:"user_latitude"`
	CurrentDate   string  `json:"current_date" gorm:"column:dates"`
}

type Attendances struct {
	UserId        int     `json:"user_id"`
	UserLongitude float32 `json:"user_longitude"`
	UserLatitude  float32 `json:"user_latitude"`
	CurrentDate   string  `json:"current_date" gorm:"column:dates"`
	CheckIn       string  `json:"checkin" gorm:"column:checkin"`
	CheckOut      string  `json:"checkout" gorm:"column:checkout"`
}
