import { User } from './../../models/user.model';
import { AuthService } from './../services/auth.service';
import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent implements OnInit {

  public password: string;
  public email: string;

  public isProcessing: boolean;

  constructor(
    private authService: AuthService
  ) {

  }

  ngOnInit(): void {
  }

  login(): void {
    this.isProcessing = true;
    this.authService.login({
      Email: this.email,
      Password: this.password
    } as User).subscribe((data: any) => {
      this.isProcessing = false;
      console.log(data);
    });
  }

}
