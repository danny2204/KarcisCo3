import { Input, Output, Component, ComponentFactoryResolver, ViewContainerRef, EventEmitter, ViewChild } from '@angular/core';
import { QuickCardPesawatComponent } from './quick-card-pesawat/quick-card-pesawat.component';
import { QuickCardHotelComponent } from './quick-card-hotel/quick-card-hotel.component';
import { QuickCardKeretaComponent } from './quick-card-kereta/quick-card-kereta.component';
import { QuickCardMobilComponent } from './quick-card-mobil/quick-card-mobil.component';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';

@Component({
  selector: 'app-quick-card-list',
  templateUrl: './quick-card-list.component.html',
  styleUrls: ['./quick-card-list.component.sass']
})

export class QuickCardListComponent {
  @ViewChild('container', {read: ViewContainerRef, static: true}) container: ViewContainerRef;
  @Input('type') type: string
  @Output() myEvent = new EventEmitter()

  components = [];
  PesawatComponent = QuickCardPesawatComponent
  HotelComponent = QuickCardHotelComponent
  KeretaComponent = QuickCardKeretaComponent
  MobilComponent = QuickCardMobilComponent
  
  constructor(
    private componentFactoryResolver: ComponentFactoryResolver,
    private eventEmitterService: EventEmitterService
  ) { }

  public removeComponent() {
    var component;
    component = this.components.find((component) => {
      if(this.type == "pesawat") {
        component.instance instanceof this.PesawatComponent
      } else if(this.type == "hotel") {
        component.instance instanceof this.HotelComponent
      } else if(this.type == "kereta") {
        component.instance instanceof this.KeretaComponent
      }
      return component
    });
    const componentIdx = this.container.indexOf(component)
    if (componentIdx !== -1) {
      this.container.remove(this.container.indexOf(component))
      this.components.splice(componentIdx, 1)
    }
  }

  public injectComponent() {
    if(this.components.length != 0){
      this.removeComponent()
    }

    if(this.type == "pesawat") {
      const factory = this.componentFactoryResolver.resolveComponentFactory(this.PesawatComponent)
      const ref = this.container.createComponent(factory)
      ref.changeDetectorRef.detectChanges()
      this.components.push(ref)
    }
    else if(this.type == "hotel") {
      const factory = this.componentFactoryResolver.resolveComponentFactory(this.HotelComponent)
      const ref = this.container.createComponent(factory)
      ref.changeDetectorRef.detectChanges()
      this.components.push(ref)
    }
    else if(this.type == "kereta") {
      const factory = this.componentFactoryResolver.resolveComponentFactory(this.KeretaComponent)
      const ref = this.container.createComponent(factory)
      ref.changeDetectorRef.detectChanges()
      this.components.push(ref)
    }
    else if(this.type == "sewa-mobil") {
      const factory = this.componentFactoryResolver.resolveComponentFactory(this.MobilComponent)
      const ref = this.container.createComponent(factory)
      ref.changeDetectorRef.detectChanges()
      this.components.push(ref)
    }

    this.myEvent.emit(null)
  }

  ngOnInit() {
    if(this.eventEmitterService.subsVar == undefined) {
      this.eventEmitterService.subsVar = this.eventEmitterService.invokeInjectFunction.subscribe(() => {
        this.injectComponent()
      })
    }
    this.type = "pesawat"
    this.injectComponent()
  }

}
