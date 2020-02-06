import { Component, AfterViewInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { Router } from '@angular/router'

@Component({
  selector: 'app-quick-card-pesawat',
  templateUrl: './quick-card-pesawat.component.html',
  styleUrls: ['./quick-card-pesawat.component.sass']
})
export class QuickCardPesawatComponent implements AfterViewInit {

  flight: Object[]
  supkrep : Subscription
  content: Object
  fromValue: string
  toValue: string
  supkrep2: Subscription

  constructor(
    private service: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private emitterService: EventEmitterService,
    private router: Router
  ) {

   }

  ngAfterViewInit() {
    this.supkrep = this.service.getAllAirport().subscribe(async query => {
      this.flight = query.data.getAllAirport
    })
  }

  public getFlight() {
    if(this.fromValue.localeCompare(this.toValue) == 0) {
      alert("Destination and From Location cannot be the same")
    } else {
      this.supkrep2 = this.service.getFlightByToAndFrom(this.toValue, this.fromValue).subscribe(async query => {
        this.sharedService.flightSearchResult = query.data.getFlightByToAndFrom
        console.table(this.sharedService.flightSearchResult)
        await this.router.navigate(["/pesawat"])
      })
    }
  }

}
