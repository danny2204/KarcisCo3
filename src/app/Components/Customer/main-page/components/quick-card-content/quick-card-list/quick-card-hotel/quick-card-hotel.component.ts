import { Component, OnInit, ViewChild } from '@angular/core';
import { Subscription } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { Router } from '@angular/router';
import { MatSelect } from '@angular/material/select';

@Component({
  selector: 'app-quick-card-hotel',
  templateUrl: './quick-card-hotel.component.html',
  styleUrls: ['./quick-card-hotel.component.sass']
})
export class QuickCardHotelComponent implements OnInit {

  @ViewChild('locationSelect') child: MatSelect;
  hotels: Object[]
  locations: Object[]
  supkrep: Subscription
  locationValue: string
  supkrep2: Subscription
  locationToView: Object[]
  amount = 10

  constructor(
    private service: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private emitterService: EventEmitterService,
    private router: Router
  ) { }
    
    init() {
      this.child._openedStream.subscribe(() => {
        this.child.panel.nativeElement.addEventListener('scroll', event => {
          if(event.target.scrollTop+256 >= event.target.scrollHeight) {
            this.amount += 5
            this.addItems(this.amount)
          }
        })
      })
      // this.addItems(this.amount)
    }

  ngOnInit() {
    this.locationToView = ["Jakarta Pusat"]
    // this.supkrep = this.service.getAllLocation().subscribe(async query => {
    //   this.locations = query.data.getAllLocation
    //   await this.init()
    // })
  }

  public getHotelByLocation() {
    this.supkrep2 = this.service.getHotelByCity(this.locationValue).subscribe(async query => {
      this.sharedService.hotelSearchResult = query.data.getHotelByCity
      await this.router.navigate(["/hotel"])
    })
  }

  addItems(count: number) {
    let index = this.locationToView.length
    for (let idx = index; idx < count; idx++) {
      this.locationToView.push(this.locations[idx])
    }
  }

  loadMore(event) {
    console.log(event.scroll.scrollY)
    // if(event.scroll.scrollY)
  }

  // scroll = (event): void => {
  //   if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
  //     this.showData += 5
  //     if(this.flightData.length >= this.showData) {
  //       for (let index = this.showData-5; index < this.showData; index++) {
  //         this.displayedFlightData.push(this.flightData[index])
  //       }
  //     }
  //   }
  // }

}