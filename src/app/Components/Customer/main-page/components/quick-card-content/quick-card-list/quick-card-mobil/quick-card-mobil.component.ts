import { Component, OnInit, ViewChild } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { MatSelect } from '@angular/material';

@Component({
  selector: 'app-quick-card-mobil',
  templateUrl: './quick-card-mobil.component.html',
  styleUrls: ['./quick-card-mobil.component.sass']
})
export class QuickCardMobilComponent implements OnInit {

  @ViewChild('fromLocationSelect') child: MatSelect;
  car: Object[]
  supkrep : Subscription
  fromValue: string
  supkrep2: Subscription
  amount = 10
  rentLocationToView: Object[] = ["Jakarta Pusat"]
  locations: Object[]

  toDate: Date
  fromDate: Date
 
  constructor(
    private service: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private router: Router
  ) { }

  init() {
    this.child._openedStream.subscribe(() => {
      this.child.panel.nativeElement.addEventListener('scroll', event => {
        if(event.target.scrollTop+256 >= event.target.scrollHeight) {
          this.amount += 5
          // this.addItems(this.amount)
        }
      })
    })
    // this.addItems(this.amount)
  }

  ngOnInit() {
    this.rentLocationToView = ["Jakarta Pusat"]
    // this.supkrep = this.service.getAllLocation().subscribe(async query => {
    //   this.locations = query.data.getAllLocation
    //   await this.init()
    // })
  }

  addItems(count: number) {
    let index = this.rentLocationToView.length
    for (let idx = index; idx < count; idx++) {
      this.rentLocationToView.push(this.locations[idx])
    }
  }

  public doValidation() {
    if(this.toDate == undefined || this.fromDate == undefined) {
      alert("From Date and To Date must be filled")
      return false
    } else if(this.toDate == this.fromDate) {
      alert("From Date and To Date cannot be the same")
      return false
    }
  }

  public getCar() {
    if(this.doValidation() == true) {

    } else {
      this.sharedService.carSearchData = {to: new Date(), from: new Date(), where: ""}
      this.sharedService.carSearchData.to = this.toDate
      this.sharedService.carSearchData.from = this.fromDate
      this.sharedService.carSearchData.where = this.fromValue
      var stringForSearch = this.fromDate.getFullYear() + "-" + (this.fromDate.getMonth()+1) + "-" + this.fromDate.getDate()
      this.supkrep2 = this.service.getCarByLocationAndToAndFrom(this.fromValue, stringForSearch).subscribe(async query => {
        this.sharedService.carSearchResult = query.data.getCarByLocationAndToAndFrom
        console.log(this.sharedService.carSearchResult)
        await this.router.navigate(["/sewa-mobil"])
      })
    }
  }

}
