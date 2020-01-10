import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Subscription } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';

@Component({
  selector: 'app-quick-card-pesawat',
  templateUrl: './quick-card-pesawat.component.html',
  styleUrls: ['./quick-card-pesawat.component.sass']
})
export class QuickCardPesawatComponent implements OnInit {

  flight: Object[]
  supkrep : Subscription
  content: Object

  constructor(
    private service: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private emitterService: EventEmitterService
  ) {

   }

  ngOnInit() {
  }

}
