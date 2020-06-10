package models

type TaskMilestoneReaction struct {
	comment string
}

type TaskMilestone struct {
	description string
	reactions []TaskMilestoneReaction
}