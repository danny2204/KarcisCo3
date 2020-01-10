import { Injectable } from '@angular/core';
import { Observable, Subscription } from 'rxjs';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { Observer } from 'apollo-client/util/Observable';

@Injectable({
  providedIn: 'root'
})
export class GraphqlServiceService {

  constructor(
    private apollo : Apollo
  ) { }

  createUser(frontName: string, backName: string, email: string, password: string, phoneNumber: string) : Observable<any> {
    return this.apollo.mutate<any>({
      mutation: gql`
        mutation createNewUser ($frontName: String!, $backName: String!, $email: String!, $password: String!, $phoneNumber: String!) {
          createUser(frontName: $frontName, backName: $backName, email: $email, password: $password, phoneNumber: $phoneNumber) {
            frontName,
            backName,
            email,
            password,
            phoneNumber
          }
        }
      `,
      variables: {
        "frontName": frontName,
        "backName": backName,
        "email": email,
        "password": password,
        "phoneNumber": phoneNumber,
      }
    })
  }

  getAllFlight() : Observable<any>{
    return this.apollo.query<any>({
      query: gql`
        query {
          getAllFlight {
            id,
            airline {
              id,
              name,
            }
          }
        }
      `
    })
  }

  getUserByEmailAndPhone(arg: string) : Observable<Query> {
    return this.apollo.query<Query>({
      query: gql`
        query GetUserByEmailAndPhone($arg: String!){
          getUserByEmailAndPhone(arg:$arg){
            id,
            frontName,
            backName,
            email,
            password,
          }
        }
      `,
      variables : {
        "arg" : arg,
      }
    })

  }

}

export class Query {
  data: any
}