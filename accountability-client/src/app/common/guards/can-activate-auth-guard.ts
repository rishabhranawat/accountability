import { AuthService } from './../services/auth.service';
import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import { combineLatest, BehaviorSubject } from 'rxjs';

@Injectable()
export class CanActivateAuthGuard implements CanActivate {
  constructor(private authService: AuthService){}

  canActivate(){
    const activated = new BehaviorSubject<boolean>(false);

    this.authService.userAuthenticated().subscribe((data: boolean) => {
      if(data){
        activated.next(true);
      } else {
        this.authService.login();
        this.authService.requestProcessing().subscribe((requestProcessing: boolean) => {
          if(!requestProcessing){
            this.authService.userAuthenticated().subscribe((userLoggedIn: boolean) => {
              if(userLoggedIn){
                activated.next(true);
              }
            });
          }
        });
      }
    });

    return activated.asObservable();
  }
}
