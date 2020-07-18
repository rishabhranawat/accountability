import { Router } from '@angular/router';
import { AuthService } from './common/services/auth.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{
  title = 'accountability-client';

  constructor(
    private authService: AuthService,
    private route: Router
  ){

  }

  ngOnInit(){
  }
}
