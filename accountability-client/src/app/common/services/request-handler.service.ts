import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class RequestHandlerService{

  private headerOptions: HttpHeaders = new HttpHeaders({
    'Access-Control-Allow-Origin': 'http://localhost:4200',
    'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept',
    'Access-Control-Allow-Methods': 'GET, POST, PUT, PATCH',
    'Access-Control-Allow-Credentials': 'true'
  });

  // TODO: move to env file
  private baseUrl = 'http://localhost:10000';

  constructor(
    private httpClient: HttpClient,
    private cookieService: CookieService
  ){}


  public post(path: string, body: any): Observable<object>{
    return this.httpClient.post(this.baseUrl + path, body, {headers: this.headerOptions, withCredentials: true});
  }

  public get(path: string): Observable<object>{
    return this.httpClient.get(this.baseUrl + path, {headers: this.headerOptions, withCredentials: true});
  }

}
