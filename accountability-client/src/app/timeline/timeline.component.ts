import { FeedService } from './../common/services/feed.service';
import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-timeline',
  templateUrl: './timeline.component.html',
  styleUrls: ['./timeline.component.scss']
})
export class TimelineComponent implements OnInit {

  constructor(
    private feedService: FeedService
  ) { }

  public tasks: any[];

  ngOnInit(): void {
    this.feedService.getFeed().subscribe(data => {
      this.tasks = (data as any).Tasks as any[];
    });
  }

}
