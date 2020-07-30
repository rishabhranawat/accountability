package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Tracker struct {
	gorm.Model

	User        User `gorm:"foreignkey:UserReferID"`
	UserReferID uint
	Task        Task `gorm:"foreignkey:TaskReferID"`
	TaskReferID uint
}

type TaskUpdate struct {
	gorm.Model

	Task        Task `gorm:"foreignKey:TaskReferID"`
	TaskReferID uint
	Description string `gorm:"type:varchar(255);"`
}

type Task struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:varchar(255);"`
	User        User   `gorm:"foreignkey:UserID"`
	UserID      uint
}
