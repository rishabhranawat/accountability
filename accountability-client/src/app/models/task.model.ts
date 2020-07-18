import { User } from './user.model';
import { TaskMilestone } from './task-milestone.model';

export interface Task {
    Name:        string;
	Description: string;
	Workers:     User;
	Trackers:    User[];
    Milestones:  TaskMilestone[];
	Id:          string;
}