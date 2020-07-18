import { UserState } from '../../store/user-details/user-state';
import { isAuthenticated } from '../../store/user-details/user-selector';
import { login, logout, create } from '../../store/user-details/user-actions';

import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/models/user.model';
import { Store, select } from '@ngrx/store';
import { isProcessing } from 'src/app/store/user-details/user-selector';

// maybe move to /store/user-details.services
@Injectable()
export class AuthService {

  constructor(
    private store: Store<UserState>
  ){
  }

  public login(data?: User) {
    this.store.dispatch(login({user: data}));
  }

  public logout() {
    this.store.dispatch(logout());
  }

  public create(data: User){
    this.store.dispatch(create({user: data}));
  }

  public requestProcessing(): Observable<boolean> {
    return this.store.pipe(select(isProcessing));
  }

  public userAuthenticated(): Observable<boolean> {
    return this.store.pipe(select(isAuthenticated));
  }


}
