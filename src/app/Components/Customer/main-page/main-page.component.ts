import { Component, OnInit, AfterViewInit, AfterContentInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.sass']
})
export class MainPageComponent implements AfterContentInit {

  constructor(
    private graphqlService: GraphqlServiceService,
  ) { }

  isLoading: boolean = true
  nearestHotel: Object[] = []
  subscriber1: Subscription
  ratingList: number[] = [1, 2, 3, 4, 5]

  isValidRating(num: number, rating: number) {
    if(num <= rating) return true
    return false
  }

  loadLowestPrice(hotel: Object) {
    var minimalPrice = 9999999
    for(let i = 0; i < hotel.Room.length; i++){
      if(hotel.Room[i].Price < minimalPrice) {
        minimalPrice = hotel.Room[i].Price
      }
    }
    return minimalPrice
  }

  hotelLoaded(){
    this.isLoading = false
  }

  getNearestHotel(longitude: number, latitude: number) {
    this.subscriber1 = this.graphqlService.getHotelByGeolocation(longitude, latitude).subscribe(async query => {
      this.nearestHotel = query.data.getHotelByGeolocation
      await this.hotelLoaded()
    })
  }

  ngAfterContentInit() {
    console.log(localStorage.getItem("history"))
    this.isLoading = true
    navigator.geolocation.getCurrentPosition(async (resp) => {
      console.log("Longitude" + resp.coords.longitude)
      console.log("Latitude" + resp.coords.latitude)
      setTimeout(() => {
        this.getNearestHotel(resp.coords.longitude, resp.coords.latitude)
      }, 3000);
    })
  }
}