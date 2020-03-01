import { Component, OnInit } from '@angular/core';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Time } from '@angular/common';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

@Component({
  selector: 'app-kereta-api',
  templateUrl: './kereta-api.component.html',
  styleUrls: ['./kereta-api.component.sass']
})
export class KeretaApiComponent implements OnInit {

  trainData: Object[]
  displayedTrainData: Object[]
  test: Object
  tempTrainData: Object[]
  trainSearchData: TrainSearchData
  fromValue: string
  toValue: string
  fromLocationToView: Object[] = []
  toLocationToView: Object[] = []

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
  kelasList: ClassFilter[] = [
    {
      class: "Ekonomi",
      checked: false
    },
    {
      class: "Bisnis",
      checked: false,
    },
    {
      class: "Eksekutif",
      checked: false
    }
  ]
  nameFilter: TrainNameFilter[] = []
  showData = 5
  detailString: string[] = [];

  config: MatDialogConfig = new MatDialogConfig();

  constructor(
    private eventEmitterService: EventEmitterService,
    private sharedService: SharedServiceService,
    private dialog: MatDialog,
    private graphqlService: GraphqlServiceService
  ) { }

  public filterTrain() {
    var trainNameList: string[] = [];
    this.trainData.forEach(element => {
      trainNameList.push(element.Train.Name)
    })
    var fliteredTrain = trainNameList.filter(
      (trainName, i, arr) => arr.findIndex( name => name === trainName ) === i
    )
    fliteredTrain.forEach(element => {
      this.nameFilter.push({
        name: element,
        checked: false
      })
    })
  }

  ngOnInit() {
    this.displayedTrainData = []
    this.config.autoFocus = false
    this.config.restoreFocus = true
    this.trainSearchData = {to: "", from: ""}
    this.trainSearchData.from = this.sharedService.trainSearchData.from
    this.trainSearchData.to = this.sharedService.trainSearchData.to
    this.fromValue = this.trainSearchData.from
    this.toValue = this.trainSearchData.to
    this.fromLocationToView = ['Surabaya', 'Jakarta Pusat']
    this.toLocationToView = ['Surabaya', 'Jakarta Pusat']
    // this.graphqlService.getAllLocation().subscribe(async query => {
      // this.locations = query.data.getAllLocation
      // for (let idx = index; idx < count; idx++) {
      //   this.fromLocationToView.push(this.locations[idx])
      //   this.toLocationToView.push(this.locations[idx])
      // }
      // this.fromLocationToView = query.data.getAllLocation
      // this.toLocationToView = query.data.getAllLocation
    // })
    this.loadData()
    window.addEventListener('scroll', this.scroll, true)
  }

  public loadData() {
    this.trainData = this.sharedService.trainSearchResult
    for (let index = 0; index < this.trainData.length; index++) {
      this.detailString.push("")
    }
    if(this.trainData.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedTrainData.push(this.trainData[index])
      }
    } else {
      for (let index = 0; index < this.trainData.length; index++) {
        this.displayedTrainData.push(this.trainData[index])
      }
    }
    this.filterTrain()
  }

  public showDetail(type: string, i: number) {
    if(this.detailString[i] == type) {
      this.detailString[i] = ''
    } else {
      this.detailString[i] = type
    }
  }

  // public changeSearch() {
  //   this.dialog.open(PesawatWidgetComponent, this.config)
  // }

  public checkValidation(idx: number) {
    if(!this.checkDepartureFilter(idx)) {
      return false
    } else if(!this.checkClassFilter(idx)) {
      return false
    } else if(!this.checkTrainFilter(idx)) {
      return false
    }
    return true
  }

  public resetFilter() {
    for (let index = 0; index < this.departureTimes.length; index++) {
      this.departureTimes[index].checked = false      
    }
    for (let index = 0; index < this.kelasList.length; index++) {
      this.kelasList[index].checked = false      
    }
    for (let index = 0; index < this.nameFilter.length; index++) {
      this.nameFilter[index].checked = false
    }
  }

  public checkTrainFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.nameFilter.length; index++) {
      if(this.nameFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.nameFilter.length; index++) {
        if(this.nameFilter[index].checked == true) {
          if(this.trainData[idx].train.name == this.nameFilter[index].name) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkClassFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.kelasList.length; index++) {
      if(this.kelasList[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.kelasList.length; index++) {
        if(this.kelasList[index].checked == true) {
          if(this.trainData[idx].train.class == this.kelasList[index].class) {
            return true;
          } else continue;
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
          var departureHour = new Date(this.trainData[idx].departure).getUTCHours()
          var departureMinute = new Date(this.trainData[idx].departure).getUTCMinutes()
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
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 5
      if(this.trainData.length >= this.showData) {
        for (let index = this.showData-5; index < this.showData; index++) {
          this.displayedTrainData.push(this.trainData[index])
        }
      } else {
        for (let index = this.showData-5; index < this.trainData.length; index++) {
          this.displayedTrainData.push(this.trainData[index])
        }        
      }
    }
  }

  changeSearch() {
    this.displayedTrainData = []
    this.detailString = []
    this.graphqlService.getTripByToAndFrom(this.toValue, this.fromValue).subscribe(async query => {
      this.sharedService.trainSearchResult = query.data.getTripByToAndFrom
      this.loadData()
    })

  }

}

export class TimeFilter {
  start: Time
  end: Time
  label: string
  checked: boolean
}

export class ClassFilter {
  class: string
  checked: boolean
}

export class TrainNameFilter {
  name: string
  checked: boolean
}

export class TrainSearchData {
  to: string
  from: string
}