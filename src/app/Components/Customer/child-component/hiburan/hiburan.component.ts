import { Component, OnInit } from '@angular/core';
import { Subscription, Subscribable } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';
import { WebSocketServiceService } from 'src/app/Service/web-socket-service.service';

@Component({
  selector: 'app-hiburan',
  templateUrl: './hiburan.component.html',
  styleUrls: ['./hiburan.component.sass']
})
export class HiburanComponent implements OnInit {

  trendingActivities: Object[] = []
  trendingAttractions: Object[] = []
  trendingEvents: Object[] = []
  typeValue : string

  subscriber1: Subscription
  subscriber2: Subscription
  subscriber3: Subscription
  subscriber4: Subscription

  constructor(
    private graphqlService: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private router: Router,
    private webSocket: WebSocketServiceService
  ) { }

  loadTrendAttractions() {
    this.subscriber2 = this.graphqlService.getEntertainmentByTrending("Attractions").subscribe(async query => {
      this.trendingAttractions = query.data.getEntertainmentByTrending
      await this.loadTrendEvents()
    })
  }

  loadTrendEvents() {
    this.subscriber3 = this.graphqlService.getEntertainmentByTrending("Events").subscribe(async query => {
      this.trendingEvents = query.data.getEntertainmentByTrending
    })
  }

  addWebSocket() {
    this.webSocket.listen("event").subscribe(async data => {
      alert("new event inserted ! Refresh the page !")
    })
  }

  ngOnInit() {
    this.addWebSocket()
    this.typeValue = ''
    this.subscriber1 = this.graphqlService.getEntertainmentByTrending("Activities").subscribe(async query => {
      this.trendingActivities = query.data.getEntertainmentByTrending
      await this.loadTrendAttractions()
    })
  }

  getActivitiesLowestPrice(idx: number) {
    var minimalPrice = 999999999
    console.log(this.trendingActivities.length)
    for (let index = 0; index < this.trendingActivities[idx].EntertainmentTicket.length; index++) {
      console.log(this.trendingActivities[idx].EntertainmentTicket[index].Price)
      if(minimalPrice > this.trendingActivities[idx].EntertainmentTicket[index].Price) {
        minimalPrice = this.trendingActivities[idx].EntertainmentTicket[index].Price
      }
    }
    return minimalPrice
  }

  getAttractionLowestPrice(idx: number) {
    var minimalPrice = 999999999
    console.log(this.trendingActivities.length)
    for (let index = 0; index < this.trendingAttractions[idx].EntertainmentTicket.length; index++) {
      console.log(this.trendingAttractions[idx].EntertainmentTicket[index].Price)
      if(minimalPrice > this.trendingAttractions[idx].EntertainmentTicket[index].Price) {
        minimalPrice = this.trendingAttractions[idx].EntertainmentTicket[index].Price
      }
    }
    return minimalPrice
  }

  getEventLowestPrice(idx: number) {
    var minimalPrice = 999999999
    console.log(this.trendingActivities.length)
    for (let index = 0; index < this.trendingEvents[idx].EntertainmentTicket.length; index++) {
      console.log(this.trendingEvents[idx].EntertainmentTicket[index].Price)
      if(minimalPrice > this.trendingEvents[idx].EntertainmentTicket[index].Price) {
        minimalPrice = this.trendingEvents[idx].EntertainmentTicket[index].Price
      }
    }
    return minimalPrice
  }
 
  searchEntertainmentByType() {
    this.sharedService.entertainmentTypeValue = this.typeValue
    this.router.navigate(['hiburan-search'])
  }

}
