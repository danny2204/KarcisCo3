import { Component, OnInit, ElementRef } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';
import { Subscription } from 'rxjs';
import { CountryCode } from '../../../material/country_code'

declare var FB: any;
declare const gapi: any;

@Component({
  selector: 'app-register-dialog',
  templateUrl: './register-dialog.component.html',
  styleUrls: ['./register-dialog.component.sass']
})
export class RegisterDialogComponent implements OnInit {

  facebookId: string = ""
  googleId: string = ""
  countryCode: Object[] = CountryCode

  private clientId = '196774946684-8vspumhpt7hbhqie6s34vb22tun649ej.apps.googleusercontent.com';

  private scope = [
    'profile',
    'email',
    'https://www.googleapis.com/auth/plus.me',
    'https://www.googleapis.com/auth/contacts.readonly',
    'https://www.googleapis.com/auth/admin.directory.user.readonly'
  ].join(' ');

  public auth2: any;

  constructor(
    private graphService : GraphqlServiceService,
    private element: ElementRef
  ) { }

  public googleInit() {
    const that = this;
    gapi.load('auth2', () => {
      that.auth2 = gapi.auth2.init({
        client_id: that.clientId,
        cookiepolicy: 'single_host_origin',
        scope: that.scope
      });
      that.attachSignin(that.element.nativeElement.firstChild);
    });
  }
  public attachSignin(element) {
    const that = this;
    this.auth2.attachClickHandler(element, {},
      (googleUser) => {
        const profile = googleUser.getBasicProfile();
        console.log('Token || ' + googleUser.getAuthResponse().id_token);
        console.log('ID: ' + profile.getId());
        console.log('Name: ' + profile.getName());
        console.log('Image URL: ' + profile.getImageUrl());
        console.log('Email: ' + profile.getEmail());
        this.googleId = profile.getId()
      }, (error) => {
        console.log(JSON.stringify(error, undefined, 2));
      });

  }

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

  user$: Subscription
  user: User

  emailError: boolean = false
  firstNameError: boolean = false
  passwordError: boolean = false
  phoneNumberError: boolean = false
  dialCode: string = ""

  doRegister(email: string, frontName: string, backName: string, password: string, phoneNumber: string) {
    if(email == "") this.emailError = true
    if(frontName == "") this.firstNameError = true
    if(password == "") this.passwordError = true
    if(phoneNumber == "") this.phoneNumberError = true
    if(email != "" && phoneNumber != "" && password != "" && frontName != "") {
      phoneNumber = this.dialCode + "-" + phoneNumber
      this.user = null
      this.user$ = this.graphService.createUser("", this.googleId, backName, email, password, phoneNumber, "", "", this.facebookId, "", frontName, "").subscribe(async query => {})
    } 
  }

  facebookLogin() {
    FB.login((response) => {
      if(response.authResponse) {
        let userID = response.authResponse.userID
        this.facebookId = userID
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
export class User {
  frontName: string
  backName: string
  email: string
  password: string
  phoneNumber: string
}