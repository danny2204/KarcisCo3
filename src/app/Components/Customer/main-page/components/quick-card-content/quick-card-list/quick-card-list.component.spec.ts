import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickCardListComponent } from './quick-card-list.component';

describe('QuickCardListComponent', () => {
  let component: QuickCardListComponent;
  let fixture: ComponentFixture<QuickCardListComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickCardListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickCardListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
