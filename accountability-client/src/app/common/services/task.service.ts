import { TaskComment } from './../../models/task-comment.model';
import { CreateTaskRequest } from '../requests/create-task-request';
import { RequestHandlerService } from 'src/app/common/services/request-handler.service';
import { Task } from 'src/app/models/task.model';
import { User } from 'src/app/models/user.model';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';

@Injectable()
export class TaskService {

  constructor(
    private requestService: RequestHandlerService
  ) { }

  public createTask(task?: CreateTaskRequest): Observable<object> {
    return this.requestService.post('/tasks/create-task', task);
  }

  public updateTask(task?: Task): Observable<object> {
    return this.requestService.post('/tasks/update-task', task);
  }

  public removeTask(task?: Task): Observable<object> {
    return this.requestService.post('/tasks/remove-task', task);
  }

  public retrieveTasks(user?: User): Observable<object> {
    return this.requestService.post('/tasks/fetch-tasks', user);
  }

  public getTaskDetails(taskId: number): Observable<object> {
    return this.requestService.get('/tasks/fetch-task-details/' + taskId);
  }

  public postComment(taskComment: TaskComment): Observable<object> {
    return this.requestService.post('/tasks/create-task-comment', taskComment);
  }

  public getComments(taskId: number): Observable<object> {
    return this.requestService.get('/tasks/fetch-task-comments/' + taskId);
  }
}
