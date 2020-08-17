import { AuthService } from './../common/services/auth.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  public toggled = false;
  public viewType = 'dashboard';

  constructor(
    private authService: AuthService
  ) {}

  ngOnInit() {
  }

  public toggle() {
    this.toggled = !this.toggled;
  }

  public switchView(switchToViewType: string) {
    this.viewType = switchToViewType;
  }

  logout() {
    this.authService.logout();
  }

}
