package models


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


type TaskMilestoneReaction struct {
	gorm.Model
	comment string;
}

type TaskMilestone struct {
	gorm.Model
	description string;
	reactions []TaskMilestoneReaction;
}