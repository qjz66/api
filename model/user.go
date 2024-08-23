package model

type User struct {
	ID       int64  `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username" `
	Password string `json:"password"`
}
