import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PesawatComponent } from './pesawat.component';

describe('PesawatComponent', () => {
  let component: PesawatComponent;
  let fixture: ComponentFixture<PesawatComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PesawatComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PesawatComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
