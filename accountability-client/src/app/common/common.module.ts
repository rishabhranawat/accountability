import { CanActivateAuthGuard } from './guards/can-activate-auth-guard';
import { AuthService } from './services/auth.service';
import { RequestHandlerService } from './services/request-handler.service';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  imports: [
    HttpClientModule,
  ],
  providers: [
    RequestHandlerService,
    AuthService,
    CanActivateAuthGuard
  ],
  bootstrap: []
})
export class CommonAccountabilityModule { }
