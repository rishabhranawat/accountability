import { CommonModule } from './common/common.module';
import { AuthFormModule } from './auth-form/auth-form.module';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AgGridModule } from 'ag-grid-angular';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { TasksGridComponent } from './tasks-grid/tasks-grid.component';
import { CreateTaskFormComponent } from './create-task-form/create-task-form.component';
import { TimelineComponent } from './timeline/timeline.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { CookieService } from 'ngx-cookie-service';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

@NgModule({
  declarations: [
    AppComponent,
    TasksGridComponent,
    CreateTaskFormComponent,
    TimelineComponent,
    DashboardComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    AgGridModule.withComponents([]),
    NgbModule,
    CommonModule,
    AuthFormModule,
    FontAwesomeModule
  ],
  providers: [CookieService],
  bootstrap: [AppComponent]
})
export class AppModule { }
