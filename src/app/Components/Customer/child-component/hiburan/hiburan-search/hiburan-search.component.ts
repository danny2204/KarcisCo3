import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { DatePipe } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hiburan-search',
  templateUrl: './hiburan-search.component.html',
  styleUrls: ['./hiburan-search.component.sass'],
  providers: [DatePipe]
})
export class HiburanSearchComponent implements OnInit {

  typeValue: string
  fromDate: Date
  toDate: Date
  showData: number = 6
  categoryList: CategoryFilter[] = [
    {
      checked: false,
      label: "Activities"
    },
    {
      checked: false,
      label: "Attractions"
    },
    {
      checked: false,
      label: "Events"
    }
  ]

  bottomPrice: number = 0
  topPrice: number = 100000000
  hiburanList: Object[] = []
  displayedHiburanList: Object[] = []
  sortValue: string

  priceList: number[] = []

  constructor(
    private graphqlService: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private datePipe: DatePipe,
    private router: Router
  ) { }

  checkPriceValidation(idx: number) {
    if(this.topPrice < this.bottomPrice) {
      return true
    }
    if(this.getEntertainmentTicketLowestPrice(idx) < this.topPrice && this.getEntertainmentTicketLowestPrice(idx) > this.bottomPrice) {
      return true
    }
    return false
  }

  checkDateValidation(idx: number) {
    if(this.fromDate != undefined && this.toDate != undefined) {
      if(this.datePipe.transform(new Date(this.hiburanList[idx].EventStartDate), "yyyy-MM-dd") >= this.datePipe.transform(new Date(this.fromDate), "yyyy-MM-dd") 
      && this.datePipe.transform(new Date(this.hiburanList[idx].EventEndDate), "yyyy-MM-dd") <= this.datePipe.transform(new Date(this.toDate), "yyyy-MM-dd")) {
        return true
      }
      return false
    }
    return true
  }

  checkCategory(idx: number) {
    let isChecked = false
    for (let index = 0; index < this.categoryList.length; index++) {
      if(this.categoryList[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.categoryList.length; index++) {
        if(this.categoryList[index].checked == true) {
          if(this.hiburanList[idx].Type == this.categoryList[index].label) {
            return true
          }
          // for(let index2 = 0; index2 < this.hiburanList[idx].length; index2++) {
          //   if(this.hotelData[idx].HotelFacilities[index2].FacilityName == this.fasilitasFilter[index].label) return true
          //   else continue
          // }
        }
      }
      return false
    } else {
      return true
    }
  }

  checkValidation(idx: number) {
    if(!this.checkPriceValidation(idx)) return false
    if(!this.checkDateValidation(idx)) return false
    if(!this.checkCategory(idx)) return false
    return true
  }

  sortNewest() {
    this.displayedHiburanList = []
    this.hiburanList.sort((a, b)=>{
      if(this.datePipe.transform(new Date(a.UploadDate), "yyyy-MM-dd") < this.datePipe.transform(b.UpdloadDate), "yyyy-MM-dd") {
        return -1
      } else if (this.datePipe.transform(new Date(a.UploadDate), "yyyy-MM-dd") > this.datePipe.transform(b.UpdloadDate), "yyyy-MM-dd"){
        return 1
      }
      return 0
    })
  }

  public sortByPriceDescending() {
    var length = this.hiburanList.length
    this.displayedHiburanList = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] < this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.hiburanList[j]
          this.hiburanList[j] = this.hiburanList[j+1]
          this.hiburanList[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public sortByPriceAscending() {
    var length = this.hiburanList.length
    this.displayedHiburanList = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] > this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.hiburanList[j]
          this.hiburanList[j] = this.hiburanList[j+1]
          this.hiburanList[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  sort() {
    if(this.sortValue == "newest") {
      this.sortNewest()
      this.loadData()
    }
    if(this.sortValue == "highToLow") {
      this.sortByPriceDescending()
    }
    if(this.sortValue == "lowToHigh" || this.sortValue == "recommend") {
      this.sortByPriceAscending()
    }
  }

  resetFilter() {
    this.fromDate = undefined
    this.toDate = undefined
    this.bottomPrice = 0
    this.topPrice = 10000000
    for(let i = 0; i < this.categoryList.length; i++) {
      this.categoryList[i].checked = false
    }
  }

  ngOnInit() {
    this.hiburanList = []
    this.displayedHiburanList = []
    this.priceList = []
    this.typeValue = this.sharedService.entertainmentTypeValue
    this.sortValue = "newest"
    if(this.typeValue == "") {
      this.graphqlService.getAllEntertainment().subscribe(async query => {
        this.hiburanList = query.data.getAllEntertainment
        for(let i = 0; i < this.hiburanList.length; i++) {
          this.priceList[i] = this.getEntertainmentTicketLowestPrice(i)
        }
        await this.loadData()
      })
    } else {
      this.graphqlService.getEntertainmentByType(this.typeValue).subscribe(async query => {
        this.hiburanList = query.data.getEntertainmentByType
        await this.loadData()
      })
    }
    this.sort()
    window.addEventListener("scroll", this.scroll, true)
  }

  public loadData() {
    // console.log(this.hiburanList.length)
    if(this.hiburanList.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedHiburanList.push(this.hiburanList[index])
      }
      // console.table(this.displayedHiburanList)
    }
    else {
      for (let index = 0; index < this.hiburanList.length; index++) {
        this.displayedHiburanList.push(this.hiburanList[index])
      }
    }
    // this.sortNewest()
    // console.table(this.hiburanList)
  }

  scroll = (event): void => {
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 6
      if(this.hiburanList.length >= this.showData) {
        for (let index = this.showData-6; index < this.showData; index++) {
          this.displayedHiburanList.push(this.hiburanList[index])
        }
      } else {
        for (let index = this.showData-6; index < this.hiburanList.length; index++) {
          this.displayedHiburanList.push(this.hiburanList[index])
        }
      }
    }
  }

  getEntertainmentTicketLowestPrice(idx: number) {
    var minimumPrice = 9999999
    for(let i = 0 ; i < this.hiburanList[idx].EntertainmentTicket.length; i++) {
      if(this.hiburanList[idx].EntertainmentTicket[i].Price < minimumPrice) {
        minimumPrice = this.hiburanList[idx].EntertainmentTicket[i].Price
      }
    }
    return minimumPrice
  }

  goToDetail(idx: number) {
    localStorage.setItem("selectedEvent", this.hiburanList[idx].Id)
    this.router.navigate(['/hiburan-detail'])
  }

}

export class CategoryFilter {
  checked: boolean
  label: string
}
