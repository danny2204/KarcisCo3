import { Component, OnInit, AfterViewInit } from '@angular/core';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';

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

  minutes: number = 0
  seconds: number = 10

  constructor(
    private sharedService: SharedServiceService,
    private router: Router
  ) { }

  titelValue: string = ""
  nameValue: string = ""
  emailValue: string = ""

  toPayment() {
    if(this.titelValue == "") {
      alert("titel cannot be empty")
    }
    else if(this.nameValue == "") {
      alert("name cannot be empty")
    }
    else if(this.emailValue == "") {
      alert("email cannot be empty")
    }
  }

  ngOnInit() {

    var interval = setInterval(() => {
      this.seconds -= 1
      if(this.seconds == 0 && this.minutes != 0) {
        this.minutes -= 1
        this.seconds = 59
      }
      if(this.minutes <= 0 && this.seconds <= 0) {
        alert("times out")
        clearInterval(interval)
        this.router.navigate(['/'])
      }
    }, 1000)

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
