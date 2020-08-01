import { RequestHandlerService } from 'src/app/common/services/request-handler.service';
import { Task } from 'src/app/models/task.model';
import { User } from 'src/app/models/user.model';
import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';


// TODO: why is this service in /store/task-management?
@Injectable()
export class TaskService {

    constructor(
        private requestService: RequestHandlerService
    ){}

    public createTask(task? : Task) : Observable<object> {
        return this.requestService.post('/tasks/create-task', task);
    }

    public updateTask(task? : Task) : Observable<object> {
        return this.requestService.post('/tasks/update-task', task);
    }

    public removeTask(task? : Task) : Observable<object> {
        return this.requestService.post('/tasks/remove-task', task);
    }

    public retrieveTasks(user? : User) : Observable<object>  {
        return this.requestService.post('/tasks/fetch-tasks', user);
    }

}
