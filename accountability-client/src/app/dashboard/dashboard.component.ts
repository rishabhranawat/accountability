import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TaskService } from '../store/task-management/services/task.service';
import { User } from '../models/user.model';
import { Task } from '../models/task.model';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  dashboard: String = "dashboard"

  constructor(
  ) {}

  ngOnInit() {
  }

  getTasks(user?: User) : Task[] {
    var result : Task[] = []; //TODO should be retrieved from taskService
    return result;
  }

}
