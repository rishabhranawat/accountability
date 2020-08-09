import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SendReminderComponent } from './send-reminder.component';

describe('SendReminderComponent', () => {
  let component: SendReminderComponent;
  let fixture: ComponentFixture<SendReminderComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SendReminderComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SendReminderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
