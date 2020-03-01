import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickCardHiburanComponent } from './quick-card-hiburan.component';

describe('QuickCardHiburanComponent', () => {
  let component: QuickCardHiburanComponent;
  let fixture: ComponentFixture<QuickCardHiburanComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickCardHiburanComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickCardHiburanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
