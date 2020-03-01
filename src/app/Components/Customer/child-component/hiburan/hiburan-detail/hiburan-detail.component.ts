import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import * as L from 'leaflet';


@Component({
  selector: 'app-hiburan-detail',
  templateUrl: './hiburan-detail.component.html',
  styleUrls: ['./hiburan-detail.component.sass']
})
export class HiburanDetailComponent implements OnInit {

  Event: Object = {}
  eventTicketList: EventTicket[] = []
  map: L.map

  constructor(
    private graphqlService: GraphqlServiceService
  ) { }

  ngOnInit() {
    this.graphqlService.getEntertainmentById(parseInt(localStorage.getItem("selectedEvent"))).subscribe(async query => {
      this.Event = query.data.getEntertainmentById

      this.map = L.map('map').setView([this.Event.Location.Latitude, this.Event.Location.Longitude], 13);
      var markersLayer = L.featureGroup().addTo(this.map)

      var firstPointer = L.marker([this.Event.Location.Latitude, this.Event.Location.Longitude])
      firstPointer.addTo(markersLayer).bindPopup(this.Event.Name);

      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
          attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
      }).addTo(this.map);

      for(let i = 0; i < this.Event.EntertainmentTicket.length; i++) {
        this.eventTicketList.push({
          Description: this.Event.EntertainmentTicket[i].Description,
          Price: this.Event.EntertainmentTicket[i].Price,
          Quantity: 0,
          Type: this.Event.EntertainmentTicket[i].Type
        })
      }
    })
  }

  allTicketQuantity() {
    let sum = 0
    for(let i = 0 ; i < this.eventTicketList.length; i++) {
      sum += this.eventTicketList[i].Quantity
    }
    return sum
  }

  getTotalPrice() {
    let totalPrice = 0
    for(let i = 0; i < this.eventTicketList.length; i++) {
      totalPrice += (this.eventTicketList[i].Price * this.eventTicketList[i].Quantity)
    }
    return totalPrice
  }

  getLowestPrice() {
    let minimumPrice = 9999999
    for(let i = 0 ; i < this.Event.EntertainmentTicket.length; i++) {
      if(this.Event.EntertainmentTicket[i].Price <  minimumPrice) {
        minimumPrice = this.Event.EntertainmentTicket[i].Price
      }
    }
    return minimumPrice
  }

  goToDetail() {

  }

}

export interface EventTicket {
  Price: number
  Quantity: number
  Type: string
  Description: string
}
