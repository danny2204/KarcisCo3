import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

@Component({
  selector: 'app-main-home',
  templateUrl: './main-home.component.html',
  styleUrls: ['./main-home.component.sass']
})
export class MainHomeComponent implements OnInit {

  toEmail: string = ""

  constructor(
    private graphqlService: GraphqlServiceService
  ) { }

  ngOnInit() {
  }

  sendMail() {
    this.graphqlService.sendMail(this.toEmail).subscribe(async query => {})
  }

}
