import { AuthService } from '../../common/services/auth.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-logout-form',
  templateUrl: './logout-form.component.html',
  styleUrls: ['./logout-form.component.scss']
})
export class LogoutFormComponent implements OnInit {

  constructor(
    private authService: AuthService,
  ) { }

  ngOnInit(): void {
    this.authService.logout();
  }

  logout(): void {
    this.authService.logout();
  }

}
