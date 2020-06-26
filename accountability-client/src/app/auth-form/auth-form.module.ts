import { CommonModule } from './../common/common.module';
import { AuthService } from './services/auth.service';
import { LoginFormComponent } from './login-form/login-form.component';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    LoginFormComponent
  ],
  imports: [
    FormsModule
  ],
  providers: [
    AuthService
  ],
  exports: [LoginFormComponent],
  bootstrap: []
})
export class AuthFormModule { }
