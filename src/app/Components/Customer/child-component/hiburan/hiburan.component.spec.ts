import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HiburanComponent } from './hiburan.component';

describe('HiburanComponent', () => {
  let component: HiburanComponent;
  let fixture: ComponentFixture<HiburanComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HiburanComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HiburanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
