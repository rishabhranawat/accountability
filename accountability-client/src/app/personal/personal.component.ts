import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { User } from '../models/user.model';
import { Task } from '../models/task.model';

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.scss']
})
export class PersonalComponent implements OnInit {

  dashboard: String = "personal"

  constructor(
  ) {}

  ngOnInit() {
  }

  getTasks(user?: User) : Task[] {
    var result : Task[] = []; //TODO should be retrieved from taskService
    return result;
  }

}
