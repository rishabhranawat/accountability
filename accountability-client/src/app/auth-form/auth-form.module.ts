import { LogoutFormComponent } from './logout-form/logout-form.component';
import { CommonModule } from './../common/common.module';
import { AuthService } from './services/auth.service';
import { LoginFormComponent } from './login-form/login-form.component';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    LoginFormComponent,
    LogoutFormComponent
  ],
  imports: [
    FormsModule
  ],
  providers: [
    AuthService
  ],
  exports: [LoginFormComponent, LogoutFormComponent],
  bootstrap: []
})
export class AuthFormModule { }
