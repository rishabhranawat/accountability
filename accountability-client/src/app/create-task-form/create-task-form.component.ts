import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TaskService } from '../store/task-management/services/task.service';
import { User } from '../models/user.model';
import { Task } from '../models/task.model';
import { NgForm } from '@angular/forms';
import { TaskMilestone } from '../models/task-milestone.model';
import { UserComponentFactory } from 'ag-grid-community';

@Component({
  selector: 'app-create-task-form',
  templateUrl: './create-task-form.component.html',
  styleUrls: ['./create-task-form.component.scss']
})
export class CreateTaskFormComponent implements OnInit {

  username: String;

  public name: string;
	public description: string;
  public milestones: TaskMilestone[];
  public currentTracker: string;

  public trackers: User[] = [];
  private worker: User;

  constructor(
    private route: ActivatedRoute,
    private taskService: TaskService
  ) {}

  ngOnInit() {

  }

  createTaskHandler(): void {
    var task: Task = {
      Name: this.name,
      Description: this.description,
      Trackers: this.trackers
    } as Task
    this.taskService.createTask(task).subscribe((data: any) => {
      console.log(data);
    }, (error: any) => {console.log("error!")});
  }

  addTracker() {
    var user: User = {
      Email: this.currentTracker,
    } as User
    this.trackers.push(user);
    this.currentTracker = '';
  }

}
