import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HiburanDetailComponent } from './hiburan-detail.component';

describe('HiburanDetailComponent', () => {
  let component: HiburanDetailComponent;
  let fixture: ComponentFixture<HiburanDetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HiburanDetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HiburanDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
