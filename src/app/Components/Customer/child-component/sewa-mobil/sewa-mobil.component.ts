import { Component, OnInit } from '@angular/core';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sewa-mobil',
  templateUrl: './sewa-mobil.component.html',
  styleUrls: ['./sewa-mobil.component.sass']
})
export class SewaMobilComponent implements OnInit {

  passengerFilter: PassengerFilter[] = [
    {
      checked: false,
      label: "Semua",
      value: 0
    },
    {
      checked: false,
      label: "< 5 Penumpang",
      value: 4
    },
    {
      checked: false,
      label: "5 - 6 Penumpang",
      value: 6
    },
    {
      checked: false,
      label: "> 6 Penumpang",
      value: 7
    }
  ]
  priceList: number[] = []
  brandFilter: BrandFilter[] = []
  modelFilter: ModelFilter[] = []
  sortBy: string = "rekomendasi"
  carData: Object[]
  displayedCarData: Object[]
  test: Object
  tempCarData: Object[]
  carSearchData: CarSearchData
  showData = 5
  config: MatDialogConfig = new MatDialogConfig();

  constructor(
    private eventEmitterService: EventEmitterService,
    private sharedService: SharedServiceService,
    private dialog: MatDialog,
    private Router: Router
  ) { }

  ngOnInit() {
    this.displayedCarData = []
    this.carData = []
    this.config.autoFocus = false
    this.config.restoreFocus = true
    this.carSearchData = {where: "", to: new Date, from: new Date}
    this.carSearchData.from = this.sharedService.carSearchData.from
    this.carSearchData.to = this.sharedService.carSearchData.to
    this.carData = this.sharedService.carSearchResult
    this.loadLowestPrice()
    this.sort()
    this.filterBrand()
    this.filterModel()
    window.addEventListener('scroll', this.scroll, true)
  }

  public showVendors(idx: number) {
    var vendorList = []
    this.sharedService.carData = this.carData[idx]
    for(let i = 0; i < this.carData[idx].CarVendor.length; i++) {
      vendorList.push(this.carData[idx].CarVendor[i])
    }
    this.sharedService.vendorsData = vendorList
    this.sharedService.carBought = this.carData[idx]
    this.Router.navigate(['/vendor-detail'])
  }

  public loadLowestPrice() {
    var length = this.carData.length
    for(let i = 0; i < length; i++) {
      var vendorLength = this.carData[i].CarVendor.length
      var minimalPrice = 9999999
      for(let j = 0; j < vendorLength; j++) {
        if (this.carData[i].CarVendor[j].Price < minimalPrice) {
          minimalPrice = this.carData[i].CarVendor[j].Price
        }
      }
      this.priceList[i] = minimalPrice
    }
  }

  public sort() {
    if(this.sortBy.localeCompare("rekomendasi") == 0 || this.sortBy.localeCompare("harga-terendah")) {
      this.sortByPriceAscending()
    }
    if(this.sortBy.localeCompare("harga-tertinggi") == 0) {
      this.sortByPriceDescending()
    }
  }

  public sortByPriceAscending() {
    var length = this.carData.length
    this.displayedCarData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] < this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.carData[j]
          this.carData[j] = this.carData[j+1]
          this.carData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public sortByPriceDescending() {
    var length = this.carData.length
    this.displayedCarData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] > this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.carData[j]
          this.carData[j] = this.carData[j+1]
          this.carData[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public filterBrand() {
    var brandNameList: string[] = [];
    this.carData.forEach(element => {
      brandNameList.push(element.Brand)
    })
    var filteredCar = brandNameList.filter(
      (brandName, i, arr) => arr.findIndex( name => name === brandName ) === i
    )
    filteredCar.forEach(element => {
      this.brandFilter.push({
        label: element,
        checked: false
      })
    })
  }

  public filterModel() {
    var modelNameList: string[] = [];
    this.carData.forEach(element => {
      modelNameList.push(element.Model)
    })
    var filteredModel = modelNameList.filter(
      (modelName, i, arr) => arr.findIndex( name => name === modelName ) === i
    )
    filteredModel.forEach(element => {
      this.modelFilter.push({
        label: element,
        checked: false
      })
    })
  }

  public changeSearch() {
    this.dialog.open(PesawatWidgetComponent, this.config)
  }

  public checkValidation(idx: number) {
    if(!this.checkBrandFilter(idx)) {
      return false
    } else if(!this.checkModelFilter(idx)) {
      return false
    } else if(!this.checkPassengerFilter(idx)) {
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
  }

  public checkBrandFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.brandFilter.length; index++) {
      if(this.brandFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.brandFilter.length; index++) {
        if(this.brandFilter[index].checked == true) {
          if(this.carData[idx].Brand == this.brandFilter[index].label) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkModelFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.modelFilter.length; index++) {
      if(this.modelFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.modelFilter.length; index++) {
        if(this.modelFilter[index].checked == true) {
          if(this.carData[idx].Model == this.modelFilter[index].label) return true
          else continue
        }
      }
      return false
    } else {
      return true
    }
  }

  public checkPassengerFilter(idx: number): boolean {
    var isChecked = false
    for (let index = 0; index < this.passengerFilter.length; index++) {
      if(this.passengerFilter[index].checked == true) {
        isChecked = true
        break
      }
    }
    if(isChecked == true) {
      for (let index = 0; index < this.passengerFilter.length; index++) {
        if(this.passengerFilter[index].checked == true) {
          if(this.passengerFilter[index].value == 0) {
            return true
          } else if(this.passengerFilter[index].value == 6) {
            if(this.carData[idx].Passenger == 5 || this.carData[idx].Passenger == 6) {
              return true
            }
          } else if(this.passengerFilter[index].value == 4) {
            if(this.carData[idx].Passenger <= 4) {
              return true
            }
          } else if(this.passengerFilter[index].value == 7) {
            if(this.carData[idx].Passenger >= 7) {
              return true
            }
          }
        }
      }
      return false
    } else {
      return true
    }
  }

  public loadData() {
    if(this.carData.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedCarData.push(this.carData[index])
      }
    }
    else {
      for (let index = 0; index < this.carData.length; index++) {
        this.displayedCarData.push(this.carData[index])
      }
    }
  }

  scroll = (event): void => {
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 5
      if(this.carData.length >= this.showData) {
        for (let index = this.showData-5; index < this.showData; index++) {
          this.displayedCarData.push(this.carData[index])
        }
      } else {
        for (let index = this.showData-5; index < this.carData.length; index++) {
          this.displayedCarData.push(this.carData[index])
        }
      }
    }
  }


}

export class PassengerFilter {
  label: string
  value: number
  checked: boolean
}

export class BrandFilter {
  label: string
  checked: boolean
}

export class ModelFilter {
  label: string
  checked: boolean
}

export class CarSearchData {
  where: string
  from: Date
  to: Date
}