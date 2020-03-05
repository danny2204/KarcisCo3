import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WebSocketServiceService } from 'src/app/Service/web-socket-service.service';
import { stringify } from 'querystring';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

@Component({
  selector: 'app-chat-detail',
  templateUrl: './chat-detail.component.html',
  styleUrls: ['./chat-detail.component.sass']
})
export class ChatDetailComponent implements OnInit {

  constructor(
    private router: Router,
    private webSocket: WebSocketServiceService,
    private graphqlService: GraphqlServiceService
  ) { }

  contentString: string

  addWebSocket() {
    this.webSocket.listen("chat").subscribe(async data => {
      this.contentString = data.toString()
      this.chatList.push({
        sender: JSON.parse(sessionStorage.getItem("user"))[0].ID,
        content: this.contentString
      })
      console.table(this.chatList)
      await this.updateChatContent()
    })

  }
  newChatId: number
  updateChatContent() {
    if(sessionStorage.getItem("chatFrom") == "addNew") {
      this.graphqlService.createChat(parseInt(JSON.parse(sessionStorage.getItem("user"))[0].ID), parseInt(localStorage.getItem("destinationUser")), JSON.stringify(this.chatList)).subscribe(async query => {
        this.newChatId = query.data.createChat.Id
        console.log("success insert")
        sessionStorage.setItem("chatFrom", "selectChat")
        localStorage.setItem("choosedChat", "")
      })
    } else if(sessionStorage.getItem("chatFrom") == "selectChat"){
      if(localStorage.getItem("choosedChat") != "") {
        console.log("asdasd")
        var chatId = JSON.parse(localStorage.getItem("choosedChat")).Id
        let user1 = JSON.parse(localStorage.getItem("choosedChat")).User1
        let user2 = JSON.parse(localStorage.getItem("choosedChat")).User2
        this.graphqlService.updateChat(chatId, JSON.stringify(this.chatList), user1, user2).subscribe(async query => {
          console.log("success update")
        })
      } else {
        let user1 = parseInt(JSON.parse(sessionStorage.getItem("user"))[0].ID)
        let user2 = parseInt(localStorage.getItem("destinationUser"))
        console.log(this.newChatId)
        console.log(user1)
        console.log(user2)
        this.graphqlService.updateChat(this.newChatId, JSON.stringify(this.chatList), user1, user2).subscribe(async query => {
          console.log("success update")
        })
      }
    }
  }

  chatList: Chat[] = []
  sender: number

  ngOnInit() { 
    this.addWebSocket()
    if(localStorage.getItem("choosedChat") == null) {
      this.chatList = []
    } else {
      var temp = JSON.parse(localStorage.getItem("choosedChat"))
      var tempContent = JSON.parse(temp.Content)
      this.chatList = []
      for(let i = 0; i < tempContent.length; i++) {
        this.chatList.push({
          sender: tempContent[i].sender,
          content: tempContent[i].content
        })
      }
    }

  }

  checkIsImage(idx: number) {
    var content = this.chatList[idx].content
    if((content.startsWith("img:") && content.endsWith(".jpg")) || ((content.startsWith("img:") && content.endsWith(".png")))) {
      return content.split(":")[1]
    }
    return ""
  }

  checkIsUser(idx: number) {
    if(JSON.parse(sessionStorage.getItem("user"))[0].ID != this.chatList[idx].sender) {
      return false
    }
    return true
  }

  backToChatPage() {
    this.router.navigate(['chat-page'])
  }
  chatContent: string

  sendMsg() {
    this.webSocket.emit("chat", this.chatContent)
  }

}

export interface Chat {
  sender: string
  content: string
}
