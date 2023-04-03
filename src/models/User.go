package models

type User struct {
	Id       uint
	Name     string `gorm:"not null;"`
	Username string `gorm:"unique; not null;"`
	Password string `json:"not null;"`
}
