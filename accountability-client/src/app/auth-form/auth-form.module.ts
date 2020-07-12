import { UserDetailsStoreModule } from './../store/user-details/user-details-store.module';
import { AuthService } from './services/auth.service';
import { SignUpFormComponent } from './sign-up-form/sign-up-form.component';
import { AuthFormComponent } from './auth-form.component';
import { LogoutFormComponent } from './logout-form/logout-form.component';
import { CommonModule } from '@angular/common';
import { LoginFormComponent } from './login-form/login-form.component';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

@NgModule({
  declarations: [
    LoginFormComponent,
    LogoutFormComponent,
    AuthFormComponent,
    SignUpFormComponent
  ],
  imports: [
    CommonModule,
    FormsModule,
    FontAwesomeModule,
    UserDetailsStoreModule
  ],
  providers: [
    AuthService
  ],
  exports: [LoginFormComponent, LogoutFormComponent, AuthFormComponent, SignUpFormComponent],
  bootstrap: []
})
export class AuthFormModule { }
