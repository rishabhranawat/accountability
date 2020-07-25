package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Tracker struct {
	User   `gorm:"foreignkey:UserID"`
	UserID uint
	Task   `gorm:"foreignkey:TaskID"`
	TaskID uint
}

type Task struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:varchar(255);"`
	User        User   `gorm:"foreignkey:UserID"`
	UserID      uint
}
