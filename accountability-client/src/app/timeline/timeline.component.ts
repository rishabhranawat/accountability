import { FeedTask } from './../models/feed-task.model';
import { FeedService } from './../common/services/feed.service';
import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import { Subscription, Observable } from 'rxjs';

@Component({
  selector: 'app-timeline',
  templateUrl: './timeline.component.html',
  styleUrls: ['./timeline.component.scss']
})
export class TimelineComponent implements OnInit, OnDestroy {

  @Input()
  public tasks$: Observable<object>;

  public tasks: FeedTask[];

  private subscription: Subscription;

  public taskView: boolean;
  public taskId: number;

  constructor(
    private feedService: FeedService
  ) {
    this.subscription = new Subscription();
  }

  ngOnInit(): void {
    if (this.tasks$ === null || this.tasks$ === undefined){
      this.tasks$ = this.feedService.getFeed();
    }
    this.subscription.add(this.tasks$.subscribe(data => {
      this.tasks = (data as any).Tasks as any[];
    }));
  }

  public switchTaskView(taskId: number): void {
    this.taskView = true;
    this.taskId = taskId;
  }

  public onSwitchBack(switchBack: boolean): void {
    this.taskView = false;
  }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }



}
