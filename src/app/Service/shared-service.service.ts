import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SharedServiceService {

  public flightData: Object[]
  fetchFlight(flight) {
    this.flightData = flight
  }

}
