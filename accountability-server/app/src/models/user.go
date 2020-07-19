package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model

	UserName string `gorm:"type:varchar(100);unique;not null"`
	Email    string `gorm:"type:varchar(100);unique_index;not null"`
	Password string
	FullName string `gorm:"type:varchar(255);"`
}

func (u User) GetUserName() string {
	return u.UserName
}
