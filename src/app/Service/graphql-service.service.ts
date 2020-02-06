import { Injectable } from '@angular/core';
import { Observable, Subscription } from 'rxjs';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';

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

  getAllLocation(): Observable<any> {
    return this.apollo.query<any>({
      query: gql`
        query {
          getAllLocation {
            City
            Country
            Id
            Latitude
            Longitude
            Province
            ZIndex
          }
        }
      `
    })
  }

  getAllFlight() : Observable<any>{
    return this.apollo.query<any>({
      query: gql`
        query {
          getAllFlight {
            to {
              id,
              code,
              name,
              city,
              city_code,
              province,
            },
            from {
              id,
              code,
              name,
              city,
              city_code,
              province,
            },
            airline {
              id,
              name,
            },
            price,
            tax,
            service_charge,
            departure,
            arrival,
            duration,
            id
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

  getHotelByCity(arg1: string): Observable<any> {
    return this.apollo.query<Query>({
      query: gql `
        query GetHotelByCity($arg1: String!) {
          getHotelByCity(city: $arg1){
            id
            name
            imagePath
            location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            rating
            HotelFacilities {
              FacilityName
              FacilityPrice
              HotelReferId
              Id
            }
            address
          }
        }
      `,
      variables : {
        "arg1" : arg1
      }
    })
  }

  getFlightByToAndFrom(arg1: string, arg2: string) : Observable<any> {
    return this.apollo.query <Query>({
      query: gql`
        query GetFlightByToAndFrom($arg1: String!, $arg2: String!) {
          getFlightByToAndFrom(to:$arg1, from:$arg2){
            to {
              id,
              code,
              name,
              city,
              city_code,
              province,
            },
            from {
              id,
              code,
              name,
              city,
              city_code,
              province,
            },
            airline {
              id,
              name,
            },
            routes {
              id,
              AirportFrom {
                id,
                code,
                name,
                city,
                city_code,
                province,
              },
              AirportTo {
                id,
                code,
                name,
                city,
                city_code,
                province,
              },
              duration
            },
            price,
            tax,
            service_charge,
            departure,
            arrival,
            duration,
            id
          }
        }
      `,
      variables: {
        "arg1": arg1,
        "arg2": arg2,
      }
    })
  }

  getAllAirport() : Observable<any> {
    return this.apollo.query<any>({
      query: gql`
        query {
          getAllAirport {
            id,
            code,
            name,
            city,
            city_code,
            province,
            country,
          }
        }
      `
    })
  }

}

export class Query {
  data: any
}