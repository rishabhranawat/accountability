import { AuthService } from './../services/auth.service';
import { User } from './../../models/user.model';
import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { faUser } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent implements OnInit {

  public password: string;
  public email: string;
  public faUser = faUser;
  public errors: string[];
  public isProcessing: boolean;

  constructor(
    private authService: AuthService,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    this.authService.userAuthenticated().subscribe((data: boolean) => {
      this.isProcessing = data;
    });

    this.authService.userAuthenticated().subscribe((data: boolean) => {
      this.router.navigate(['dashboard']);
    });
  }

  // Todo validations
  login(): void {
    this.isProcessing = true;
    this.authService.login({
      Email: this.email,
      Password: this.password
    } as User);
  }



}
