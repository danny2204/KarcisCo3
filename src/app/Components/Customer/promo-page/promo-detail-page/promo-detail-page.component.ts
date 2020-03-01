import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

@Component({
  selector: 'app-promo-detail-page',
  templateUrl: './promo-detail-page.component.html',
  styleUrls: ['./promo-detail-page.component.sass']
})
export class PromoDetailPageComponent implements OnInit {

  promo: Object = {}
  otherPromo: Object[] = []
  currentUrl: string

  constructor(
    private route: ActivatedRoute,
    private graphqlService: GraphqlServiceService,
    private router: Router
  ) { }

  copyToClipboard() {
    const selBox = document.createElement('textarea')
    selBox.style.opacity = '0'
    selBox.value = this.currentUrl
    document.body.appendChild(selBox)
    selBox.focus()
    selBox.select()
    document.execCommand("copy")
    document.body.removeChild(selBox)
  }

  backToList() {
    this.router.navigate(['/promo-page'])
  }

  loadOtherPromo() {
    this.graphqlService.getOtherPromo(+this.route.snapshot.paramMap.get("id")).subscribe(async query => {
      this.otherPromo = query.data.getOtherPromo
    })
  }

  ngOnInit() {
    this.currentUrl = window.location.href
    this.graphqlService.getPromoById(+this.route.snapshot.paramMap.get("id")).subscribe(async query => {
      this.promo = query.data.getPromoById
      await this.loadOtherPromo()
    })
  }

}
