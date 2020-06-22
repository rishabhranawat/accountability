import { User } from './../../models/user.model';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {

  constructor(private http: HttpClient){
  }

  public login(user: User): Observable<object> {
    return this.http.post('http://localhost:10000/auth/login', user);
  }

  public logout(user: User): Observable<object> {
    return this.http.post('http://localhost:10000/auth/logout', user);
  }

  public create(user: User): Observable<object> {
    return this.http.post('http://localhost:10000/auth/create', user);
  }

}
