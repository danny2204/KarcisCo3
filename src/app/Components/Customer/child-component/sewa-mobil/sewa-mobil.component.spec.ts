import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SewaMobilComponent } from './sewa-mobil.component';

describe('SewaMobilComponent', () => {
  let component: SewaMobilComponent;
  let fixture: ComponentFixture<SewaMobilComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SewaMobilComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SewaMobilComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
