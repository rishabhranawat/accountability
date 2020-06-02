package models

import (
	"User",
	"TaskMilestone"
)

type Task struct {
	name string
	description string
	worker User
	trackers []User
	milestones []TaskMilestone
}