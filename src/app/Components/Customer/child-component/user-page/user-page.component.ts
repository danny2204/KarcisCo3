import { Component, OnInit, ElementRef } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

declare const gapi: any;
declare var FB: any;

@Component({
  selector: 'app-user-page',
  templateUrl: './user-page.component.html',
  styleUrls: ['./user-page.component.sass']
})
export class UserPageComponent implements OnInit {

  User: Object = {}
  firstNameValue: string
  secondNameValue: string
  titelValue: string
  kotaValue: string
  addressValue: string
  postCodeValue: string
  emailValue: string
  phoneNumberValue: string
  locations: Object[] = []

  firstNameError: boolean = false
  secondNameError: boolean = false
  titelError: boolean = false
  kotaError: boolean = false
  addressError: boolean = false
  codePostError: boolean = false

  editEmail:boolean = false
  editPhone: boolean = false

  constructor(
    private graphqlService: GraphqlServiceService,
    private element: ElementRef
  ) { }

  private clientId = '196774946684-8vspumhpt7hbhqie6s34vb22tun649ej.apps.googleusercontent.com';

  private scope = [
    'profile',
    'email',
    'https://www.googleapis.com/auth/plus.me',
    'https://www.googleapis.com/auth/contacts.readonly',
    'https://www.googleapis.com/auth/admin.directory.user.readonly'
  ].join(' ');

  public auth2: any;

  openEdit() {
    this.editPhone = true
    this.editEmail = true
  }

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
        this.User[0].GoogleId = profile.getId()
        this.updateUser()
      }, (error) => {
        console.log(JSON.stringify(error, undefined, 2));
      });

  }

  clearFacebookId() {
    this.User[0].FacebookId = ''
    this.updateUser()
  }

  clearGoogleId() {
    this.User[0].GoogleId = ''
    this.updateUser()
  }

  facebookLogin() {
    FB.login((response) => {
      if(response.authResponse) {
        let userID = response.authResponse.userID
        this.User[0].FacebookId = userID
        console.log(this.User[0].FacebookId)
        this.updateUser()
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

  checkTitelError() {
    if(this.titelValue == undefined || this.titelValue == "") {
      this.titelError = true
    }
  }

  checkFirstNameError() {
    if(this.firstNameValue == "" || this.firstNameValue == undefined) this.firstNameError = true
    else this.firstNameError = false
  }

  checkSecondNameError() {
    if(this.secondNameValue == "" || this.secondNameValue == undefined) this.secondNameError = true
    else this.secondNameError = false
  }

  checkPostCodeError() {
    if(this.postCodeValue == "" || this.postCodeValue == undefined) this.codePostError = true
    else this.codePostError = false
  }

  checkAddressError() {
    if(this.addressValue == "" || this.addressValue == undefined) this.addressError = true
    else this.addressError = false
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

    this.graphqlService.getUserByEmailAndPhone(JSON.parse(sessionStorage.getItem('user'))[0].Email).subscribe(async query => {
      this.User = query.data.getUserByEmailAndPhone
      this.firstNameValue = this.User[0].FrontName
      this.secondNameValue = this.User[0].BackName
      this.titelValue = this.User[0].Title
      this.kotaValue = this.User[0].City
      this.postCodeValue = this.User[0].PostCode
      this.addressValue = this.User[0].Address
      this.emailValue = this.User[0].Email
      this.phoneNumberValue = this.User[0].PhoneNumber
    })
  }

  updateEmailPhone() {

    var fetchPromise = fetch("https://api.veriphone.io/v2/verify?phone="+this.phoneNumberValue+"&key=C11A8FD445A34296A726CE181697C6FB");
    fetchPromise.then(response => {
      return response.json();
    }).then(q => {
      console.log(q)
      if(q.phone_valid == false) alert("invalid phone number")
      else {
        this.graphqlService.updateUser(this.User[0].password, this.User[0].Title, this.User[0].Address, this.User[0].GoogleId, this.User[0].Gender, this.User[0].FacebookId, this.User[0].ID, this.User[0].FrontName, this.User[0].BackName, this.emailValue, this.phoneNumberValue, this.User[0].City, this.User[0].PostCode).subscribe(async query => {})
        this.editEmail = false
        this.editPhone = false
        alert("update success")
      }
    });

  }

  updateUser() {

    if(this.firstNameError == false && this.secondNameError == false && this.titelError == false && this.kotaError == false && this.addressError == false && this.codePostError == false) {
      this.graphqlService.updateUser(this.User[0].password, this.titelValue, this.addressValue, this.User[0].GoogleId, this.User[0].Gender, this.User[0].FacebookId, this.User[0].id, this.firstNameValue, this.secondNameValue, this.User[0].email, this.User[0].phoneNumber, this.kotaValue, this.postCodeValue).subscribe(async query => {})
      alert("update success")
    }

  }

}
