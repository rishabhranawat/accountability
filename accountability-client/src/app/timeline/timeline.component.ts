import { FeedService } from './../common/services/feed.service';
import { Component, OnInit, Input, OnDestroy } from '@angular/core';
import { Subscription, Observable } from 'rxjs';
import { isNullOrUndefined } from 'util';

@Component({
  selector: 'app-timeline',
  templateUrl: './timeline.component.html',
  styleUrls: ['./timeline.component.scss']
})
export class TimelineComponent implements OnInit, OnDestroy {

  @Input()
  public tasks$: Observable<object>;

  public tasks: any[];

  private subscription: Subscription;

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

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
  }

}
