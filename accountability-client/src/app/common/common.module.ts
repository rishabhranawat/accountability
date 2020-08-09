import { FeedService } from './services/feed.service';
import { CanActivateAuthGuard } from './guards/can-activate-auth-guard';
import { AuthService } from './services/auth.service';
import { RequestHandlerService } from './services/request-handler.service';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { TaskService } from './services/task.service';

@NgModule({
  imports: [
    HttpClientModule,
  ],
  providers: [
    RequestHandlerService,
    AuthService,
    TaskService,
    FeedService,
    CanActivateAuthGuard
  ],
  bootstrap: []
})
export class CommonAccountabilityModule { }
