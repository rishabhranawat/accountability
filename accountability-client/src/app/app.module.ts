import { CommonModule } from '@angular/common';
import { UserDetailsStoreModule } from './store/user-details/user-details-store.module';
import { StoreModule } from '@ngrx/store';
import { CommonAccountabilityModule } from './common/common.module';
import { AuthFormModule } from './auth-form/auth-form.module';
import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { AgGridModule } from 'ag-grid-angular';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CreateTaskFormComponent } from './create-task-form/create-task-form.component';
import { TimelineComponent } from './timeline/timeline.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { CookieService } from 'ngx-cookie-service';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import { PersonalComponent } from './personal/personal.component';
import { ProfileComponent } from './profile/profile.component';
import { EffectsModule } from '@ngrx/effects';
import { FormsModule } from '@angular/forms';
import { SendReminderComponent } from './timeline/task/send-reminder/send-reminder.component';
import { TaskComponent } from './timeline/task/task.component';

@NgModule({
  declarations: [
    AppComponent,
    CreateTaskFormComponent,
    TimelineComponent,
    DashboardComponent,
    PersonalComponent,
    ProfileComponent,
    SendReminderComponent,
    TaskComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    AgGridModule.withComponents([]),
    NgbModule,
    CommonModule,
    CommonAccountabilityModule,
    AuthFormModule,
    FormsModule,
    FontAwesomeModule,
    StoreModule.forRoot({}),
    EffectsModule.forRoot([]),
    UserDetailsStoreModule
  ],
  providers: [CookieService],
  bootstrap: [AppComponent]
})
export class AppModule { }
