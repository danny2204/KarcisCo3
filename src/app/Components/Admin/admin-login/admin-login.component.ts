import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

@Component({
  selector: 'app-admin-login',
  templateUrl: './admin-login.component.html',
  styleUrls: ['./admin-login.component.sass']
})
export class AdminLoginComponent implements OnInit {

  constructor(
    private router: Router,
    private graphqlService: GraphqlServiceService
  ) { }

  adminEmail: string = ""
  adminPassword: string = ""
  user: Object[] = []

  ngOnInit() {
  }

  doLogin() {
    this.graphqlService.getAdminByEmail(this.adminEmail).subscribe(async query => {
      this.user = query.data.getAdminByEmail
      console.log(this.adminPassword)
      console.log(this.user[0].password)
      if(this.user[0].password == this.adminPassword) {
        sessionStorage.setItem("admin", JSON.stringify(this.user[0]))
        this.router.navigate(['/admin-main'])
      } else {
        alert("there is no such admin !")
      }
    })
  }

}
