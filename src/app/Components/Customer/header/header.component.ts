import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material';
import { LoginDialogComponent } from '../login-dialog/login-dialog.component';
import { RegisterDialogComponent } from '../register-dialog/register-dialog.component';
import { Subscription } from 'rxjs';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { EventEmitterService } from 'src/app/Service/event-emitter.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.sass']
})
export class HeaderComponent implements OnInit {

  config: MatDialogConfig = new MatDialogConfig();
  supkrep: Subscription
  flight: Object[]

  constructor(
    public dialog: MatDialog,
    private service: GraphqlServiceService,
    private sharedService: SharedServiceService,
    private emitterService: EventEmitterService,
    private router: Router
  ) { }

  ngOnInit() {
    this.config.autoFocus = false;
    this.config.restoreFocus = true;
  }

  openLoginDialog() {
    this.dialog.open(LoginDialogComponent, this.config);
  }

  openRegisterDialog() {
    this.dialog.open(RegisterDialogComponent, this.config);
  }

  public getFlight() {
    this.supkrep = this.service.getAllFlight().subscribe(async query => {
      this.sharedService.flightSearchResult = query.data.getAllFlight
      await this.router.navigate(["/pesawat"])
    })
  }
}
