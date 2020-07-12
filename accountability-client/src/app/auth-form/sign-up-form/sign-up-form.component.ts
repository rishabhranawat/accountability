import { AuthService } from './../services/auth.service';
import { Router } from '@angular/router';
import { User } from './../../models/user.model';
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
    this.authService.userAuthenticated().subscribe((data: boolean) => {
      this.isProcessing = data;
    });

    this.authService.userAuthenticated().subscribe((data: boolean) => {
      if (data){
        this.router.navigate(['dashboard']);
      }
    });
  }

  create(): void {
    const user: User = {
      UserName: this.username,
      Email: this.email,
      Password: this.password
    } as User;
    this.authService.create(user);
  }





}
