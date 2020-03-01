import { Component, OnInit } from '@angular/core';
import * as L from 'leaflet';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hotel-map',
  templateUrl: './hotel-map.component.html',
  styleUrls: ['./hotel-map.component.sass']
})
export class HotelMapComponent implements OnInit {

  markerList: L.marker[] = []
  hotelList: Object[]
  map: L.map
  selectedHotel: Object = {}

  constructor(
    private sharedService: SharedServiceService,
    private router: Router
  ) { }

  ngOnInit() {
    this.map = L.map('map').setView([51.505, -0.09], 13);
    var markersLayer = L.featureGroup().addTo(this.map)
    this.hotelList = this.sharedService.hotelSearchResult
    if(localStorage.getItem("openMapFrom") == "hotel-page") {
      this.markerList = []
      for(let idx = 0; idx < this.hotelList.length; idx++) {
        var firstPointer = L.marker([this.hotelList[idx].Latitude, this.hotelList[idx].Longitude])
        firstPointer.Name = this.hotelList[idx].Name
        firstPointer.addTo(markersLayer).bindPopup(this.hotelList[idx].Name).openPopup();
      }
    } else if(localStorage.getItem("openMapFrom") == "hotel-detail") {
      this.markerList = []
      var selectedHotel = this.sharedService.selectedHotel
      var pointer = L.marker([selectedHotel.Longitude, selectedHotel.Latitude])

      this.markerList.push(this.sharedService.selectedHotel)
    }

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(this.map);

    markersLayer.on("click", this.onclick)
  }

  public changeView(idx: number) {
    this.map.setView([this.hotelList[idx].Latitude, this.hotelList[idx].Longitude], 13)
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(this.map);
    document.getElementById("right-panel").style.display = "flex"
    this.selectedHotel = this.hotelList[idx]
  }

  public toDetail() {
    this.sharedService.selectedHotel = this.selectedHotel
    this.router.navigate(['/hotel-detail'])
  }

  public onclick(event) {
    console.log(event.layer.test)
  }

  public getLowestPrice() {
    var minimalPrice = 9999999
    for(let idx = 0; idx < this.selectedHotel.Room.length; idx++) {
      if(this.selectedHotel.Room[idx].Price < minimalPrice) {
        minimalPrice = this.selectedHotel.Room[idx].Price
      }
    }
    return minimalPrice
  }

}
