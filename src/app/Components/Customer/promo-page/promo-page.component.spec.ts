import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PromoPageComponent } from './promo-page.component';

describe('PromoPageComponent', () => {
  let component: PromoPageComponent;
  let fixture: ComponentFixture<PromoPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PromoPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PromoPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
