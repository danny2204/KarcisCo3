import { Component, OnInit } from '@angular/core';
declare var FB: any;

@Component({
  selector: 'app-facebook-button',
  templateUrl: './facebook-button.component.html',
  styleUrls: ['./facebook-button.component.sass']
})
export class FacebookButtonComponent implements OnInit {

  constructor() { }

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

  facebookLogin() {
    FB.login((response) => {
      if(response.authResponse) {
        let userID = response.authResponse.userID
        //tarik data user
        FB.api(
          userID, (res) => {
            console.log(res);
          }
        )
      } else {
        console.log("login failed");
      }
    })
  }

}
