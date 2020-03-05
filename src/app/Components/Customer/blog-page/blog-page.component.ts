import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Router } from '@angular/router';
import { WebSocketServiceService } from 'src/app/Service/web-socket-service.service';

@Component({
  selector: 'app-blog-page',
  templateUrl: './blog-page.component.html',
  styleUrls: ['./blog-page.component.sass']
})
export class BlogPageComponent implements OnInit {

  blogList: Object[] = []
  showData: number = 6
  displayedBlog: Object[] = []
  isInfiniteScroll: boolean = false

  constructor(
    private graphqlService: GraphqlServiceService,
    private router: Router,
    private webSocket: WebSocketServiceService
  ) { }

  addInfiniteScroll() {
    window.addEventListener('scroll', this.scroll, true)
    this.isInfiniteScroll = true
  }

  toDetail(id: number) {
    this.router.navigate(['/blog-detail/', id])
  }

  ngOnInit() {
    this.addWebSocket()
    this.graphqlService.getAllBlog().subscribe(async query => {
      this.blogList = query.data.getAllBlog
      await this.loadData()
    })
  }

  public loadData() {
    if(this.blogList.length >= this.showData) {
      for (let index = 0; index < this.showData; index++) {
        this.displayedBlog.push(this.blogList[index])
      }
    }
    else {
      for (let index = 0; index < this.blogList.length; index++) {
        this.displayedBlog.push(this.blogList[index])
      }
    }
  }

  addWebSocket() {
    this.webSocket.listen("blog").subscribe(async data => {
      alert("new blog inserted ! Refresh the page !")
    })
  }

  scroll = (event): void => {
    // console.log("scroll Y : ", window.scrollY, " document height : ", document.documentElement.clientHeight)
    if(window.scrollY + window.innerHeight >= document.body.scrollHeight) {
      this.showData += 5
      if(this.blogList.length >= this.showData) {
        for (let index = this.showData-5; index < this.showData; index++) {
          this.displayedBlog.push(this.blogList[index])
        }
      } else {
        for (let index = this.showData-5; index < this.blogList.length; index++) {
          this.displayedBlog.push(this.blogList[index])
        }        
      }
    }
  }

}
