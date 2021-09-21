package models

type Dashboard struct {
	FullName        string `json:"full_name"`
	RoleDescription string `json:"role_description"`
	OfficeLongitude string `json:"office_longitude"`
	OfficeLatitude  string `json:"office_latitude"`
}
