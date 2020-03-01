import { Component, OnInit } from '@angular/core';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { timer } from 'rxjs';
import { Time } from '@angular/common';
import { getTreeMissingMatchingNodeDefError } from '@angular/cdk/tree';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { PesawatWidgetComponent } from './pesawat-widget/pesawat-widget.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-pesawat',
  templateUrl: './pesawat.component.html',
  styleUrls: ['./pesawat.component.sass']
})
export class PesawatComponent implements OnInit {

  sortBy: string = "rekomendasi"
  flightData: Object[]
  displayedFlightData: Object[]
  test: Object
  tempFlightData: Object[]
  flightSearchData: FlightSearchData
  departureTimes:TimeFilter[]=[
    {
      "label":"00:00-06:00",
      "checked":false,
      "start":{
        hours:0,
        minutes:0
      },
      "end":{
        hours:6,
        minutes:0
      }
    },
    {
      "label":"06:00-12:00",
      "checked":false,
      "start":{
        hours:6,
        minutes:0
      },
      "end":{
        hours:12,
        minutes:0
      }
    },
    {
      "label":"12:00-18:00",
      "checked":false,
      "start":{
        hours:12,
        minutes:0
      },
      "end":{
        hours:18,
        minutes:0
      }
    },
    {
      "label":"18:00-24:00",
      "checked":false,
      "start":{
        hours:18,
        minutes:0
      },
      "end":{
        hours:24,
        minutes:0
      }
    }
  ]
  arrivalTimes: TimeFilter[] = [
    {
      "label":"00:00-06:00",
      "checked":false,
      "start":{
        hours:0,
        minutes:0
      },
      "end":{
        hours:6,
        minutes:0
      }
    },
    {
      "label":"06:00-12:00",
      "checked":false,
      "start":{
        hours:6,
        minutes:0
      },
      "end":{
        hours:12,
        minutes:0
      }
    },
    {
      "label":"12:00-18:00",
      "checked":false,
      "start":{
        hours:12,
        minutes:0
      },
      "end":{
        hours:18,
        minutes:0
      }
    },
    {
      "label":"18:00-24:00",
      "checked":false,
      "start":{
        hours:18,
        minutes:0
      },
      "end":{
        hours:24,
        minutes:0
      }
    }
  ]
  airlineList: AirlineFilter[] = []
  facilityFilter: FacilityFilter[] = []
  transitFilter: TransitFilter[] = [
    {
      label: "Langsung",
      checked: false,
      amount: 0
    },
    {
      label: "1 Transit",
      checked: false,
      amount: 1
    },
    {
      label: "2+ Transit",
      checked: false,
      amount: 2
    }
  ]
  showData = 5
  detailString: string[] = [];
  facilityString: string[] = [];

  config: MatDialogConfig = new MatDialogConfig();

  constructor(
    private eventEmitterService: EventEmitterService,
    private sharedService: SharedServiceService,
    private dialog: MatDialog,
    private router: Router
  ) { }

  public sort() {
    console.log(this.sortBy)
    if(this.sortBy.localeCompare("rekomendasi") == 0 || this.sortBy.localeCompare("harga-terendah")) {
      this.sortByPrice()
    }
    if(this.sortBy.localeCompare("keberangkatan-awal") == 0) {
      this.sortDepartureByAscending()
    }
    if(this.sortBy.localeCompare("keberangkatan-akhir") == 0) {
      this.sortDepartureByDescending()
    }
    if(this.sortBy.localeCompare("kedatangan-awal") == 0) {
      this.sortArrivalByAscending()
    }
    if(this.sortBy.localeCompare("kedatangan-akhir") == 0) {
      this.sortArrivalByDescending()
    }
    if(this.sortBy.localeCompare("durasi") == 0) {
      this.sortByDuration()
    }
  }

  public sortByDuration() {
    var length = this.flightData.length
    this.displayedFlightData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.flightData[j].duration > this.flightData[j+1].duration) {
          var temp = this.flightData[j]
          this.flightData[j] = this.flightData[j+1]
          this.flightData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public sortArrivalByAscending() {
    var length = this.flightData.length
    this.displayedFlightData = []
    for (let i = 0; i < length-1; i++) { 
      for (let j = 0; j < length-i-1; j++) {
        if(new Date(this.flightData[j].arrival).getUTCHours() > new Date(this.flightData[j+1].arrival).getUTCHours()) {
          var temp = this.flightData[j]
          this.flightData[j] = this.flightData[j+1]
          this.flightData[j+1] = temp
        } else if(new Date(this.flightData[j].arrival).getUTCHours() == new Date(this.flightData[j+1].arrival).getUTCHours()) {
          if(new Date(this.flightData[j].arrival).getUTCMinutes() > new Date(this.flightData[j+1].arrival).getUTCMinutes()) {
            var temp = this.flightData[j]
            this.flightData[j] = this.flightData[j+1]
            this.flightData[j+1] = temp          
          }
        }
      }
    }
    this.loadData()    
  }

  public sortArrivalByDescending() {
    var length = this.flightData.length
    this.displayedFlightData = []
    for (let i = 0; i < length-1; i++) { 
      for (let j = 0; j < length-i-1; j++) {
        if(new Date(this.flightData[j].arrival).getUTCHours() < new Date(this.flightData[j+1].arrival).getUTCHours()) {
          var temp = this.flightData[j]
          this.flightData[j] = this.flightData[j+1]
          this.flightData[j+1] = temp
        } else if(new Date(this.flightData[j].arrival).getUTCHours() == new Date(this.flightData[j+1].arrival).getUTCHours()) {
          if(new Date(this.flightData[j].arrival).getUTCMinutes() < new Date(this.flightData[j+1].arrival).getUTCMinutes()) {
            var temp = this.flightData[j]
            this.flightData[j] = this.flightData[j+1]
            this.flightData[j+1] = temp          
          }
        }
      }
    }
    this.loadData()    
  }

  public sortDepartureByDescending() {
    var length = this.flightData.length
    this.displayedFlightData = []
    for (let i = 0; i < length-1; i++) { 
      for (let j = 0; j < length-i-1; j++) {
        if(new Date(this.flightData[j].departure).getUTCHours() < new Date(this.flightData[j+1].departure).getUTCHours()) {
          var temp = this.flightData[j]
          this.flightData[j] = this.flightData[j+1]
          this.flightData[j+1] = temp
        } else if(new Date(this.flightData[j].departure).getUTCHours() == new Date(this.flightData[j+1].departure).getUTCHours()) {
          if(new Date(this.flightData[j].departure).getUTCMinutes() < new Date(this.flightData[j+1].departure).getUTCMinutes()) {
            var temp = this.flightData[j]
            this.flightData[j] = this.flightData[j+1]
            this.flightData[j+1] = temp          
          }
        }
      }
    }
    this.loadData()    
  }

  public sortDepartureByAscending() {
    var length = this.flightData.length
    this.displayedFlightData = []
    for (let i = 0; i < length-1; i++) { 
      for (let j = 0; j < length-i-1; j++) {
        if(new Date(this.flightData[j].departure).getUTCHours() > new Date(this.flightData[j+1].departure).getUTCHours()) {
          var temp = this.flightData[j]
          this.flightData[j] = this.flightData[j+1]
          this.flightData[j+1] = temp
        } else if(new Date(this.flightData[j].departure).getUTCHours() == new Date(this.flightData[j+1].departure).getUTCHours()) {
          if(new Date(this.flightData[j].departure).getUTCMinutes() > new Date(this.flightData[j+1].departure).getUTCMinutes()) {
            var temp = this.flightData[j]
            this.flightData[j] = this.flightData[j+1]
            this.flightData[j+1] = temp          
          }
        }
      }
    }
    this.loadData()    
  }

  public sortByPrice() {
    var length = this.flightData.length
    this.displayedFlightData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.flightData[j].price > this.flightData[j+1].price) {
          var temp = this.flightData[j]
          this.flightData[j] = this.flightData[j+1]
          this.flightData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public filterAirline() {
    var airlineNameList: string[] = [];
    this.flightData.forEach(element => {
      airlineNameList.push(element.Airline.Name)
    })
    var filteredAirline = airlineNameList.filter(
      (airlineName, i, arr) => arr.findIndex( name => name === airlineName ) === i
    )
    filteredAirline.forEach(element => {
      this.airlineList.push({
        airline: element,
        checked: false
      })
    })
  }

  public filterFacility() {
    var facilityNameList: string[] = [];
    this.flightData.forEach(element => {
      element.Facilities.forEach(facility => {
        facilityNameList.push(facility.FacilityName)
      })
    })
    var filteredAirline = facilityNameList.filter(
      (airlineName, i, arr) => arr.findIndex( name => name === airlineName ) === i
    )
    filteredAirline.forEach(element => {
      this.facilityFilter.push({
        facility: element,
        checked: false
      })
    })
  }

  ngOnInit() {
    this.displayedFlightData = []
    this.config.autoFocus = false
    this.config.restoreFocus = true
    this.flightSearchData = {to: "", from: ""}
    this.flightSearchData.from = this.sharedService.flightSearchData.from
    this.flightSearchData.to = this.sharedService.flightSearchData.to
    this.flightData = this.sharedService.flightSearchResult
    this.sort()
    this.filterAirline()
    this.filterFacility()
    window.addEventListener('scroll', this.scroll, true)
  }

  public loadData() {
    for (let index = 0; index < this.flightData.length; index++) {
      this.detailString.push("")      
      this.facilityString.push("")
    }
    if(this.flightData.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedFlightData.push(this.flightData[index])
      }
      // console.log(this.displayedFlightData)
    }
  }

  public showDetail(type: string, i: number) {
    if(this.detailString[i] == type) {
      this.detailString[i] = ''
    } else {
      this.detailString[i] = type
    }
  }

  public changeSearch() {
    this.dialog.open(PesawatWidgetComponent, this.config)
  }

  public checkValidation(idx: number) {
    if(!this.checkDepartureFilter(idx)) {
      return false
    } else if(!this.checkArrivalFilter(idx)) {
      return false
    } else if(!this.checkAirlineFilter(idx)) {
      return false
    } else if(!this.checkTransitFilter(idx)) {
      return false
    } else if(!this.checkFacilityFilter(idx)) {
      return false
    }
    return true
  }

  public resetFilter() {
    for (let index = 0; index < this.transitFilter.length; index++) {
      this.transitFilter[index].checked = false
    }
    for (let index = 0; index < this.airlineList.length; index++) {
      this.airlineList[index].checked = false
    }
    for (let index = 0; index < this.departureTimes.length; index++) {
      this.departureTimes[index].checked = false      
    }
    for (let index = 0; index < this.arrivalTimes.length; index++) {
      this.arrivalTimes[index].checked = false      
    }
    for (let index = 0; index < this.facilityFilter.length; index++) {
      this.facilityFilter[index].checked = false
    }
  }

  public checkFacilityFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.facilityFilter.length; index++) {
      if(this.facilityFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.facilityFilter.length; index++) {
        if(this.facilityFilter[index].checked == true) {
          for (let index2 = 0; index2 < this.flightData[idx].facilities.length; index2++) {
            if(this.facilityFilter[index].facility == this.flightData[idx].facilities[index2].FacilityName) {
              return true;
            } else continue
          }
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkTransitFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.transitFilter.length; index++) {
      if(this.transitFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.transitFilter.length; index++) {
        if(this.transitFilter[index].checked == true) {
          if(this.transitFilter[index].amount == 2) {
            if(this.flightData[idx].routes.length >= 2) {
              return true
            } else continue
          }
          else {
            if(this.flightData[idx].routes.length == this.transitFilter[index].amount) {
              return true
            } else continue
          }
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkAirlineFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.airlineList.length; index++) {
      if(this.airlineList[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.airlineList.length; index++) {
        if(this.airlineList[index].checked == true) {
          if(this.flightData[idx].airline.name == this.airlineList[index].airline) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkArrivalFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.arrivalTimes.length; index++) {
      if(this.arrivalTimes[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.arrivalTimes.length; index++) {
        if(this.arrivalTimes[index].checked == true) {
          var arrivalHour = new Date(this.flightData[idx].arrival).getUTCHours()
          var arrivalMinute = new Date(this.flightData[idx].arrival).getUTCMinutes()
          if(arrivalHour < this.departureTimes[index].start.hours || arrivalHour > this.departureTimes[index].end.hours) {
            continue
          }
          if(arrivalMinute < this.departureTimes[index].start.minutes || arrivalMinute > this.departureTimes[index].end.minutes) {
            continue
          }
          return true
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkDepartureFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.departureTimes.length; index++) {
      if(this.departureTimes[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.departureTimes.length; index++) {
        if(this.departureTimes[index].checked == true) {
          var departureHour = new Date(this.flightData[idx].departure).getUTCHours()
          var departureMinute = new Date(this.flightData[idx].departure).getUTCMinutes()
          if(departureHour < this.departureTimes[index].start.hours || departureHour > this.departureTimes[index].end.hours) {
            continue
          }
          if(departureMinute < this.departureTimes[index].start.minutes || departureMinute > this.departureTimes[index].end.minutes) {
            continue
          }
          return true
        }
      }
      return false
    } else {
      return true
    }
  }

  public toCheckout(idx: number) {
    localStorage.setItem("whichBought", "flight")
    this.sharedService.flightBought = this.displayedFlightData[idx]
    this.router.navigate(['/checkout'])
  }

  scroll = (event): void => {
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 5
      if(this.flightData.length >= this.showData) {
        for (let index = this.showData-5; index < this.showData; index++) {
          this.displayedFlightData.push(this.flightData[index])
        }
      } else {
        for (let index = this.showData-5; index < this.flightData.length; index++) {
          this.displayedFlightData.push(this.flightData[index])
        }        
      }
    }
  }

}

export class TimeFilter {
  start: Time
  end: Time
  label: string
  checked: boolean
}

export class AirlineFilter {
  airline: string
  checked: boolean
}

export class TransitFilter {
  label: string
  checked: boolean
  amount: number
}

export class FacilityFilter {
  facility: string
  checked: boolean
}

export class FlightSearchData {
  to: string
  from: string
}