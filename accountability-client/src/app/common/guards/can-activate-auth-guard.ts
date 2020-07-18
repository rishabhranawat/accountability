import { isAuthenticated } from './../../store/user-details/user-selector';
import { UserService } from './../../store/user-details/services/user.service';
import { AuthService } from './../services/auth.service';
import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { combineLatest, BehaviorSubject, of } from 'rxjs';
import { map, tap, catchError } from 'rxjs/operators';

@Injectable()
export class CanActivateAuthGuard implements CanActivate {
  constructor(private authService: AuthService, private userService: UserService) { }

  //Todo: check store first
  canActivate() {
    return this.userService.login().pipe(
      map(response => true),
      catchError(error => of(false))
    );
  }
}
