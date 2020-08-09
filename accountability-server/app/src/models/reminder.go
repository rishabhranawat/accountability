package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Reminder struct {
	gorm.Model

	Task              Task `gorm:"foreignkey:TaskReferID"`
	TaskReferID       uint
	Subject           string `gorm:"type:varchar(255);"`
	Message           string `gorm:"type:text"`
	RecepientEmail    string `gorm:"type:varchar(255);"`
	RecepientUsername string `gorm:"type:varchar(255);"`
}
