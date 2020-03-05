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
  user: Object[] = []

  doLogin(emailInput : string) {
    this.user = []
    console.log((emailInput))
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
    if(this.user[0].Password == password) {
      console.log("success login")
      sessionStorage.setItem("user", JSON.stringify(this.user))
      window.location.reload()
    } else {
      alert("invalid password !")
    }
  }

  toogleVisible() {
    var x = document.getElementById("passwordInput")
    if (x.getAttribute("type") === "password") {
      x.setAttribute("type", "text")
    } else {
      x.setAttribute("type", "password")
    }
  }

  doCheck() {
    if (this.user.length == 0) {
      alert("there is no such user")
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
