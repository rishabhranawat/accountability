import { User } from './../../models/user.model';
import { AuthService } from './../services/auth.service';
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
  }

  // Todo validations
  login(): void {
    this.isProcessing = true;
    this.authService.login({
      Email: this.email,
      Password: this.password
    } as User).subscribe(
      (data: any) => {
        this.isProcessing = false;
        this.router.navigate(['/dashboard']);
      },
      (error: any) => {
        this.isProcessing = false;
        console.log(error);
      }
    );
  }

}
