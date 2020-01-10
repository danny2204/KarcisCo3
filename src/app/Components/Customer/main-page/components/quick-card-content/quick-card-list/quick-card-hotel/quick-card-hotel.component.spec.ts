import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickCardHotelComponent } from './quick-card-hotel.component';

describe('QuickCardHotelComponent', () => {
  let component: QuickCardHotelComponent;
  let fixture: ComponentFixture<QuickCardHotelComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickCardHotelComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickCardHotelComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
