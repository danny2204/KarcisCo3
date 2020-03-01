import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { QuickCardMobilComponent } from './quick-card-mobil.component';

describe('QuickCardMobilComponent', () => {
  let component: QuickCardMobilComponent;
  let fixture: ComponentFixture<QuickCardMobilComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ QuickCardMobilComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(QuickCardMobilComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
