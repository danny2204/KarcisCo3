import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-register-dialog',
  templateUrl: './register-dialog.component.html',
  styleUrls: ['./register-dialog.component.sass']
})
export class RegisterDialogComponent implements OnInit {

  constructor(
    private graphService : GraphqlServiceService
  ) { }

  ngOnInit() {
  }

  user$: Subscription
  user: User

  doRegister(email: string, frontName: string, backName: string, password: string, phoneNumber: string) {
    this.user = null
    this.user$ = this.graphService.createUser(frontName, backName, email, password, phoneNumber).subscribe(async query => {})
  }

}

export class User {
  frontName: string
  backName: string
  email: string
  password: string
  phoneNumber: string
}