import { Component, OnInit } from '@angular/core';
import * as L from 'leaflet';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hotel-map',
  templateUrl: './hotel-map.component.html',
  styleUrls: ['./hotel-map.component.sass']
})
export class HotelMapComponent implements OnInit {

  markerList: L.marker[] = []
  map: L.map
  hotelList: Object[]
  selectedHotel: Object = {}
  displayedFasilitasFilter: FasilitasFilter[] = []
  displayedTipeFilter: TipePropertiFilter[] = []

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

  markersLayer: any

  loadMoreFasilitas: boolean = false
  loadMoreTipe: boolean = false

  showMoreFasilitas() {
    this.loadMoreFasilitas = true
    for(let i = 5 ; i < this.fasilitasFilter.length; i++) {
      this.displayedFasilitasFilter.push(this.fasilitasFilter[i])
    }
  }

  showMoreTipe() {
    this.loadMoreTipe = true
    for(let i = 5; i < this.tipeFilter.length; i++) {
      this.displayedTipeFilter.push(this.tipeFilter[i])
    }
  }

  constructor(
    private sharedService: SharedServiceService,
    private router: Router
  ) { }

  ngOnInit() {
    this.map = L.map('map').setView([51.505, -0.09], 13);
    this.markersLayer = L.featureGroup().addTo(this.map)
    this.hotelList = this.sharedService.hotelSearchResult

    for(let i = 0 ; i < 5; i++) {
      this.displayedFasilitasFilter.push(this.fasilitasFilter[i])
      this.displayedTipeFilter.push(this.tipeFilter[i])
    }

    if(localStorage.getItem("openMapFrom") == "hotel-page") {
      this.markerList = []
      for(let idx = 0; idx < this.hotelList.length; idx++) {
        var firstPointer = L.marker([this.hotelList[idx].Latitude, this.hotelList[idx].Longitude])
        firstPointer.Name = this.hotelList[idx].Name
        firstPointer.addTo(this.markersLayer).bindPopup(this.hotelList[idx].Name).openPopup();
        this.markerList.push(firstPointer)
      }
    } else if(localStorage.getItem("openMapFrom") == "hotel-detail") {
      this.markerList = []
      var selectedHotel = this.sharedService.selectedHotel
      var pointer = L.marker([selectedHotel.Longitude, selectedHotel.Latitude])
      this.markerList.push(pointer)
    }

    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(this.map);

    this.markersLayer.on("click", this.onclick)
  }

  public changeView(idx: number) {
    this.map.setView([this.hotelList[idx].Latitude, this.hotelList[idx].Longitude], 13)
    L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
    }).addTo(this.map);
    document.getElementById("right-panel").style.display = "flex"
    this.selectedHotel = this.hotelList[idx]
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
          if(this.hotelList[idx].Rating == this.starsFilter[index].value) return true
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
          for(let index2 = 0; index2 < this.hotelList[idx].HotelFacilities.length; index2++) {
            if(this.hotelList[idx].HotelFacilities[index2].FacilityName == this.fasilitasFilter[index].label) return true
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
          if(this.hotelList[idx].Type == this.tipeFilter[index].label) return true
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
          if(this.hotelList[idx].TiketComReview.AverageScore >= this.starsFilter[index].value) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
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

    // if(localStorage.getItem("openMapFrom") == "hotel-page") {
    //   this.markerList = []
    //   var firstPointer = L.marker([this.hotelList[idx].Latitude, this.hotelList[idx].Longitude])
    //   firstPointer.Name = this.hotelList[idx].Name
    //   firstPointer.addTo(this.markersLayer).bindPopup(this.hotelList[idx].Name).openPopup();
    // } else if(localStorage.getItem("openMapFrom") == "hotel-detail") {
    //   this.markerList = []
    //   var selectedHotel = this.sharedService.selectedHotel
    //   var pointer = L.marker([selectedHotel.Longitude, selectedHotel.Latitude])
    //   this.markerList.push(this.sharedService.selectedHotel)
    // }

    return true
  }

  public toDetail() {
    this.sharedService.selectedHotel = this.selectedHotel
    this.router.navigate(['/hotel-detail'])

    
  }

  public onclick(event) {
    console.log(event.layer.test)
  }

  public backToList() {
    this.router.navigate(['/hotel'])
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

  public getLowestPrice() {
    if(this.selectedHotel != undefined){
      var minimalPrice = 9999999
      for(let idx = 0; idx < this.selectedHotel.Room.length; idx++) {
        if(this.selectedHotel.Room[idx].Price < minimalPrice) {
          minimalPrice = this.selectedHotel.Room[idx].Price
        }
      }
      return minimalPrice
    }
    return 0
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