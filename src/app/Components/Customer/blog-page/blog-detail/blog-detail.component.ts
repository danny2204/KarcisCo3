import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-blog-detail',
  templateUrl: './blog-detail.component.html',
  styleUrls: ['./blog-detail.component.sass']
})
export class BlogDetailComponent implements OnInit {

  blog: Object = {}
  trendingBlog: Object[] = []
  currentUrl: string

  constructor(
    private graphqlService: GraphqlServiceService,
    private route: ActivatedRoute,
    private router: Router
  ) { }

  changeDetail(idx: number) {
    this.router.navigate(['blog-detail', idx])
    this.ngOnInit()
  }

  loadOtherBlog() {
    this.graphqlService.getAllBlogExcept(+this.route.snapshot.paramMap.get("id")).subscribe(async query => {
      this.trendingBlog = query.data.getAllBlogExcept
    })
  }

  ngOnInit() {
    this.trendingBlog = []
    this.blog = {}
    this.currentUrl = window.location.href
    this.graphqlService.getBlogById(+this.route.snapshot.paramMap.get("id")).subscribe(async query => {
      this.blog = query.data.getBlogById
      await this.loadOtherBlog()
    })
  }

}
