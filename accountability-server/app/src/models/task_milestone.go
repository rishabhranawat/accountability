package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type TaskMilestoneReaction struct {
	gorm.Model

	Comment string `gorm:"type:varchar(255);"`
}

type TaskMilestone struct {
	gorm.Model

	Description string `gorm:"type:varchar(255);"`
	Reactions   []TaskMilestoneReaction
}
