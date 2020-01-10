import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-login-dialog',
  templateUrl: './login-dialog.component.html',
  styleUrls: ['./login-dialog.component.sass']
})
export class LoginDialogComponent implements OnInit {

  showPass = false
  text = "LANJUTKAN"

  constructor(
    private GraphQLService: GraphqlServiceService
  ) { }

  ngOnInit() {
  }

  user$: Subscription
  user: User[]

  doLogin(emailInput : string) {
    this.user = []
    this.user$ = this.GraphQLService.getUserByEmailAndPhone(emailInput).subscribe(async query => {
      this.user = query.data.getUserByEmailAndPhone;
      await this.doCheck()
    })
  }

  showPassField() {
    if(this.showPass == false) {
      return "none"
    } else if(this.showPass == true) {
      return "block"
    }
  }

  checkPass(password: string){
    if(this.user[0].password == password) {
      console.log("success login")
      sessionStorage.setItem("user", JSON.stringify(this.user))
    } else {
      console.log("invalild password")
    }
  }

  doCheck() {
    if (this.user == null) {
      console.log("no user")
    } else {
      console.table(this.user)
      this.showPass = true
      this.text = "LOGIN"
    }
  }

}

export class User {
  id: number
  frontName: string
  backName: string
  email: string
  password: string
  phoneNumber: string
}
