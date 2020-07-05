import { AuthFormComponent } from './auth-form.component';
import { LogoutFormComponent } from './logout-form/logout-form.component';
import { CommonModule } from '@angular/common';
import { AuthService } from './services/auth.service';
import { LoginFormComponent } from './login-form/login-form.component';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

@NgModule({
  declarations: [
    LoginFormComponent,
    LogoutFormComponent,
    AuthFormComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    FontAwesomeModule
  ],
  providers: [
    AuthService
  ],
  exports: [LoginFormComponent, LogoutFormComponent, AuthFormComponent],
  bootstrap: []
})
export class AuthFormModule { }
