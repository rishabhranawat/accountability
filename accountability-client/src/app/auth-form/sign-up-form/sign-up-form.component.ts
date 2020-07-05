import { Router } from '@angular/router';
import { User } from './../../models/user.model';
import { AuthService } from './../services/auth.service';
import { Component, OnInit } from '@angular/core';
import { finalize } from 'rxjs/operators';

@Component({
  selector: 'app-sign-up-form',
  templateUrl: './sign-up-form.component.html',
  styleUrls: ['./sign-up-form.component.scss']
})
export class SignUpFormComponent implements OnInit {

  public email: string;
  public username: string;
  public password: string;
  public confirmPassword: string;
  public errors: string;

  public isProcessing: boolean;

  constructor(
    private authService: AuthService,
    private router: Router
  ) { }

  ngOnInit(): void {
  }

  create(): void {
    const user: User = {
      UserName: this.username,
      Email: this.email,
      Password: this.password
    } as User;
    this.isProcessing = true;
    this.authService.create(user)
    .pipe(finalize(() => this.isProcessing = false))
    .subscribe(
      (success: any) => {
        this.router.navigate(['dashboard']);
      },
      (error: any) => {
        this.errors = error.error;
      },
    );
  }





}
