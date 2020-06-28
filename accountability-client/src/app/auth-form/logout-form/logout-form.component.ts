import { TasksService } from './../../common/services/tasks.service';
import { User } from './../../models/user.model';
import { AuthService } from './../services/auth.service';
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
    // this.authService.logout({UserName: 'rish 4'} as User).subscribe((data: any) => {
    //   console.log(data);
    // });

    this.tasksService.getTasks().subscribe((data: any) => {
      console.log(data);
    });
  }

}
