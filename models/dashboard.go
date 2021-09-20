package models

type Dashboard struct {
	FullName        string  `json:"full_name"`
	RoleDescription string  `json:"role_description"`
	OfficeLongitude float32 `json:"office_longitude"`
	OfficeLatitude  float32 `json:"office_latitude"`
	CurrentDate     string  `json:"dates"`
	UserLongitude   float32 `json:"user_longitude"`
	UserLatitude    float32 `json:"user_latitude"`
}
