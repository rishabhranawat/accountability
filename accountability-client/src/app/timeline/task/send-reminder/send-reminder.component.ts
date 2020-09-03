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

  public imagePath;
  imageURL: any;
  public message: string;
  private file: any;

  url: string | ArrayBuffer = '';

  constructor(
    private taskService: TaskService
  ) { }

  ngOnInit(): void {
    this.taskService.getComments(this.taskID).subscribe((data) => {
      this.comments = data as TaskComment[];
    });
  }

  public postComment() {
    const formData = new FormData();
    formData.append('uploadFile', this.file);
    formData.append('TaskReferID', this.taskID.toString());
    formData.append('Description', this.comment);

    this.taskService.postTaskUpdate(formData).subscribe((data) => {
      console.log(data);
    });
    // this.taskService.postComment({TaskReferID: this.taskID, Comment: this.comment} as TaskComment).subscribe((data) => {
    //   this.comments = data as TaskComment[];
    // });
  }


  onSelectFile(event) {
    if (event.target.files && event.target.files[0]) {
      const reader = new FileReader();
      this.file = event.target.files[0];
      reader.readAsDataURL(event.target.files[0]); // read file as data url

      reader.onload = (e) => { // called once readAsDataURL is completed
        this.url = e.target.result;
      }
    }
  }

}
