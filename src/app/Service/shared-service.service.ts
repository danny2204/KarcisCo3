import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SharedServiceService {

  public flightData: Object[]
  public flightSearchResult: Object[]
  public hotelSearchResult: Object[]
  public trainSearchResult: Object[]
  public carSearchResult: Object[]
  public airports: Object[]
  public flightSearchData: FlightSearchData = {to: "", from: ""}
  public trainSearchData: TrainSearchData = {to: "", from: ""}
  public carSearchData: CarSearchData = {where: "", from: new Date(), to: new Date()}
  public vendorsData: Object[]
  public carData: Object = {}
  public selectedHotel: Object = {}

  public flightBought: Object = {}
  public hotelBought: Object = {}
  public carBought: Object = {}
  public trainBought: Object = {}
  public entertainmentBought: Object = {}
  public roomChoosed: Object = {}
  public vendorChoosed: Object = {}

  public entertainmentsData: Object = {}

  public detailPromoData: Object = {}

  public entertainmentTypeValue: string

  fetchFlight(flight) {
    this.flightSearchResult = this.flightData
  }

}

export class FlightSearchData {
  to: string
  from: string
}

export class TrainSearchData {
  to: string
  from: string
}

export class CarSearchData {
  where: string
  from: Date
  to: Date
}