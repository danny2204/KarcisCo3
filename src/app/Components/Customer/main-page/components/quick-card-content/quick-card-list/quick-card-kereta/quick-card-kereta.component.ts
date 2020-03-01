import { Component, OnInit, ViewChild } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Router } from '@angular/router';
import { Subscription } from 'rxjs';
import { MatSelect } from '@angular/material';

@Component({
  selector: 'app-quick-card-kereta',
  templateUrl: './quick-card-kereta.component.html',
  styleUrls: ['./quick-card-kereta.component.sass']
})
export class QuickCardKeretaComponent implements OnInit {

  @ViewChild('fromLocationSelect') child: MatSelect;
  @ViewChild('toLocationSelect') child2: MatSelect;
  train: Object[]
  supkrep : Subscription
  content: Object
  fromValue: string
  toValue: string
  supkrep2: Subscription
  amount = 10
  fromLocationToView: Object[]
  toLocationToView: Object[]
  locations: Object[]
 
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

  ngOnInit() {
    this.fromLocationToView = ['Surabaya', 'Jakarta Pusat']
    this.toLocationToView = ['Surabaya', 'Jakarta Pusat']
    // this.supkrep = this.service.getAllLocation().subscribe(async query => {
    //   this.locations = query.data.getAllLocation
    //   await this.init()
    // })
    // this.init()
  }

  addItems(count: number) {
    let index = this.fromLocationToView.length
    for (let idx = index; idx < count; idx++) {
      this.fromLocationToView.push(this.locations[idx])
      this.toLocationToView.push(this.locations[idx])
    }
  }

  public getTrain() {
    if(this.toValue.localeCompare(this.fromValue) == 0) {
      alert("From and Destination cannot be the same")
    } else {
      this.sharedService.trainSearchData = {to: "", from: ""}
      this.sharedService.trainSearchData.to = this.toValue
      this.sharedService.trainSearchData.from = this.fromValue
      this.supkrep2 = this.service.getTripByToAndFrom(this.toValue, this.fromValue).subscribe(async query => {
        this.sharedService.trainSearchResult = query.data.getTripByToAndFrom
        console.table(this.sharedService.trainSearchResult)
        await this.router.navigate(["/kereta-api"])
      })
    }
  }

}
