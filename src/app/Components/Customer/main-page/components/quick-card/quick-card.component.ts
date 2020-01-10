import { Output, Component, OnInit, ViewEncapsulation, ComponentFactoryResolver, ViewContainerRef } from '@angular/core';
import { QuickCardListComponent } from '../quick-card-content/quick-card-list/quick-card-list.component';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';

@Component({
  selector: 'app-quick-card',
  templateUrl: './quick-card.component.html',
  styleUrls: ['./quick-card.component.sass'],
  encapsulation: ViewEncapsulation.None
})
export class QuickCardComponent implements OnInit {

  type: string
  injectableComponent: QuickCardListComponent
  flightData: Object
  airlineName: string

  constructor(
    private eventEmitterService: EventEmitterService,
    private sharedServie: SharedServiceService
  ) { }

  inject(type: string) {
    console.log(type)
    this.type = type
    this.eventEmitterService.onInject(type);
  }

  ngOnInit() {
    
  }

  scrollToTop(e1: HTMLElement, e2:HTMLElement){
    e1.scrollIntoView({behavior:'smooth'})
    e1.classList.add('focus')
    e2.style.display='block'
  }

  backToNormal(e: HTMLElement){
    e.style.display='none'
  }

}
