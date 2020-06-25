import { User } from './../../models/user.model';
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {

  constructor(private http: HttpClient){
  }

  public login(user: User): Observable<object> {
    const headerOptions: HttpHeaders = new HttpHeaders({
      'Access-Control-Allow-Origin': '*',
      'Content-Type': 'application/json',
      'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept',
      'Access-Control-Allow-Methods': 'GET, POST, PUT, PATCH'
    });
    return this.http.post('http://localhost:10000/auth/login', user, {headers: headerOptions});
  }

  public logout(user: User): Observable<object> {
    return this.http.get('http://localhost:10000/');
  }

  public create(user: User): Observable<object> {
    return this.http.post('http://localhost:10000/auth/create', user);
  }

}
