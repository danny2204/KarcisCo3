import { Component, OnInit } from '@angular/core';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { timer } from 'rxjs';
import { Time } from '@angular/common';
import { getTreeMissingMatchingNodeDefError } from '@angular/cdk/tree';

@Component({
  selector: 'app-pesawat',
  templateUrl: './pesawat.component.html',
  styleUrls: ['./pesawat.component.sass']
})
export class PesawatComponent implements OnInit {

  flightData: Object[]
  displayedFlightData: Object[]
  test: Object
  tempFlightData: Object[]
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
  airlineList: AirlineFilter[] = [
    {
      airline: "Batik Air",
      checked: false
    }, 
    {
      airline: "Citilink Indonesia",
      checked: false
    },
    {
      airline: "Garuda Indonesia",
      checked: false
    },
    {
      airline: "Lion Air",
      checked: false
    }
  ]
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


  constructor(
    private eventEmitterService: EventEmitterService,
    private sharedService: SharedServiceService
  ) { }

  ngOnInit() {
    window.addEventListener('scroll', this.scroll, true)
    this.displayedFlightData = []
    this.flightData = this.sharedService.flightSearchResult
    if(this.flightData.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedFlightData.push(this.flightData[index])
      }
    }
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

  scroll = (event): void => {
    console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 5
      if(this.flightData.length >= this.showData) {
        for (let index = this.showData-5; index < this.showData; index++) {
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