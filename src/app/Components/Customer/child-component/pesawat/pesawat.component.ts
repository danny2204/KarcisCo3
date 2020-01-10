import { Component, OnInit } from '@angular/core';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';

@Component({
  selector: 'app-pesawat',
  templateUrl: './pesawat.component.html',
  styleUrls: ['./pesawat.component.sass']
})
export class PesawatComponent implements OnInit {

  flightData: Object[]
  test: Object

  constructor(
    private eventEmitterService: EventEmitterService,
    private sharedService: SharedServiceService
  ) { }

  ngOnInit() {
    if(this.eventEmitterService.subsVar2 == undefined) {
      this.eventEmitterService.subsVar2 = this.eventEmitterService.invokeAssignFunction.subscribe( async () => {
        this.flightData = this.sharedService.flightData
      })
    }
  }

}
