package models

type Task struct {
	name string;
	description string;
	worker User;
	trackers []User;
	milestones []TaskMilestone;
}