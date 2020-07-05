import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-sign-up-form',
  templateUrl: './sign-up-form.component.html',
  styleUrls: ['./sign-up-form.component.scss']
})
export class SignUpFormComponent implements OnInit {

  public email: string;
  public username: string;
  public password: string;
  public isProcessing: boolean;

  constructor() { }

  ngOnInit(): void {
  }

}
