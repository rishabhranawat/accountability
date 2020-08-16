import { TaskService } from './../../../common/services/task.service';
import { TaskComment } from './../../../models/task-comment.model';
import { Component, OnInit, Input } from '@angular/core';

@Component({
  selector: 'app-send-reminder',
  templateUrl: './send-reminder.component.html',
  styleUrls: ['./send-reminder.component.scss']
})
export class SendReminderComponent implements OnInit {

  @Input()
  public taskID: number;
  public comment: string;
  public comments: TaskComment[];

  constructor(
    private taskService: TaskService
  ) { }

  ngOnInit(): void {
    this.taskService.getComments(this.taskID).subscribe((data) => {
      this.comments = data as TaskComment[];
    });
  }

  public postComment() {
    this.taskService.postComment({TaskReferID: this.taskID, Comment: this.comment} as TaskComment).subscribe((data) => {
      this.comments = data as TaskComment[];
    });
  }

}
