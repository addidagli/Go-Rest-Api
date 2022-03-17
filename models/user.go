package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
