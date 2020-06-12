package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


type Task struct {
	gorm.Model

	name string;
	description string;
	worker User;
	trackers []User;
	milestones []TaskMilestone;
}