import { login, logout, setUserDetails, expireUserDetails, create } from './user-actions';
import { UserService } from './services/user.service';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { Injectable } from '@angular/core';
import { mergeMap, map, catchError } from 'rxjs/operators';
import { EMPTY } from 'rxjs';
import { User } from 'src/app/models/user.model';

@Injectable()
export class UserEffects {

  constructor(
    private actions$: Actions,
    private userService: UserService
  ) { }

  userLogin$ = createEffect(() => this.actions$.pipe(
    ofType(login),
    mergeMap((action) => this.userService.login(action.user)
      .pipe(
        map((success: any) => (setUserDetails({ user: {UserName: success.UserName, Email: success.Email} as User }))),
        catchError(() => EMPTY)
      ))
    )
  );

  userLogout$ = createEffect(() => this.actions$.pipe(
    ofType(logout),
    mergeMap(() => this.userService.logout()
      .pipe(
        map(success => (expireUserDetails()),
          catchError(() => EMPTY)
        ))
    )
  )
  );

  create$ = createEffect(() => this.actions$.pipe(
    ofType(create),
    mergeMap((action) => this.userService.create(action.user)
      .pipe(
        map((succes: any) => (setUserDetails({ user: {} as User }))),
        catchError(() => EMPTY)
      ))
    )
  );
}
