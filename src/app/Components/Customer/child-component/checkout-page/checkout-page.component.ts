import { Component, OnInit, AfterViewInit } from '@angular/core';
import { SharedServiceService } from 'src/app/Service/shared-service.service';

@Component({
  selector: 'app-checkout-page',
  templateUrl: './checkout-page.component.html',
  styleUrls: ['./checkout-page.component.sass']
})
export class CheckoutPageComponent implements OnInit {

  flightData: Object = {}
  hotelData: Object = {}
  trainData: Object = {}
  carData: Object = {}
  entertainmentData: Object = {}
  roomData: Object = {}
  vendorData: Object = {}

  isFlightBought: boolean = false
  isTrainBought: boolean = false
  isCarBought: boolean = false
  isHotelBought: boolean = false
  isEntertainmentBought: boolean = false

  constructor(
    private sharedService: SharedServiceService
  ) { }

  ngOnInit() {
    if(localStorage["whichBought"] == "flight") {
      this.flightData = []
      this.isFlightBought = true
      this.flightData = this.sharedService.flightBought
      console.log(this.flightData)
    }
    else if(localStorage["whichBought"] == "train") {
      this.isTrainBought = true
    }
    else if(localStorage["whichBought"] == "car") {
      this.isCarBought = true
      this.carData = this.sharedService.carBought
      this.vendorData = this.sharedService.vendorChoosed
    }
    else if(localStorage["whichBought"] == "hotel") {
      this.isHotelBought = true
      this.hotelData = this.sharedService.hotelBought
      this.roomData = this.sharedService.roomChoosed
    }
    else if(localStorage["whichBought"] == "entertainment") {
      this.isEntertainmentBought = true
    }
  }

}
