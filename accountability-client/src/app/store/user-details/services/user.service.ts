import { RequestHandlerService } from '../../../common/services/request-handler.service';
import { User } from '../../../models/user.model';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { isNullOrUndefined } from 'util';

@Injectable()
export class UserService {

  constructor(
    private requestService: RequestHandlerService){
  }

  public login(user?: User): Observable<object> {
    if (user === null || user === undefined){
      return this.requestService.post('/auth/login', {});
    }
    return this.requestService.post('/auth/login', user);

  }

  public logout(user?: User): Observable<object> {
    return this.requestService.post('/auth/logout', user);
  }

  public create(user: User): Observable<object> {
    return this.requestService.post('/auth/create', user);
  }

}
