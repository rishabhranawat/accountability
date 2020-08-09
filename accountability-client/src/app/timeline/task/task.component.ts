import { Task } from './../../models/task.model';
import { TaskService } from '../../common/services/task.service';
import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-task',
  templateUrl: './task.component.html',
  styleUrls: ['./task.component.scss']
})
export class TaskComponent implements OnInit {

  @Input()
  public taskId: number;

  public task: Task;

  constructor(
    private taskService: TaskService
  ) { }

  ngOnInit(): void {
    this.taskService.getTaskDetails(this.taskId).subscribe((data) => {
      this.task  = data as Task;
    });
  }

}
