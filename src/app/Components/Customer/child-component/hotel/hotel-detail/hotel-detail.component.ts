import { Component, OnInit } from '@angular/core';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { PopUpImageComponent } from './pop-up-image/pop-up-image.component';
import { Router } from '@angular/router';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import * as L from 'leaflet';

@Component({
  selector: 'app-hotel-detail',
  templateUrl: './hotel-detail.component.html',
  styleUrls: ['./hotel-detail.component.sass']
})
export class HotelDetailComponent implements OnInit {

  test: string = "url('../../../../../../assets/house.png')"
  hotel: Object
  config: MatDialogConfig = new MatDialogConfig();
  nearestHotel: Object[] = []
  markerList: L.marker[] = []
  map: L.map


  constructor(
    private sharedService: SharedServiceService,
    private dialog: MatDialog,
    private router: Router,
    private graphqlService: GraphqlServiceService
  ) { }

  currentUrl: string
  
  sendThroughEmail() {
    location.href = 'mailto:?subject=Tiket.com&body=links:http://127.0.0.1:4200' + this.currentUrl;
  }

  copyToClipboard() {
    const selBox = document.createElement('textarea')
    selBox.style.opacity = '0'
    selBox.value = this.currentUrl
    document.body.appendChild(selBox)
    selBox.focus()
    selBox.select()
    document.execCommand("copy")
    document.body.removeChild(selBox)
  }

  ngOnInit() {
    this.currentUrl = window.location.href
    console.table(this.sharedService.selectedHotel)
    this.hotel = this.sharedService.selectedHotel
    this.config.autoFocus = false;
    this.config.restoreFocus = true;

    this.getNearestHotel(this.hotel.Longitude, this.hotel.Latitude)

  }

  getNearestHotel(longitude: number, latitude: number) {
    this.graphqlService.getHotelByGeolocation(longitude, latitude).subscribe(async query => {
      this.nearestHotel = query.data.getHotelByGeolocation
      console.table(this.nearestHotel)
      await this.setMarker()
    })
  }

  setMarker() {
    this.map = L.map('map').setView([this.hotel.Latitude, this.hotel.Longitude], 13);
    var markersLayer = L.featureGroup().addTo(this.map)
    this.markerList = []

    L.marker([this.hotel.Latitude, this.hotel.Longitude]).addTo(markersLayer).bindPopup(this.hotel.Name)

    for(let i = 0; i < this.nearestHotel.length; i++) {
      var firstPointer = L.marker([this.nearestHotel[i].Latitude, this.nearestHotel[i].Longitude])
      firstPointer.addTo(markersLayer).bindPopup(this.nearestHotel[i].Name).openPopup();
    }

    console.table(this.nearestHotel)

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(this.map);

  }

  public openDialog() {
    this.dialog.open(PopUpImageComponent, this.config);
  }

  public toCheckout(idx: number) {
    localStorage.setItem("whichBought", "hotel")
    this.sharedService.hotelBought = this.hotel
    this.sharedService.hotelBought.Price = this.hotel.Room[idx].Price
    this.sharedService.roomChoosed = this.hotel.Room[idx]
    this.router.navigate(['/checkout'])
  }

}
