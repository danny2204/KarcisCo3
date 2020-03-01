import { Component, OnInit, ViewChild, AfterViewInit, ChangeDetectorRef } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Subscribable, Subscription } from 'rxjs';
import { async } from '@angular/core/testing';
import { Time } from '@angular/common';
import { FormControl } from '@angular/forms';
import { MatTableDataSource } from '@angular/material';

@Component({
  selector: 'app-admin-main',
  templateUrl: './admin-main.component.html',
  styleUrls: ['./admin-main.component.sass']
})
export class AdminMainComponent implements OnInit {
  
  @ViewChild('fromLocationSelect') child: MatSelect;
  @ViewChild('toLocationSelect') child2: MatSelect;
  @ViewChild('selecctedFlightId') child3: MatSelect;
  amount = 10
  fromLocationToView: Object[]
  toLocationToView: Object[]

  tripDatas: TripData[] = []
  flightDatas: FlightData[] = []
  blogDatas: BlogData[] = []
  flightDataSource = new MatTableDataSource<FlightData>(this.flightDatas)
  tripDataSource = new MatTableDataSource<TripData>(this.tripDatas)
  blogDataSource = new MatTableDataSource<BlogData>(this.blogDatas)
  displayedColumns: String[] = ['index', 'from', 'to', 'airline']

  locations: Object[]
  selectedAirlines: number
  airlineList: Object[]
  price: number
  tax: number
  service: number

  fromValue: number
  toValue: number
  airlineFilter: AirlineFilter[] = [
    {
      checked: false,
      airline: 'Batik Air'
    },
    {
      checked: false,
      airline: 'Lion Air'
    },
    {
      checked: false,
      airline: 'Garuda Indonesia'
    },
    {
      checked: false,
      airline: 'Citilink'
    }
  ]
  trainFilter: TrainFilter[] = [
    {
      checked: false,
      train: 'Api Tanpa Kereta'
    }
  ]

  flightPageNow: number = 1
  flightPageCount: number[] = []
  flightPage: number = 5

  tripPageNow: number = 1
  tripPageCount: number[] = []
  tripPage: number = 5

  blogPageNow: number = 1
  blogPageCount: number[] = []
  blogPage: number = 5

  trips: Object[]
  selectedTripId: number
  tripFromValue: number
  tripToValue: number
  tripBerangkatValue: Date
  tripSampaiValue: Date
  selectedTrain: number
  trainList: Object[]
  tripPrice: number
  tripTax: number
  tripService: number
  tripLocationToView: Object[]
  tripToUpdateValue: number
  tripFromUpdateValue: number
  tripBerangkatUpdateValue: Date
  tripSampaiUpdateValue: Date
  tripArrivalUpdateTimeValue: Time = {hours: 0, minutes: 0}
  tripDepartureUpdateTimeValue: Time = {hours: 0, minutes: 0}
  selectedTrainUpdate: number
  tripUpdatePrice: number
  tripUpdateTax: number
  tripUpdateService: number
  tripDuration: number
  selectedDeleteTripId: number

  blogs: Object[] = []
  selectedBlogId: number
  updateBlogTitle: string
  updateBlogCategory: string
  updateBlogDetail: string
  selectedRemoveBlogId: number

  fromUpdateValue: number
  toUpdateValue: number
  berangkatValue: Date
  sampaiValue: Date
  berangkatUpdateValue: Date
  sampaiUpdateValue: Date
  flights: Object[]
  selectedFlightId: number
  departureUpdateTimeValue: Time = {hours: 0, minutes: 0}
  arrivalUpdateTimeValue: Time = {hours: 0, minutes: 0}
  flightToUpdate: Object

  selectedUpdateAirlines: number
  updatePrice: number
  updateTax: number
  updateService: number

  selectedDeleteFlightId: number

  blogTitle: string
  blogCategory: string
  blogDetails: string

  blogInterval: any

  constructor(
    private graphqlService: GraphqlServiceService,
    private changedetectorRef: ChangeDetectorRef
  ) { }

  format(cmd){
    if(cmd === 'createlink') {
      let url = prompt("Enter the link here: ", "http:\/\/")
      document.execCommand(cmd, false, url)
    } else {
      document.execCommand(cmd, false, null)
    }
  }

  removeBlog() {
    this.graphqlService.removeBlog(this.selectedRemoveBlogId).subscribe(async query => {})
    this.blogInterval = setInterval(() => {
      this.graphqlService.getAllBlog().subscribe(async query => {
        this.blogs = query.data.getAllBlog
        await this.loadBlogTableData()
      })
    }, 1000)
  }

  updateBlog() {
    this.updateBlogDetail = document.getElementById("updateBlogDetail").innerHTML
    this.graphqlService.getBlogById(this.selectedBlogId).subscribe(async query => {
      var blogToUpdate = query.data.getBlogById
      console.log(blogToUpdate)
      await this.graphqlService.updateBlog(this.selectedBlogId, blogToUpdate.AuthorName, this.updateBlogCategory, "../../../../assets/blog-image-2.jpg", this.updateBlogTitle, this.updateBlogDetail, blogToUpdate.View).subscribe(async query => {})
    })
    this.blogInterval = setInterval(() => {
      this.graphqlService.getAllBlog().subscribe(async query => {
        this.blogs = query.data.getAllBlog
        await this.loadBlogTableData()
      })
    }, 1000)
  }

  saveBlog() {
    this.blogDetails = document.getElementById("blogDetail").innerHTML
    this.graphqlService.createBlog("admin", this.blogCategory, "../../../../assets/blog-image-1.jpg", this.blogTitle, this.blogDetails, 0).subscribe(async query => {
      console.log(query.data.createBlog.Id)
    })
  }

  init() {
    this.child._openedStream.subscribe(() => {
      this.child.panel.nativeElement.addEventListener('scroll', event => {
        if(event.target.scrollTop+256 >= event.target.scrollHeight) {
          this.amount += 5
          this.addItems(this.amount)
        }
      })
    })
    this.child2._openedStream.subscribe(() => {
      this.child2.panel.nativeElement.addEventListener('scroll', event => {
        if(event.target.scrollTop+256 >= event.target.scrollHeight) {
          this.amount += 5
          this.addItems(this.amount)
        }
      })
    })
    this.addItems(this.amount)
  }

  public setValue() {
    this.subscribet4 = this.graphqlService.getFlightById(this.selectedFlightId).subscribe(async query => {
      this.flightToUpdate = query.data.getFlightById
    })
  }

  public insertFlight() {
    var berangkat = this.berangkatValue.getFullYear() + "-" + this.berangkatValue.getMonth()+1 + "-" + this.berangkatValue.getDate() + "T00:00:00Z"
    var sampai = this.sampaiValue.getFullYear() + "-" + this.sampaiValue.getMonth()+1 + "-" + this.sampaiValue.getDate() + "T00:00:00Z"
    this.graphqlService.createFlight(this.fromValue, this.toValue, berangkat, sampai, this.selectedAirlines, this.price, this.tax, this.service, 0).subscribe(async query => {
      console.log(query.data.createFlight.ID)
    })
  }

  public updateFlight() {
    var berangkat = this.berangkatUpdateValue.getFullYear() + "-" + this.berangkatUpdateValue.getMonth()+1 + "-" + this.berangkatUpdateValue.getDate() + "T00:00:00Z"
    var sampai = this.sampaiUpdateValue.getFullYear() + "-" + this.sampaiUpdateValue.getMonth()+1 + "-" + this.sampaiUpdateValue.getDate() + "T00:00:00Z"
    this.graphqlService.updateFlight(this.selectedFlightId, this.fromUpdateValue, this.toUpdateValue, berangkat, sampai, this.selectedUpdateAirlines, this.updatePrice, this.updateTax, this.updateService, 0).subscribe(async query => {})
  }

  public removeFlight() {
    this.graphqlService.removeFlight(this.selectedDeleteFlightId).subscribe(async query => {})
    this.subscriber3 = this.graphqlService.getAllFlight().subscribe(async query => {
      this.flights = query.data.getAllFlight
    })
  }

  public insertTrip() {
    var berangkat = this.tripBerangkatValue.getFullYear() + "-" + this.tripBerangkatValue.getMonth()+1 + "-" + this.tripBerangkatValue.getDate() + "T00:00:00Z"
    var sampai = this.tripSampaiValue.getFullYear() + "-" + this.tripSampaiValue.getMonth()+1 + "-" + this.tripSampaiValue.getDate() + "T00:00:00Z"
    this.tripDuration = 1
    this.graphqlService.createTrip(this.tripToValue, sampai, this.tripDuration, this.tripTax, this.tripService, this.selectedTrain, this.tripFromValue, berangkat, this.tripPrice).subscribe(async query => {
      console.log(query.data.createTrip.Id)
    })
  }

  public updateTrip() {
    var berangkat = this.tripBerangkatUpdateValue.getFullYear() + "-" + this.tripBerangkatUpdateValue.getMonth()+1 + "-" + this.tripBerangkatUpdateValue.getDate() + "T00:00:00Z"
    var sampai = this.tripSampaiUpdateValue.getFullYear() + "-" + this.tripSampaiUpdateValue.getMonth()+1 + "-" + this.tripSampaiUpdateValue.getDate() + "T00:00:00Z"
    this.tripDuration = 1
    this.graphqlService.updateTrip(this.selectedTripId, this.tripToUpdateValue, sampai, this.tripDuration, this.tripUpdateTax, this.tripUpdateService, this.selectedTrainUpdate, this.tripFromUpdateValue, berangkat, this.tripUpdatePrice).subscribe(async query => {})
  }

  public removeTrip() {
    this.graphqlService.removeTrip(this.selectedDeleteTripId).subscribe(async query => {})
    // this.subscriber3 = this.graphqlService.getAllFlight().subscribe(async query => {
    //   this.flights = query.data.getAllFlight
    // })
  }

  public resetTrainFilter() {
    this.tripDatas = []
    for(let idx = 0; idx < this.trips.length; idx++) {
      this.tripDatas.push({index: this.trips[idx].Id, from: this.trips[idx].From.Name, to: this.trips[idx].To.Name, train: this.trips[idx].Train.Name})
    }
    console.table(this.tripDatas)
    this.tripDataSource = new MatTableDataSource<TripData>(this.tripDatas)
  }

  public resetFilter() {
    for (let index = 0; index < this.airlineFilter.length; index++) {
      this.airlineFilter[index].checked = false
    }
    for(let index = 0; index < this.trainFilter.length; index++) {
      this.trainFilter[index].checked = false
    }
    this.flightDatas = []
    for(let idx = 0; idx < this.flights.length; idx++) {
      this.flightDatas.push({index: this.flights[idx].ID, from: this.flights[idx].From.City, to: this.flights[idx].To.City, airline: this.flights[idx].Airline.Name})
    }
    this.flightDataSource = new MatTableDataSource<FlightData>(this.flightDatas)
  }

  public doFilterAirline() {

    this.flightDatas = []
    for(let idx = 0; idx < this.airlineFilter.length; idx++) {
      if(this.airlineFilter[idx].checked == true) {
        for(let idx2 = 0; idx2 < this.flights.length; idx2++) {
          if(this.flights[idx2].Airline.Name == this.airlineFilter[idx].airline) {
            this.flightDatas.push({index: this.flights[idx2].ID, from: this.flights[idx2].From.City, to: this.flights[idx2].To.City, airline: this.flights[idx2].Airline.Name})
          }
        }
      }
    }
    this.flightDataSource = new MatTableDataSource<FlightData>(this.flightDatas)
  }

  public doFilterTrain() {
    this.tripDatas = []
    for(let idx = 0; idx < this.trainFilter.length; idx++) {
      if(this.trainFilter[idx].checked == true) {
        for(let idx2 = 0; idx2 < this.trips.length; idx2++) {
          if(this.trips[idx2].Train.Name == this.trainFilter[idx].train) {
            this.tripDatas.push({index: this.trips[idx2].Id, from: this.trips[idx2].From.Name, to: this.trips[idx2].To.Name, train: this.trips[idx2].Train.Name})
          }
        }
      }
    }
    this.tripDataSource = new MatTableDataSource<TripData>(this.tripDatas)    
  }

  subscriber: Subscription
  subscriber2: Subscription
  subscriber3: Subscription
  subscribet4: Subscription
  subscriber5: Subscription
  subscriber6: Subscription
  subscriber7: Subscription
  subscriber8: Subscription

  displayedFlightColumn: string[] = ['index', 'from', 'to', 'airline']
  displayedTripColumns: string[] = ['index', 'from', 'to', 'train']
  displayedBlogColumns: string[] = ['index', 'Blog Title', 'Author Name', 'Upload Date', 'View']

  ngOnInit() {
    this.fromLocationToView = []
    this.toLocationToView = []
    this.tripLocationToView = []
    this.graphqlService.getAirline().subscribe(async query => {
      this.airlineList = query.data.getAirline
      await this.graphqlService.getAllFlight().subscribe(async query => {
        this.flights = query.data.getAllFlight
        for (let i = 1; i <= Math.round(this.flights.length/this.flightPage); i++) {
          this.flightPageCount.push(i)
        }
        if(this.flights.length % this.flightPage != 0 && this.flights.length > this.flightPage && Math.round(this.flights.length/this.flightPage) * this.flightPage<this.flights.length){
          this.flightPageCount.push(Math.round(this.flights.length/this.flightPage)+1)
        }
        this.loadFlightTableData()
        await this.graphqlService.getAllStation().subscribe(async query => {
          this.tripLocationToView = query.data.getAllStation
          await this.graphqlService.getAllTrip().subscribe(async query => {
            this.trips = query.data.getAllTrip
            for (let i = 1; i <= Math.round(this.trips.length/this.tripPage); i++) {
              this.tripPageCount.push(i)
            }
            if(this.trips.length % this.flightPage != 0 && this.trips.length > this.tripPage && Math.round(this.trips.length/this.tripPage) * this.tripPage<this.trips.length){
              this.tripPageCount.push(Math.round(this.trips.length/this.tripPage)+1)
            }
            this.loadTripTableData()
            await this.graphqlService.getAllTrip().subscribe(async query => {
              this.trips = query.data.getAllTrip
              await this.graphqlService.getAllTrain().subscribe(async query => {
                this.trainList = query.data.getAllTrain
                await this.graphqlService.getAllBlog().subscribe(async query => {
                  this.blogs = query.data.getAllBlog
                  for (let i = 1; i <= Math.round(this.blogs.length/this.blogPage); i++) {
                    this.blogPageCount.push(i)
                  }
                  if(this.blogs.length % this.blogPage != 0 && this.blogs.length > this.blogPage && Math.round(this.blogs.length/this.blogPage) * this.blogPage<this.blogs.length){
                    this.blogPageCount.push(Math.round(this.blogs.length/this.blogPage)+1)
                  }
                  await this.loadBlogTableData()
                })
              })
            })
          })
        })
      })
    })
  }

  initialize() {
    this.subscriber = this.graphqlService.getAllAirport().subscribe(async query => {
      this.locations = query.data.getAllAirport
      await this.init()
    })
  }

  loadBlog() {
    this.graphqlService.getAllBlog().subscribe(async query => {
      this.blogs = query.data.getAllBlog
      await this.loadBlogTableData()
    })
  }

  loadBlogTableData() {
    this.blogDatas = []
    for(let i = 0; i < this.blogs.length; i++) {
      if(this.blogPagination(i)) {
        this.blogDatas.push({
          index: this.blogs[i].Id, authorName: this.blogs[i].AuthorName, blogTitle: this.blogs[i].BlogTitle, uploadDate: this.blogs[i].UploadDate, view: this.blogs[i].View
        })
      }
    }
    this.blogDataSource = new MatTableDataSource<BlogData>(this.blogDatas)
    this.initialize()
  }

  loadTrain() {
    this.subscriber7 = this.graphqlService.getAllTrain().subscribe(async query => {
      this.trainList = query.data.getAllTrain
    })
  }

  loadTrip() {
    this.subscriber6 = this.graphqlService.getAllTrip().subscribe(async query => {
      this.trips = query.data.getAllTrip
      await this.loadTripTableData()
    })
  }

  loadTripTableData() {
    this.tripDatas = []
    for(let idx = 0; idx < this.trips.length; idx++) {
      if(this.tripPagination(idx)) {
        this.tripDatas.push(
          {index: this.trips[idx].Id, from: this.trips[idx].From.Name, to: this.trips[idx].To.Name, train: this.trips[idx].Train.Name}
        )
      }
    }
    this.tripDataSource = new MatTableDataSource<TripData>(this.tripDatas)
    console.table(this.tripDatas)
  }

  loadAirline() {
    this.subscriber3 = this.graphqlService.getAllFlight().subscribe(async query => {
      this.flights = query.data.getAllFlight
      await this.loadFlightTableData()
    })
  }

  loadFlightTableData() {
    this.flightDatas = []
    for(let idx = 0; idx < this.flights.length; idx++) {
      if(this.flightPagination(idx)) {
        this.flightDatas.push(
          {index: idx, from: this.flights[idx].From.City, to: this.flights[idx].To.City, airline: this.flights[idx].Airline.Name}
        )
      }
    }
    this.flightDataSource = new MatTableDataSource<FlightData>(this.flightDatas)
    
  }

  addItems(count: number) {
    let index = this.fromLocationToView.length
    if(this.locations.length >= count) {
      for (let idx = index; idx < count; idx++) {
        this.fromLocationToView.push(this.locations[idx])
        this.toLocationToView.push(this.locations[idx])
      }
    } else {
      for(let idx = index; idx < this.locations.length; idx++) {
        this.fromLocationToView.push(this.locations[idx])
        this.toLocationToView.push(this.locations[idx])        
      }
    }
  }

  flightPagination(idx: number) {
    return (idx < this.flightPage * this.flightPageNow) && (idx >= this.flightPage * (this.flightPageNow - 1))
  }

  flightChangePageNow(idx: number) {
    this.flightPageNow = idx
    this.loadFlightTableData()
  }

  tripPagination(idx: number) {
    return (idx < this.tripPage * this.tripPageNow) && (idx >= this.tripPage * (this.tripPageNow - 1))
  }

  tripChangePageNow(idx: number) {
    this.tripPageNow = idx
    this.loadTripTableData()
  }

  blogPagination(idx: number) {
    return (idx < this.blogPage * this.blogPageNow) && (idx >= this.blogPage * (this.blogPageNow - 1))
  }

  blogChangePageNow(idx: number) {
    this.blogPageNow = idx
    this.loadBlogTableData()
  }

}

export interface FlightData {
  index: number,
  from: string,
  to: string,
  airline: string
}

export interface TripData {
  index: number,
  from: string,
  to: string,
  train: string
}

export interface BlogData {
  index: number,
  blogTitle: string,
  authorName: string,
  uploadDate: string,
  view: number
}

export class AirlineFilter {
  checked: boolean
  airline: string
}

export class TrainFilter {
  checked: boolean
  train: string
}