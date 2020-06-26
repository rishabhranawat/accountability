import { Observable } from 'rxjs';
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class RequestHandlerService{

  private headerOptions: HttpHeaders = new HttpHeaders({
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json',
    'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept',
    'Access-Control-Allow-Methods': 'GET, POST, PUT, PATCH'
  });

  // TODO: move to env file
  private baseUrl = 'http://localhost:10000';

  constructor(
    private httpClient: HttpClient
  ){}


  public post(path: string, body: any): Observable<object>{
    return this.httpClient.post(this.baseUrl + path, body, {headers: this.headerOptions});
  }

  public get(path: string): Observable<object>{
    return this.httpClient.get(this.baseUrl + path);
  }

}
