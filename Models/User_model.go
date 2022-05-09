package models

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (User) TableName() string {
	return "user"
}
