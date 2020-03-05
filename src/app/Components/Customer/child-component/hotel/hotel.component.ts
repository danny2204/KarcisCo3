import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MatDialog, MatDialogConfig } from '@angular/material';
import { SharedServiceService } from 'src/app/Service/shared-service.service';

@Component({
  selector: 'app-hotel',
  templateUrl: './hotel.component.html',
  styleUrls: ['./hotel.component.sass']
})
export class HotelComponent implements OnInit {

  hotelData: Object[]
  displayedHotelData: Object[]
  showData: number = 5
  sortBy: string
  showAllTipeFilter = false
  showAllFasilitasFilter = false
  starsFilter: StarFilter[] = [
    {
      label: "1 star",
      checked: false,
      value: 1
    },
    {
      label: "2 star",
      checked: false,
      value: 2
    },
    {
      label: "3 star",
      checked: false,
      value: 3
    },
    {
      label: "4 star",
      checked: false,
      value: 4
    },
    {
      label: "5 star",
      checked: false,
      value: 5
    }
  ]
  fasilitasFilter: FasilitasFilter[] = [
    {
      label: "Resepsionis 24 jam",
      checked: false
    },
    {
      label: "AC",
      checked: false
    },
    {
      label: "Lift",
      checked: false
    },
    {
      label: "Tempat Parkir",
      checked: false
    },
    {
      label: "Restoran",
      checked: false
    },
    {
      label: "Spa",
      checked: false
    },
    {
      label: "Kolam Renang",
      checked: false
    },
    {
      label: "WIFI",
      checked: false
    }
  ]
  tipeFilter: TipePropertiFilter[] = [
    {
      label: "Apartment",
      checked: false
    },
    {
      label: "Guest house",
      checked: false
    },
    {
      label: "Hostel",
      checked: false
    },
    {
      label: "Vacation Home",
      checked: false
    },
    {
      label: "Hotel-unknown",
      checked: false
    },
    {
      label: "Villa",
      checked: false
    },
    {
      label: "Resort villa",
      checked: false
    },
    {
      label: "Condominium",
      checked: false
    },
    {
      label: "Hotel",
      checked: false
    },
    {
      label: "Hostal",
      checked: false
    },
    {
      label: "Apart-hotel",
      checked: false
    },
    {
      label: "Resort",
      checked: false
    },
    {
      label: "Bed & breakfast",
      checked: false
    },
    {
      label: "Hotel Bersejarah",
      checked: false
    },
    {
      label: "Capsule Hotel",
      checked: false
    }
  ]
  ratingFilter: RatingFilter[] = [
    {
      checked: false,
      label: "9+ (Fantastis)",
      value: 9
    },
    {
      checked: false,
      label: "8+ (Mengesankan)",
      value: 8
    },
    {
      checked: false,
      label: "7+ (Keren)",
      value: 7
    },
    {
      checked: false,
      label: "6+ (baik)",
      value: 6
    }
  ]
  config: MatDialogConfig = new MatDialogConfig();
  priceList: number[] = []

  displayedTipeFilter: TipePropertiFilter[] = []
  displayedFasilitasFilter: FasilitasFilter[] = []

  showAllTipe() {
    this.showAllTipeFilter = true
    for(let i = 5; i < this.tipeFilter.length; i++) {
      this.displayedTipeFilter.push(this.tipeFilter[i])
    }
  }

  showAllFasilitas() {
    this.showAllFasilitasFilter = true
    for (let i = 5; i < this.fasilitasFilter.length; i++) {
      this.displayedFasilitasFilter.push(this.fasilitasFilter[i])
    }
  }

  constructor(
    private sharedService: SharedServiceService,
    private dialog: MatDialog,
    private Router: Router
  ) { }

  public goToMap() {
    localStorage.setItem("openMapFrom", "hotel-page")
    this.Router.navigate(['/hotel-map'])
  }

  ngOnInit() {
    this.hotelData = [];
    this.displayedHotelData = [];
    this.config.autoFocus = false
    this.config.restoreFocus = true
    this.hotelData = this.sharedService.hotelSearchResult
    for(let i = 0 ; i < 5; i++){
      this.displayedTipeFilter.push(this.tipeFilter[i])
      this.displayedFasilitasFilter.push(this.fasilitasFilter[i])
    }
    this.loadLowestPrice()
    this.loadData()
    window.addEventListener('scroll', this.scroll, true)
  }

  public selectHotel(idx: number) {
    this.sharedService.selectedHotel = this.hotelData[idx]
    this.Router.navigate(['/hotel-detail'])
  }

  public checkValidation(idx: number) {
    if(!this.checkStarFilter(idx)) {
      return false
    } else if(!this.checkFacilityFilter(idx)) {
      return false
    } else if(!this.checkTipeFilter(idx)) {
      return false
    } else if(!this.checkRatingFilter(idx)) {
      return false
    }
    return true
  }

  public sort() {
    if(this.sortBy == "harga-murah") {
      this.sortByPriceAscending();
    }
    if(this.sortBy == "harga-mahal") {
      this.sortByPriceDescending();
    }
    if(this.sortBy == "rating") {
      this.sortByRating();
    }
    if(this.sortBy == "bintang") {
      this.sortByBintang()
    }
    if(this.sortBy == "rekomendasi") {

    }
  }

  public loadData() {
    if(this.hotelData.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedHotelData.push(this.hotelData[index])
      }
    }
    else {
      for (let index = 0; index < this.hotelData.length; index++) {
        this.displayedHotelData.push(this.hotelData[index])
      }
    }
  }

  public loadLowestPrice() {
    var length = this.hotelData.length
    for(let i = 0; i < length; i++) {
      var vendorLength = this.hotelData[i].Room.length
      var minimalPrice = 9999999
      for(let j = 0; j < vendorLength; j++) {
        if (this.hotelData[i].Room[j].Price < minimalPrice) {
          minimalPrice = this.hotelData[i].Room[j].Price
        }
      }
      this.priceList[i] = minimalPrice
    }
  }

  public checkStarFilter(idx: number) {
    var isChecked = false
    for (let index = 0; index < this.starsFilter.length; index++) {
      if(this.starsFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.starsFilter.length; index++) {
        if(this.starsFilter[index].checked == true) {
          if(this.hotelData[idx].Rating == this.starsFilter[index].value) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkFacilityFilter(idx: number) {
    var isChecked = false
    for (let index = 0; index < this.fasilitasFilter.length; index++) {
      if(this.fasilitasFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.fasilitasFilter.length; index++) {
        if(this.fasilitasFilter[index].checked == true) {
          for(let index2 = 0; index2 < this.hotelData[idx].HotelFacilities.length; index2++) {
            if(this.hotelData[idx].HotelFacilities[index2].FacilityName == this.fasilitasFilter[index].label) return true
            else continue
          }
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkTipeFilter(idx: number) {
    var isChecked = false
    for (let index = 0; index < this.tipeFilter.length; index++) {
      if(this.tipeFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.tipeFilter.length; index++) {
        if(this.tipeFilter[index].checked == true) {
          if(this.hotelData[idx].Type == this.tipeFilter[index].label) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkRatingFilter(idx: number) {
    var isChecked = false
    for (let index = 0; index < this.ratingFilter.length; index++) {
      if(this.ratingFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.ratingFilter.length; index++) {
        if(this.ratingFilter[index].checked == true) {
          if(this.hotelData[idx].TiketComReview.AverageScore >= this.starsFilter[index].value) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public resetFilter() {
    for (let index = 0; index < this.starsFilter.length; index++) {
      this.starsFilter[index].checked = false
    }
    for (let index = 0; index < this.fasilitasFilter.length; index++) {
      this.fasilitasFilter[index].checked = false
    }
    for (let index = 0; index < this.tipeFilter.length; index++) {
      this.tipeFilter[index].checked = false      
    }
    for (let index = 0; index < this.ratingFilter.length; index++) {
      this.ratingFilter[index].checked = false
    }    
  }

  public sortByPriceAscending() {
    var length = this.hotelData.length
    this.displayedHotelData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] < this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.hotelData[j]
          this.hotelData[j] = this.hotelData[j+1]
          this.hotelData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public sortByPriceDescending() {
    var length = this.hotelData.length
    this.displayedHotelData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] > this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.hotelData[j]
          this.hotelData[j] = this.hotelData[j+1]
          this.hotelData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public sortByRating() {
    var length = this.hotelData.length
    this.displayedHotelData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.hotelData[j].TiketComReview.AverageScore < this.hotelData[j+1].TiketComReview.AverageScore) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.hotelData[j]
          this.hotelData[j] = this.hotelData[j+1]
          this.hotelData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public sortByBintang() {
    var length = this.hotelData.length
    this.displayedHotelData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.hotelData[j].Rating < this.hotelData[j+1].Rating) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.hotelData[j]
          this.hotelData[j] = this.hotelData[j+1]
          this.hotelData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  scroll = (event): void => {
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 5
      if(this.hotelData.length >= this.showData) {
        for (let index = this.showData-5; index < this.showData; index++) {
          this.displayedHotelData.push(this.hotelData[index])
        }
      } else {
        for (let index = this.showData-5; index < this.hotelData.length; index++) {
          this.displayedHotelData.push(this.hotelData[index])
        }
      }
    }
  }

}

export class StarFilter {
  label: string
  value: number
  checked: boolean
}

export class FasilitasFilter {
  label: string
  checked: boolean
}

export class TipePropertiFilter {
  label: string
  checked: boolean
}

export class RatingFilter {
  label: string
  value: number
  checked: boolean
}