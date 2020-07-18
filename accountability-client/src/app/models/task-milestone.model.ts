import { TaskMilestoneReaction } from './task-milestone-reaction.model';

export interface TaskMilestone {
    description: string;
    reactions: TaskMilestoneReaction[];
}