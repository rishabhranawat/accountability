import { Task } from 'src/app/models/task.model';

export interface CreateTaskRequest {
  UserTask: Task;
  TrackerEmails: string[];
}
