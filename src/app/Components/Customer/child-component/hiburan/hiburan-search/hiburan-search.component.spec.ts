import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HiburanSearchComponent } from './hiburan-search.component';

describe('HiburanSearchComponent', () => {
  let component: HiburanSearchComponent;
  let fixture: ComponentFixture<HiburanSearchComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HiburanSearchComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HiburanSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
