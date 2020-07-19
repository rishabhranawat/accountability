package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Tracker struct {
	Tracker User
}

type Task struct {
	gorm.Model

	Name        string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:varchar(255);"`
	Workers     User
	Trackers    []Tracker
	Milestones  []TaskMilestone
	TaskId      string `gorm:"type:varchar(100);unique_index;not null"`
}
