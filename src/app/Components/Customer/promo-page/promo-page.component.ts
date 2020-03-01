import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-promo-page',
  templateUrl: './promo-page.component.html',
  styleUrls: ['./promo-page.component.sass']
})
export class PromoPageComponent implements OnInit {

  constructor(
    private graphqlService: GraphqlServiceService,
    private router: Router
  ) { }

  subscriber1: Subscription
  promoList: Object[]

  ngOnInit() {
    this.subscriber1 = this.graphqlService.getAllPromo().subscribe(async query => {
      this.promoList = query.data.getAllPromo
    })
  }

  goToDetail(i: Object) {
    this.router.navigate(['promo-detail', i.Id])
  }

}
