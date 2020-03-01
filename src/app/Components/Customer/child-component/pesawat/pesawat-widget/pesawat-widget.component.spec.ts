import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PesawatWidgetComponent } from './pesawat-widget.component';

describe('PesawatWidgetComponent', () => {
  let component: PesawatWidgetComponent;
  let fixture: ComponentFixture<PesawatWidgetComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PesawatWidgetComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PesawatWidgetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
