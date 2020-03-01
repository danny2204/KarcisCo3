import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PromoDetailPageComponent } from './promo-detail-page.component';

describe('PromoDetailPageComponent', () => {
  let component: PromoDetailPageComponent;
  let fixture: ComponentFixture<PromoDetailPageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PromoDetailPageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PromoDetailPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
