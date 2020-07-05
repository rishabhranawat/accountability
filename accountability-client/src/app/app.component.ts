import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'accountability-client';

  //todo:  if the user is not logged in, redirect to login page
  // else redirect to something else
}
