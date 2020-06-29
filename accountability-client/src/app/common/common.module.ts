import { RequestHandlerService } from './services/request-handler.service';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  imports: [
    HttpClientModule,
  ],
  providers: [
    RequestHandlerService
  ],
  bootstrap: []
})
export class CommonModule { }
