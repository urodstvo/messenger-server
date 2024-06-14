package models

import "time"

type User struct {
	Id        uint32     `gorm:"primary_key;autoIncrement" json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
