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
    AuthFormModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
