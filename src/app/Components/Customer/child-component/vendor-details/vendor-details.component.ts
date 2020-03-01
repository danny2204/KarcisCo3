import { Component, OnInit } from '@angular/core';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-vendor-details',
  templateUrl: './vendor-details.component.html',
  styleUrls: ['./vendor-details.component.sass']
})
export class VendorDetailsComponent implements OnInit {

  carData: Object
  vendorList: Object[]
  carSearchResult: Object
  displayedVendorData: Object[] = []
  priceList: number[] = []
  showData: number = 3

  constructor(
    private sharedService: SharedServiceService,
    private router: Router
  ) { }

  public toCheckout(idx: number) {
    localStorage.setItem("whichBought", "car")
    this.sharedService.vendorChoosed = this.vendorList[idx].Vendor
    this.sharedService.carBought.Price = this.vendorList[idx].Price
    this.router.navigate(['/checkout'])
  }

  public sortByPriceAscending() {
    var length = this.vendorList.length
    this.displayedVendorData = []
    for (let i = 0; i < length-1; i++) {
      for (let j = 0; j < length-i-1; j++) {
        if(this.priceList[j] < this.priceList[j+1]) {

          var tempPrice = this.priceList[j]
          this.priceList[j] = this.priceList[j+1]
          this.priceList[j+1] = tempPrice

          var temp = this.vendorList[j]
          this.vendorList[j] = this.vendorList[j+1]
          this.vendorList[j+1] = temp
        }
      }
    }
    this.loadData()
  }

  public loadData() {
    if(this.vendorList.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedVendorData.push(this.vendorList[index])
      }
    }
    else {
      for (let index = 0; index < this.vendorList.length; index++) {
        this.displayedVendorData.push(this.vendorList[index])
      }
    }
  }

  ngOnInit() {
    this.vendorList = this.sharedService.vendorsData
    this.carData = this.sharedService.carData
    this.loadData()
  }

  scroll = (event): void => {
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 3
      if(this.vendorList.length >= this.showData) {
        for (let index = this.showData-3; index < this.showData; index++) {
          this.displayedVendorData.push(this.vendorList[index])
        }
      } else {
        for (let index = this.showData-3; index < this.vendorList.length; index++) {
          this.displayedVendorData.push(this.vendorList[index])
        }
      }
    }
  }
}
