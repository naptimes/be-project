package models

type Dashboard struct {
	FullName        string  `json:"full_name"`
	RoleDescription string  `json:"role_description"`
	OfficeLongitude float64 `json:"office_longitude"`
	OfficeLatitude  float64 `json:"office_latitude"`
}
