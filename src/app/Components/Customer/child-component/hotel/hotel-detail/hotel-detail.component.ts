import { Component, OnInit } from '@angular/core';
import { SharedServiceService } from 'src/app/Service/shared-service.service';
import { MatDialogConfig, MatDialog } from '@angular/material';
import { PopUpImageComponent } from './pop-up-image/pop-up-image.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-hotel-detail',
  templateUrl: './hotel-detail.component.html',
  styleUrls: ['./hotel-detail.component.sass']
})
export class HotelDetailComponent implements OnInit {

  test: string = "url('../../../../../../assets/house.png')"
  hotel: Object
  config: MatDialogConfig = new MatDialogConfig();

  constructor(
    private sharedService: SharedServiceService,
    private dialog: MatDialog,
    private router: Router
  ) { }

  ngOnInit() {
    console.table(this.sharedService.selectedHotel)
    this.hotel = this.sharedService.selectedHotel
    this.config.autoFocus = false;
    this.config.restoreFocus = true;
  }

  public openDialog() {
    this.dialog.open(PopUpImageComponent, this.config);
  }

  public toCheckout(idx: number) {
    localStorage.setItem("whichBought", "hotel")
    this.sharedService.hotelBought = this.hotel
    this.sharedService.hotelBought.Price = this.hotel.Room[idx].Price
    this.sharedService.roomChoosed = this.hotel.Room[idx]
    this.router.navigate(['/checkout'])
  }

}
