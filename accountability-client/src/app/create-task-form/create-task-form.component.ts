import { Task } from 'src/app/models/task.model';
import { CreateTaskRequest } from './../common/requests/create-task-request';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TaskService } from '../store/task-management/services/task.service';
import { User } from '../models/user.model';
import { NgForm } from '@angular/forms';
import { TaskMilestone } from '../models/task-milestone.model';
import { UserComponentFactory } from 'ag-grid-community';

@Component({
  selector: 'app-create-task-form',
  templateUrl: './create-task-form.component.html',
  styleUrls: ['./create-task-form.component.scss']
})
export class CreateTaskFormComponent implements OnInit {

  public taskRequest: CreateTaskRequest = { UserTask: {} as Task, TrackerEmails: [] as string[]} as CreateTaskRequest;
  public currentTracker: string;

  constructor(
    private route: ActivatedRoute,
    private taskService: TaskService
  ) { }

  ngOnInit() {

  }

  createTaskHandler(): void {
    this.taskService.createTask(this.taskRequest).subscribe((data: any) => {
      console.log(data);
    }, (error: any) => { console.log('error!'); });
  }

  addTracker() {
    this.taskRequest.TrackerEmails.push(this.currentTracker);
  }

}
