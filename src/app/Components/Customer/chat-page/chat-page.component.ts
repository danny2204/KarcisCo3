import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-chat-page',
  templateUrl: './chat-page.component.html',
  styleUrls: ['./chat-page.component.sass']
})
export class ChatPageComponent implements OnInit {

  constructor(
    private graphqlService: GraphqlServiceService,
    private router: Router
  ) { }

  chatList: Object[] = []

  ngOnInit() {
    var userId = (JSON.parse(sessionStorage.getItem("user"))[0].ID)
    this.graphqlService.getAllChat(userId).subscribe(async query => {
      this.chatList = query.data.getAllChat
    })
  }

  backToHome() {
    this.router.navigate(['/'])
  }

  addNewChat() {
    var userTo = prompt("enter user id")
    localStorage.setItem("destinationUser", userTo)
    sessionStorage.setItem("chatFrom", "addNew")
    this.router.navigate(['/chat-detail'])
  }

  toDetail(idx: number) {
    sessionStorage.setItem("chatFrom", "selectChat")
    localStorage.setItem("choosedChat", JSON.stringify(this.chatList[idx]))
    this.router.navigate(['/chat-detail'])
  }

}
