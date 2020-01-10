import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickCardPesawatComponent } from './quick-card-pesawat.component';

describe('QuickCardPesawatComponent', () => {
  let component: QuickCardPesawatComponent;
  let fixture: ComponentFixture<QuickCardPesawatComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickCardPesawatComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickCardPesawatComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
