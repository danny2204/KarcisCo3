import { Component, AfterViewInit, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { Router } from '@angular/router'
import { JsonPipe } from '@angular/common';
import { MatDialog, MatDialogConfig } from '@angular/material';
import { HistorySearchComponent } from '../../../../history-search/history-search.component';

@Component({
  selector: 'app-quick-card-pesawat',
  templateUrl: './quick-card-pesawat.component.html',
  styleUrls: ['./quick-card-pesawat.component.sass']
})
export class QuickCardPesawatComponent implements OnInit {

  flight: Object[]
  supkrep : Subscription
  content: Object
  fromValue: string
  toValue: string
  supkrep2: Subscription
  historyList: Object[] = []
  fromDateValue: Date
  toDateValue: Date
  config: MatDialogConfig = new MatDialogConfig();

  constructor(
    private service: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private emitterService: EventEmitterService,
    private router: Router,
    private dialog: MatDialog
  ) { }

  ngOnInit() {
    this.config.autoFocus = false;
    this.config.restoreFocus = true;
    this.supkrep = this.service.getAllAirport().subscribe(async query => {
      this.flight = query.data.getAllAirport
      this.sharedService.airports = this.flight
    })
  }

  public openHistory() {
    this.dialog.open(HistorySearchComponent, this.config)
  }

  public getFlight() {
    if(localStorage.getItem("history") != null) {
      this.historyList = JSON.parse(localStorage.getItem("history"))
    } else {
      localStorage.setItem("history", "")
      this.historyList = []
    }
    this.historyList.push({
      label: "Flight",
      date: this.fromDateValue,
      from: this.fromValue,
      to: this.toValue
    })
    localStorage.setItem("history", JSON.stringify(this.historyList))

    if(this.fromValue.localeCompare(this.toValue) == 0) {
      alert("Destination and From Location cannot be the same")
    } else {
      this.sharedService.flightSearchData = {to: "", from: ""}
      this.sharedService.flightSearchData.to = this.toValue
      this.sharedService.flightSearchData.from = this.fromValue
      this.supkrep2 = this.service.getFlightByToAndFrom(this.toValue, this.fromValue).subscribe(async query => {
        this.sharedService.flightSearchResult = query.data.getFlightByToAndFrom
        console.table(this.sharedService.flightSearchResult)
        await this.router.navigate(["/pesawat"])
      })
    }
  }

}
