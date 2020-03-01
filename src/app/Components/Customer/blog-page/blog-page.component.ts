import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

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
    private graphqlService: GraphqlServiceService
  ) { }

  addInfiniteScroll() {
    window.addEventListener('scroll', this.scroll, true)
    this.isInfiniteScroll = true
  }

  ngOnInit() {
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
