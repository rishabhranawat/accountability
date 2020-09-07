import { Observable } from 'rxjs';
import { RequestHandlerService } from 'src/app/common/services/request-handler.service';
import { Injectable } from '@angular/core';

@Injectable()
export class RelationshipMgmtService {

  constructor(
    private requestService: RequestHandlerService
  ) { }

  public createRelationship(data: any): Observable<object> {
    return this.requestService.post('/relationship/create-relationship', data);
  }

  public approveRelationship(data: any): Observable<object> {
    return this.requestService.post('/relationship/approve-relationship', data);
  }

  public deleteRelationship(data: any): Observable<object> {
    return this.requestService.post('/relationship/delete-relationship', data);
  }

}
