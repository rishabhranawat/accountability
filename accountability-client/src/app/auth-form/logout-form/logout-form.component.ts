import { AuthService } from './../services/auth.service';
import { TasksService } from './../../common/services/tasks.service';
import { User } from './../../models/user.model';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-logout-form',
  templateUrl: './logout-form.component.html',
  styleUrls: ['./logout-form.component.scss']
})
export class LogoutFormComponent implements OnInit {

  constructor(
    private authService: AuthService,
    private tasksService: TasksService
  ) { }

  ngOnInit(): void {
  }

  logout(): void {
    this.authService.logout();
  }

}
