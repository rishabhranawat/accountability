import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { RequestHandlerService } from './request-handler.service';

@Injectable({
  providedIn: 'root'
})
export class TasksService {

  constructor(
    private requestService: RequestHandlerService
  ){}

  public getTasks(): Observable<object> {
    return this.requestService.get('/tasks/get-tasks');
  }

}
