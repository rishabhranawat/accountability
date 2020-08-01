import { Observable } from 'rxjs';
import { RequestHandlerService } from './request-handler.service';
import { Injectable } from '@angular/core';

@Injectable()
export class FeedService {

  constructor(
    private requestHandlerService: RequestHandlerService
  ) {}

  public getFeed(): Observable<object>{
    return this.requestHandlerService.get('/tasks/user-feed');
  }

}
