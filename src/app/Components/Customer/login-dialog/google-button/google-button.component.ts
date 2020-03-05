import { Component, AfterViewInit, ElementRef } from '@angular/core';
import { GraphqlServiceService } from 'src/app/Service/graphql-service.service';

declare const gapi: any;

@Component({
  selector: 'app-google-button',
  templateUrl: './google-button.component.html',
  styleUrls: ['./google-button.component.sass']
})

export class GoogleButtonComponent implements AfterViewInit {

  private clientId = '196774946684-8vspumhpt7hbhqie6s34vb22tun649ej.apps.googleusercontent.com';

  private scope = [
    'profile',
    'email',
    'https://www.googleapis.com/auth/plus.me',
    'https://www.googleapis.com/auth/contacts.readonly',
    'https://www.googleapis.com/auth/admin.directory.user.readonly'
  ].join(' ');

  public auth2: any;

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

  user: Object[] = []

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
        this.graphqlService.getUserByAuthId(profile.getId().toString()).subscribe(async query => {
          this.user = query.data.getUserByAuthId
          console.table(this.user)
          if(this.user.length == 0) alert("no user")
          else {
            alert("success")
            sessionStorage.setItem("user", JSON.stringify(this.user))
          }
        })
      }, (error) => {
        console.log(JSON.stringify(error, undefined, 2));
      });

  } 
  constructor(
    private element: ElementRef,
    private graphqlService: GraphqlServiceService
  ) {
    console.log('ElementRef: ', this.element);
  }

  ngAfterViewInit() {
    this.googleInit();
  }
}