import { Component, OnInit } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
declare var FB: any;

@Component({
  selector: 'app-facebook-button',
  templateUrl: './facebook-button.component.html',
  styleUrls: ['./facebook-button.component.sass']
})
export class FacebookButtonComponent implements OnInit {

  constructor(
    private graphqlService: GraphqlServiceService
  ) { }

  ngOnInit() {
    FB.init({
      appId: '1582559861881888',
      cookie: true,
      xfbml: true,
      version: 'v3.1'
    });
    FB.AppEvents.logPageView();

    (function(d, s, id){
      var js, fjs = d.getElementsByTagName(s)[0];
      if(d.getElementById(id)) {
        return;
      }
      js = d.createElement(s);
      js.id = id;
      js.src = 'https://connect.facebook.net/en_US/sdk.js';
      fjs.parentNode.insertBefore(js, fjs);
    }(document, 'script', 'facebook-jssdk'));
  }
  user: Object[] = []
  facebookLogin() {
    FB.login((response) => {
      if(response.authResponse) {
        let userID = response.authResponse.userID
        //tarik data user
        this.graphqlService.getUserByAuthId(userID).subscribe(async query => {
          this.user = query.data.getUserByAuthId
          if(this.user.length == 0) {
            alert("no such account")
          } else {
            alert("success")
            sessionStorage.setItem("user", JSON.stringify(this.user))
          }
        })
        FB.api(
          userID, (res) => {

          }
        )
      } else {
        console.log("login failed");
      }
    })
  }

}
