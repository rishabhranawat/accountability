import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-auth-form',
  templateUrl: './auth-form.component.html',
  styleUrls: ['./auth-form.component.scss']
})
export class AuthFormComponent implements OnInit {

  public showLogin: boolean;

  constructor() {
  }

  ngOnInit(): void {
    this.showLogin = true;
  }

  switchView(): void {
    this.showLogin = !this.showLogin;
  }

}
