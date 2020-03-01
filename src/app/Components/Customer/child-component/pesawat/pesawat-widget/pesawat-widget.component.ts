import { Component, OnInit } from '@angular/core';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { Subscription } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Router } from '@angular/router';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { PesawatComponent } from '../pesawat.component';

@Component({
  selector: 'app-pesawat-widget',
  templateUrl: './pesawat-widget.component.html',
  styleUrls: ['./pesawat-widget.component.sass']
})
export class PesawatWidgetComponent implements OnInit {

  airports: Object[]
  fromValue: string
  toValue: string
  supkrep2: Subscription

  constructor(
    private sharedService: SharedServiceService,
    private service: GraphqlServiceService,
    private router: Router,
  ) { }

  ngOnInit() {
    this.airports = this.sharedService.airports
  }

  public getFlight() {
    console.log(this.fromValue)
    console.log(this.toValue)
    if(this.fromValue.localeCompare(this.toValue) == 0) {
      alert("Destination and From Location cannot be the same")
    } else {
      this.sharedService.flightSearchData = {to: "", from: ""}
      this.sharedService.flightSearchData.to = this.toValue
      this.sharedService.flightSearchData.from = this.fromValue
      this.supkrep2 = this.service.getFlightByToAndFrom(this.toValue, this.fromValue).subscribe(async query => {
        this.sharedService.flightSearchResult = query.data.getFlightByToAndFrom
        console.table(this.sharedService.flightSearchResult)
        // await 
      })
    }
  }
}
