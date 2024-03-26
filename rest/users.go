package rest

//RESTUser represents REST data for users
type User struct {
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Login    string `gorm:"type:varchar(100);unique" json:"login"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
