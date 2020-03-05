import { Component, OnInit, ViewChild, TemplateRef } from '@angular/core';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Time, DatePipe } from '@angular/common';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { CalendarEventTimesChangedEvent, CalendarEvent, CalendarView } from 'angular-calendar';
import { isSameMonth, isSameDay, startOfDay, endOfDay } from 'date-fns';
import { Subject } from 'rxjs';
import { WebSocketServiceService } from 'src/app/Service/web-socket-service.service';

@Component({
  selector: 'app-kereta-api',
  templateUrl: './kereta-api.component.html',
  styleUrls: ['./kereta-api.component.sass'],
  providers: [DatePipe]
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
    private graphqlService: GraphqlServiceService,
    private datePipe: DatePipe,
    private webSocket: WebSocketServiceService
  ) { }

  addWebSocket() {
    this.webSocket.listen("train").subscribe(async data => {
      alert("new train inserted ! Refresh the page !")
    })
  }

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
    this.addWebSocket()
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
    var fl = []
    var dateExist = false
    for(let i = 0 ; i < this.trainData.length; i++) {
      dateExist = false
      for(let j = 0 ; j < fl.length; j++){
        if(this.datePipe.transform(fl[j].Date, "yyyy-MM-dd") == this.datePipe.transform(this.trainData[i].Departure, "yyyy-MM-dd")) {
          dateExist = true
          if(fl[j].Price > this.trainData[i].Price) {
            fl[j].Price = this.trainData[i].Price
          }
          break
        }
      }
      if(dateExist == false) {
        fl.push({
          "Date": this.trainData[i].Departure,
          "Price": this.trainData[i].Price
        })
      }
    }

    for(let i = 0 ; i < fl.length; i++) {
      this.events.push({
        start: startOfDay(new Date(fl[i].Date)),
        title: fl[i].Price,
      })
    }

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
          if(this.trainData[idx].Train.Name == this.nameFilter[index].name) return true
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
          if(this.trainData[idx].Train.Class == this.kelasList[index].class) {
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
          var departureHour = new Date(this.trainData[idx].Departure).getUTCHours()
          var departureMinute = new Date(this.trainData[idx].Departure).getUTCMinutes()
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

  openCalendar() {
    document.getElementById('calendar').style.zIndex = (parseInt(document.getElementById('calendar').style.zIndex)*-1).toString()
  }

  @ViewChild('modalContent', { static: true }) modalContent: TemplateRef<any>;

  view: CalendarView = CalendarView.Month;

  CalendarView = CalendarView;

  viewDate: Date = new Date();

  modalData: {
    action: string;
    event: CalendarEvent;
  };

  refresh: Subject<any> = new Subject();

  events: CalendarEvent[] = [
    // {
    //   start: startOfDay(new Date()),
    //   title: 'An event with no end date',
    //   color: colors.yellow,
    // },
  ];

  activeDayIsOpen: boolean = true;

  dayClicked({ date, events }: { date: Date; events: CalendarEvent[] }): void {
    if (isSameMonth(date, this.viewDate)) {
      if (
        (isSameDay(this.viewDate, date) && this.activeDayIsOpen === true) ||
        events.length === 0
      ) {
        this.activeDayIsOpen = false;
      } else {
        this.activeDayIsOpen = true;
      }
      this.viewDate = date;
    }
  }

  eventTimesChanged({
    event,
    newStart,
    newEnd
  }: CalendarEventTimesChangedEvent): void {
    this.events = this.events.map(iEvent => {
      if (iEvent === event) {
        return {
          ...event,
          start: newStart,
          end: newEnd
        };
      }
      return iEvent;
    });
  }

  addEvent(): void {
    this.events = [
      ...this.events,
      {
        title: 'New event',
        start: startOfDay(new Date()),
        end: endOfDay(new Date()),
        color: colors.red,
        draggable: true,
        resizable: {
          beforeStart: true,
          afterEnd: true
        }
      }
    ];
  }


  setView(view: CalendarView) {
    this.view = view;
  }

  closeOpenMonthViewDay() {
    this.activeDayIsOpen = false;
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