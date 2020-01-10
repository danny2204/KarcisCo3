import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { KeretaApiComponent } from './kereta-api.component';

describe('KeretaApiComponent', () => {
  let component: KeretaApiComponent;
  let fixture: ComponentFixture<KeretaApiComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KeretaApiComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeretaApiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
