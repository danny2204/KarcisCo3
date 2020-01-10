import { Component, OnInit } from '@angular/core';
import { MatDialog, MatDialogConfig } from '@angular/material';
import { LoginDialogComponent } from '../login-dialog/login-dialog.component';
import { RegisterDialogComponent } from '../register-dialog/register-dialog.component';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.sass']
})
export class HeaderComponent implements OnInit {

  config: MatDialogConfig = new MatDialogConfig();

  constructor(
    public dialog: MatDialog
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

}
