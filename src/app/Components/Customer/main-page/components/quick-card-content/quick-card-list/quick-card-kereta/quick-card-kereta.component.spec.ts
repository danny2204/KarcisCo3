import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickCardKeretaComponent } from './quick-card-kereta.component';

describe('QuickCardKeretaComponent', () => {
  let component: QuickCardKeretaComponent;
  let fixture: ComponentFixture<QuickCardKeretaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickCardKeretaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickCardKeretaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
