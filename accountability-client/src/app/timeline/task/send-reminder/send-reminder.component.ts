import { TaskUpdate } from './../../../models/task-update.model';
import { TaskService } from './../../../common/services/task.service';
import { TaskComment } from './../../../models/task-comment.model';
import { Component, OnInit, Input } from '@angular/core';
import { forkJoin } from 'rxjs';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-send-reminder',
  templateUrl: './send-reminder.component.html',
  styleUrls: ['./send-reminder.component.scss']
})
export class SendReminderComponent implements OnInit {

  @Input()
  public taskID: number;
  public comment: string;
  public update: string;
  public comments: TaskComment[];
  public updates: TaskUpdate[];

  public message: string;
  private file: any;
  private updateFile;

  commentUrl: string | ArrayBuffer = '';
  updateUrl: string | ArrayBuffer = '';

  constructor(
    private taskService: TaskService
  ) { }

  ngOnInit(): void {

    const comments$ = this.taskService.getComments(this.taskID);
    const updates$ = this.taskService.getUpdates(this.taskID);


    forkJoin([comments$, updates$]).subscribe(response => {
      this.comments = response[0] as TaskComment[];
      this.updates = response[1] as TaskUpdate[];
    });
  }

  public postComment() {

    this.taskService.postComment({TaskReferID: this.taskID, Comment: this.comment} as TaskComment).subscribe((data) => {
      this.comments = data as TaskComment[];
    });
  }

  public postUpdate() {
    const formData = new FormData();

    if (this.updateFile !== null && this.updateFile) {
      formData.append('uploadFile', this.updateFile);
    }
    formData.append('TaskReferID', this.taskID.toString());
    formData.append('Description', this.update);

    this.taskService.postTaskUpdate(formData).subscribe((data) => {
      this.updates = data as TaskUpdate[];
    });
  }


  onSelectFile(event, commentOrUpdate) {
    if (event.target.files && event.target.files[0]) {
      const reader = new FileReader();
      if (commentOrUpdate === 'comment'){
        this.file = event.target.files[0];
      } else {
        this.updateFile = event.target.files[0];
      }
      reader.readAsDataURL(event.target.files[0]); // read file as data url

      reader.onload = (e) => { // called once readAsDataURL is completed
        if (commentOrUpdate === 'comment') {
          this.commentUrl = e.target.result;
        }
        else {
          this.updateUrl = e.target.result;
        }
      };
    }
  }

}
