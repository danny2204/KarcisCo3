import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SharedServiceService {

  public flightData: Object[]
  public flightSearchResult: Object[]
  public hotelSearchResult: Object[]

  fetchFlight(flight) {
    this.flightSearchResult = this.flightData
  }

}
