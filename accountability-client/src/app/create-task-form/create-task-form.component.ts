import { FeedService } from './../common/services/feed.service';
import { Subscription, Observable } from 'rxjs';
import { Task } from 'src/app/models/task.model';
import { CreateTaskRequest } from './../common/requests/create-task-request';
import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { TaskService } from '../store/task-management/services/task.service';

@Component({
  selector: 'app-create-task-form',
  templateUrl: './create-task-form.component.html',
  styleUrls: ['./create-task-form.component.scss']
})
export class CreateTaskFormComponent implements OnInit, OnDestroy {

  public taskRequest: CreateTaskRequest = { UserTask: {} as Task, TrackerEmails: [] as string[]} as CreateTaskRequest;
  public currentTracker: string;
  public tasks$: Observable<object>;

  private subscription: Subscription;

  constructor(
    private route: ActivatedRoute,
    private taskService: TaskService,
    private feedService: FeedService
  ) {
    this.subscription = new Subscription();
  }

  ngOnInit() {
    this.tasks$ = this.feedService.getUserSpecificFeed();
  }

  createTaskHandler(): void {
    this.taskService.createTask(this.taskRequest).subscribe((data: any) => {
      console.log(data);
    }, (error: any) => { console.log('error!'); });
  }

  addTracker() {
    this.taskRequest.TrackerEmails.push(this.currentTracker);
  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

}
