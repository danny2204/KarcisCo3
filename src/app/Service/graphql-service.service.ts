import { Injectable } from '@angular/core';
import { Observable, Subscription, generate } from 'rxjs';
import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';
import { Time } from '@angular/common';
import { stringify } from 'querystring';

@Injectable({
  providedIn: 'root'
})
export class GraphqlServiceService {
  
  constructor(
    private apollo : Apollo
    ) { }

  getHotelById(id : number): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getHotelById($id: Int!) {
          getHotelById(Id: $id) {
            Address
            Id
            Latitude
            Longitude
            Name
            Rating
            Type
            Location {
              City
            }
            Room {
              Price
            }
          }
        }
      `, variables: {
        "id": id
      }
    })
  }

  getAllHotel(): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query {
          getAllHotel {
            Address
            Id
            Latitude
            Longitude
            Name
            Rating
            Type
            Location {
              City
            }
            Room {
              Price
            }
          }
        }
      `
    })
  }

  updateUser(password: string, title: string, address: string, GoogleId: string, Gender: string, FacebookId: string, id: number, frontName: string, backName: string, email: string, phoneNumber: string, city: string, postCode: string, mainLanguage: string) {
    return this.apollo.mutate<any>({
      mutation: gql`
        mutation updateUser($password: String!, $title: String!, $address: String!, $googleId: String!, $gender: String!, $facebookId: String!, $id: Int!, $frontName: String!, $backName: String!, $email: String!, $phoneNumber: String!, $city: String!, $postCode: String!, $mainLanguage: String!) {
          updateUser(Password: $password, Title: $title, Address: $address, GoogleId: $googleId, Gender: $gender, FacebookId: $facebookId, ID: $id, FrontName: $frontName, BackName: $backName, Email: $email, PhoneNumber: $phoneNumber, City: $city, PostCode: $postCode, MainLanguage: $mainLanguage) {
            ID
          }
        }
      `,
      variables: {
        "password": password,
        "title": title,
        "address": address,
        "googleId": GoogleId,
        "gender": Gender,
        "facebookId": FacebookId,
        "id": id,
        "frontName": frontName,
        "backName": backName,
        "email": email,
        "phoneNumber": phoneNumber,
        "city": city,
        "postCode": postCode,
        "mainLanguage": mainLanguage
      }
    })
  }

  removeEvent(id: number) {
    return this.apollo.mutate<any>({
      mutation: gql`
        mutation removeEntertainment($id: Int!) {
          removeEntertainment(Id: $id) {
            Id
          }
        }
      `, variables: {
        "id": id
      }
    })
  }

  updateEntertainment(id: number, name: string, locationReferId: number, type: string, eventStartDate: string, eventEndDate: string, description: string): Observable<any> {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation updateEntertainment($id: Int!, $name: String!, $type: String!, $locationReferId: Int!, $eventStartDate: DateTime!, $eventEndDate: DateTime!, $description: String!) {
          updateEntertainment(Id: $id, Name: $name, Type: $type, LocationReferId: $locationReferId, EventStartDate: $eventStartDate, EventEndDate: $eventEndDate, IsTrending: false, Description: $description) {
            Id
          }
        }
      `, variables: {
        "id": id,
        "name": name,
        "type": type,
        "locationReferId": locationReferId,
        "eventStartDate": eventStartDate,
        "eventEndDate": eventEndDate,
        "description": description
      }
    })
  }

  createEntertainment(name: string, locationReferId: number, type: string, eventStartDate: string, eventEndDate: string, description: string):Observable<any> {
    return this.apollo.mutate<any>({
      mutation: gql`
        mutation createEntertainment($name: String!, $type: String!, $locationReferId: Int!, $eventStartDate: DateTime!, $eventEndDate: DateTime!, $description: String!) {
          createEntertainment(Name: $name, Type: $type, LocationReferId: $locationReferId, EventStartDate: $eventStartDate, EventEndDate: $eventEndDate, IsTrending: false, Description: $description) {
            Id
          }
        }
      `, variables: {
        "name": name,
        "type": type,
        "locationReferId": locationReferId,
        "eventStartDate": eventStartDate,
        "eventEndDate": eventEndDate,
        "description": description
      }
    })
  }

  getEntertainmentById(id: number): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getEntertainmentById($id: Int!){
          getEntertainmentById(Id: $id) {
            Id
            EntertainmentTicket {
              Id
              Price
              Type
              Description
            }
            EventEndDate
            EventStartDate
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Name
            Type
            UploadDate
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  getAllEntertainment(): Observable<any> {
    return this.apollo.query<Query>({
      query: gql`
        query {
          getAllEntertainment {
            Id
            EntertainmentTicket {
              Id
              Price
              Type
            }
            EventEndDate
            EventStartDate
            UploadDate
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Name
            Type
            Description
          }
        }
      `
    })
  }

  getEntertainmentByType(type: string): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getEntertainmentByType($type: String!) {
          getEntertainmentByType(Type: $type) {
            Id
            EntertainmentTicket {
              Id
              Price
              Type
            }
            EventEndDate
            EventStartDate
            UploadDate
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Name
            Type
          }
        }
      `,
      variables: {
        "type": type
      }
    })
  }

  removeBlog(id: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation removeBlog($id: Int!) {
          removeBlog(Id: $id) {
            Id
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  updateBlog(id: number, authorName: string, category: string, imagePath: string, blogTitle: string, blogDetail: string, view: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation updateBlog($id: Int!, $authorName: String!, $category: String!, $imagePath: String!, $blogTitle: String!, $blogDetail: String!, $view: Int!) {
          updateBlog(Id: $id, AuthorName: $authorName, Category: $category, ImagePath: $imagePath, BlogTitle: $blogTitle, BlogDetail: $blogDetail, View: $view) {
            Id
          }
        }
      `,
      variables: {
        "id": id,
        "authorName": authorName,
        "category": category,
        "imagePath": imagePath,
        "blogTitle": blogTitle,
        "blogDetail": blogDetail,
        "view": view
      }
    })
  }

  createBlog(authorName: string, category: string, imagePath: string, blogTitle: string, blogDetail: string, view: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
      mutation createBlog($authorName: String!, $category: String!, $imagePath: String!, $blogTitle: String!, $blogDetail: String!, $view: Int!) {
        createBlog(AuthorName: $authorName, Category: $category, ImagePath: $imagePath, BlogTitle: $blogTitle, BlogDetail: $blogDetail, View: $view) {
          Id
        }
      }
      `,
      variables: {
        "authorName": authorName,
        "category": category,
        "imagePath": imagePath,
        "blogTitle": blogTitle,
        "blogDetail": blogDetail,
        "view": view
      }
    })
  }

  getAllBlogExcept(id: number): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getAllBlogExcept($id: Int!) {
          getAllBlogExcept(Id: $id) {
            AuthorName
            BlogDetail
            BlogTitle
            Category
            Id
            ImagePath
            UploadDate
            View
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  getBlogById(id: number): Observable<any> {
    return this.apollo.query<Query>({
      query: gql`
        query getBlogById($id: Int!) {
          getBlogById(Id: $id) {
            AuthorName
            BlogDetail
            BlogTitle
            Category
            Id
            ImagePath
            UploadDate
            View
          }
        }
      `, 
      variables: {
        "id": id
      }
    })
  }

  getAllBlog(): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query {
          getAllBlog {
            AuthorName
            BlogDetail
            BlogTitle
            Category
            Id
            ImagePath
            UploadDate
            View
          }
        }
      `, fetchPolicy: 'no-cache'
    })
  }

  getOtherPromo(id: number): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getOtherPromo($id: Int!) {
          getOtherPromo(Id: $id) {
            Id
            Name
            Platform
            PromoDetail
            PromoDetailList {
              DiscountAmount
              Id
              KodePromo
            }
            PromoEnd
            PromoStart
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  getPromoById(id: number): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getPromoById($id: Int!) {
          getPromoById(Id: $id) {
            Id
            Name
            Platform
            PromoDetail
            PromoDetailList {
              DiscountAmount
              Id
              KodePromo
            }
            PromoEnd
            PromoStart
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  getAllPromo(): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query {
          getAllPromo {
            Id
            Name
            Platform
            PromoDetail
            PromoEnd
            PromoStart
            PromoDetailList {
              DiscountAmount
              Id
              KodePromo
            }
          }
        }
      `
    })
  }

  getHotelByGeolocation(longitude: number, latitude: number): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getHotelByGeolocation($longitude: Float!, $latitude: Float!) {
          getHotelByGeolocation(longitude: $longitude, latitude: $latitude) {
            Id
            Address
            Type
            HotelFacilities {
              FacilityName
              FacilityPrice
              Id
            }
            Latitude
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Longitude
            Name
            Rating
            TiketComReview {
              AverageScore
              CleanlinessScore
              Id
              LocationScore
              Reviews {
                Conclusion
                Id
                Name
                ReviewString
                Score
                UploadDate
              }
              ServiceScore
              ValueScore
            }
            Room {
              Id
              FreeBreakfast
              FreeWifi
              Bed
              MaxPassenger
              Name
              Price
              RoomLeft
              RoomSize
            }
            TripAdvisorReview {
              AverageScore
              CleanlinessScore
              Id
              LocationScore
              Reviews {
                Conclusion
                Id
                Name
                ReviewString
                Score
                UploadDate
              }
              ServiceScore
              ValueScore
            }
          }
        }
      `,
      variables: {
        "longitude": longitude,
        "latitude": latitude
      }
    })
  }

  getEntertainmentByTrending(type: string): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getEntertainmentByTrending($type: String!) {
          getEntertainmentByTrending(Type: $type){
            EntertainmentTicket {
              Id
              Price
              Type
            }
            EventEndDate
            EventStartDate
            Id
            IsTrending
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Name
            Type
          }
        }
      `,
      variables: {
        "type": type
      }
    })
  }

  getAllTrain(): Observable<any> {
    return this.apollo.query <Query> ({
      query: gql`
        query {
          getAllTrain {
            Id
            Class
            Code
            Name
            Subclass
          }
        }
      `
    })
  }

  updateTrip(id: number, toRefer: number, arrival: String, duration: number, tax: number, service: number, trainRefer: number, fromRefer: number, departure: String, price: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation updateTrip($id: Int!, $toRefer: Int!, $arrival: DateTime!, $tax: Int!, $trainRefer: Int!, $fromRefer: Int!, $departure: DateTime!, $duration: Int!, $price: Int!, $service: Int!) {
          updateTrip(Id: $id, ToRefer: $toRefer, Arrival: $arrival, Tax: $tax, TrainRefer: $trainRefer, FromRefer: $fromRefer, Departure: $departure, Duration: $duration, Price: $price, Service: $service) {
            Id
          }
        }
      `,
      variables: {
        "id": id,
        "toRefer": toRefer,
        "arrival": arrival,
        "duration": duration,
        "tax": tax,
        "service": service,
        "trainRefer": trainRefer,
        "fromRefer": fromRefer,
        "departure": departure,
        "price": price
      }
    })
  }

  removeTrip(id: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation removeTrip($id: Int!) {
          removeTrip(Id: $id) {
            Id
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  createTrip(toRefer: number, arrival: String, duration: number, tax: number, service: number, trainRefer: number, fromRefer: number, departure: String, price: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation createTrip($toRefer: Int!, $arrival: DateTime!, $tax: Int!, $trainRefer: Int!, $fromRefer: Int!, $departure: DateTime!, $duration: Int!, $price: Int!, $service: Int!) {
          createTrip(ToRefer: $toRefer, Arrival: $arrival, Tax: $tax, TrainRefer: $trainRefer, FromRefer: $fromRefer, Departure: $departure, Duration: $duration, Price: $price, Service: $service) {
            Id
          }
        }
      `,
      variables: {
        "toRefer": toRefer,
        "arrival": arrival,
        "duration": duration,
        "tax": tax,
        "service": service,
        "trainRefer": trainRefer,
        "fromRefer": fromRefer,
        "departure": departure,
        "price": price,
      }
    })
  }

  removeFlight(id: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation removeFlight($id: Int!) {
          removeFlight(ID: $id) {
            ID
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  updateFlight(id: number, fromReferId: number, toReferId: number, departure: string, arrival: string, airlineReferId: number, price: number, tax: number, service: number, duration: number) {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation updateFlight($id: Int!, $airlineRefer: Int!, $price: Int!, $serviceCharge: Int!, $arrival: DateTime!, $toRefer: Int!, $fromRefer: Int!, $tax: Int!, $departure: DateTime!, $duration: Int!) {
          updateFlight(ID: $id, AirlineRefer: $airlineRefer, Price: $price, ServiceCharge: $serviceCharge, Arrival: $arrival, ToRefer: $toRefer, FromRefer: $fromRefer, Tax: $tax, Departure: $departure, Duration: $duration) {
            Airline {
              ID,
              Name,
            },
            Arrival,
            Departure,
            Duration,
            Facilities {
              FacilityName,
              FacilityPrice,
              id,
            },
            From {
              City,
              CityCode,
              Code,
              Country,
              Id,
              Name,
              Province,
            },
            ID,
            Price,
            Routes {
              AirportFrom {
                City,
                CityCode,
                Code,
                Country,
                Id,
                Name,
                Province,
              },
              AirportTo {
                City,
                CityCode,
                Code,
                Country,
                Id,
                Name,
                Province,
              },
              duration,
              id,
            },
            ServiceCharge,
            Tax,
            To {
              City,
              CityCode,
              Code,
              Country,
              Id,
              Name,
              Province
            }
          }
        }
      `,
      variables: {
        "id": id,
        "fromRefer": fromReferId,
        "toRefer": toReferId,
        "departure": departure,
        "arrival": arrival,
        "airlineRefer": airlineReferId,
        "price": price,
        "tax": tax,
        "serviceCharge": service,
        "duration": duration
      }
    })
  }
    
  getFlightById(id: number): Observable<any> {
    return this.apollo.query <Query>({
      query: gql`
        query getFlightById($id: Int!) {
          getFlightById(ID: $id) {
            Airline {
              ID,
              Name,
            },
            Arrival,
            Departure,
            Duration,
            Facilities {
              FacilityName,
              FacilityPrice,
              id,
            },
            From {
              City,
              CityCode,
              Code,
              Country,
              Id,
              Name,
              Province,
            },
            ID,
            Price,
            Routes {
              AirportFrom {
                City,
                CityCode,
                Code,
                Country,
                Id,
                Name,
                Province,
              },
              AirportTo {
                City,
                CityCode,
                Code,
                Country,
                Id,
                Name,
                Province,
              },
              duration,
              id,
            },
            ServiceCharge,
            Tax,
            To {
              City,
              CityCode,
              Code,
              Country,
              Id,
              Name,
              Province
            }
          }
        }
      `,
      variables: {
        "id": id
      }
    })
  }

  createFlight(fromReferId: number, toReferId: number, departure: Date, arrival: Date, airlineReferId: number, price: number, tax: number, service: number, duration: number) {
      return this.apollo.mutate<any>({
        mutation: gql`
        mutation createFlight($airlineRefer: Int!, $price: Int!, $serviceCharge: Int!, $arrival: DateTime!, $toRefer: Int!, $fromRefer: Int!, $tax: Int!, $departure: DateTime!, $duration: Int!) {
          createFlight(Price: $price, ServiceCharge: $serviceCharge, Duration: $duration, FromRefer: $fromRefer, AirlineRefer: $airlineRefer, Departure: $departure, Arrival: $arrival, ToRefer: $toRefer, Tax: $tax) {
            ID
          }
        }
        `,
        variables: {
        "fromRefer": fromReferId,
        "toRefer": toReferId,
        "departure": departure,
        "arrival": arrival,
        "airlineRefer": airlineReferId,
        "price": price,
        "tax": tax,
        "serviceCharge": service,
        "duration": duration
      }
    })
  }

  createUser(postCode: string, googleId: string, backName: string, email: string, password: string, phoneNumber: string, title: string, address: string, facebookId: string, gender: string, frontName: string, city: string) : Observable<any> {
    return this.apollo.mutate<any>({
      mutation: gql`
      mutation createUser($postCode: String!, $googleId: String!, $backName: String!, $email: String!, $password: String!, $phoneNumber: String!, $title: String!, $address: String!, $facebookId: String!, $gender: String!, $frontName: String!, $city: String!){
        createUser(PostCode: $postCode, GoogleId: $googleId, BackName: $backName, Email: $email, Password: $password, PhoneNumber: $phoneNumber, Title: $title, Address: $address, FacebookId: $facebookId, Gender: $gender, FrontName: $frontName, City: $city, MainLanguage: 'ID') {
          ID
        }
      }
      `,
      variables: {
        "frontName": frontName,
        "backName": backName,
        "email": email,
        "password": password,
        "phoneNumber": phoneNumber,
        "postCode": postCode,
        "googleId": googleId,
        "facebookId": facebookId,
        "title": title,
        "address": address,
        "gender": gender,
        "city": city
      }
    })
  }

  getAllUser(): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query {
          getAllUser {
            Address
            City
            FacebookId
            Gender
            GoogleId
            PostCode
            Title
            BackName
            Email
            FrontName
            ID
            Password
            PhoneNumber
            MainLanguage  
          }
        }
      `
    })
  }

  getAllStation(): Observable<any> {
    return this.apollo.query<any>({
      query: gql`
        query {
          getAllStation {
            Id
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Code
            Name
          }
        }
      `
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

  getAllTrip() : Observable<any> {
    return this.apollo.query<any>({
      query: gql `
        query {
          getAllTrip {
            Id
            Train {
              Id
              Name
              Class
              Subclass
              Code
            }
            To {
              Id
              Code
              Name
              Location {
                Id
                Longitude
                Latitude
                ZIndex
                City
                Province
                Country
              }
            }
            From {
              Id
              Code
              Name
              Location {
                Id
                Longitude
                Latitude
                ZIndex
                City
                Province
                Country
              }
            }
            Departure
            Arrival
            Duration
            Price
            Tax
            Service
          }
        }
      `, fetchPolicy: "no-cache"
    })
  }

  getAllFlight() : Observable<any>{
    return this.apollo.query<any>({
      query: gql`
        query {
          getAllFlight {
            ID,
            FromRefer,
            Airline {
              ID
              Name
            },
            Arrival,
            Departure,
            Duration,
            Facilities {
              FacilityName,
              FacilityPrice,
              FlightId,
              id
            },
            From {
              City,
              CityCode,
              Code,
              Country,
              Id,
              Name,
              Province
            },
            Price,
            ServiceCharge,
            Tax,
            To {
              City,
              CityCode,
              Code,
              Country,
              Id,
              Name,
              Province
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
            Address
            City
            FacebookId
            Gender
            GoogleId
            PostCode
            Title
            BackName
            Email
            FrontName
            ID
            Password
            PhoneNumber
            MainLanguage
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
          getHotelByCity(city: $arg1) {
            Id
            Address
            HotelFacilities {
              FacilityName
              FacilityPrice
              Id
            }
            Latitude
            Location {
              City
              Country
              Id
              Latitude
              Longitude
              Province
              ZIndex
            }
            Longitude
            Name
            Rating
            TiketComReview {
              AverageScore
              CleanlinessScore
              Id
              LocationScore
              Reviews {
                Conclusion
                Id
                Name
                ReviewString
                Score
                UploadDate
              }
              ServiceScore
              ValueScore
            }
            Room {
              Id
              FreeBreakfast
              FreeWifi
              Bed
              MaxPassenger
              Name
              Price
              RoomLeft
              RoomSize
            }
            TripAdvisorReview {
              AverageScore
              CleanlinessScore
              Id
              LocationScore
              Reviews {
                Conclusion
                Id
                Name
                ReviewString
                Score
                UploadDate
              }
              ServiceScore
              ValueScore
            }
          }
        }
      `,
      variables : {
        "arg1" : arg1
      }, fetchPolicy: 'no-cache'
    })
  }

  getTripByToAndFrom(arg1: string, arg2: string) : Observable<any> {
    return this.apollo.query <Query>({
      query: gql`
        query GetTripByToAndFrom($arg1: String!, $arg2: String!) {
          getTripByToAndFrom(to:$arg1, from:$arg2) {
            Id
            Train {
              Id
              Name
              Class
              Subclass
              Code
            }
            To {
              Id
              Code
              Name
              Location {
                Id
                Longitude
                Latitude
                ZIndex
                City
                Province
                Country
              }
            }
            From {
              Id
              Code
              Name
              Location {
                Id
                Longitude
                Latitude
                ZIndex
                City
                Province
                Country
              }
            }
            Departure
            Arrival
            Duration
            Price
            Tax
            Service
          }
        }
      `,
      variables: {
        "arg1": arg1,
        "arg2": arg2
      }
    })
  }

  getFlightByToAndFrom(arg1: string, arg2: string) : Observable<any> {
    return this.apollo.query <Query>({
      query: gql`
        query GetFlightByToAndFrom($arg1: String!, $arg2: String!) {
          getFlightByToAndFrom(to:$arg1, from:$arg2){
            ID
            Airline {
              ID
              Name
            }
            Arrival
            Departure
            Duration
            Facilities {
              FacilityName
              FacilityPrice
            }
            From {
              City
              CityCode
              Code
              Country
              Id
              Name
              Province
            }
            Price
            Routes {
              AirportFrom {
                City
                CityCode
                Code
                Country
                Id
                Name
                Province
              }
              AirportTo {
                City
                CityCode
                Code
                Country
                Id
                Name
                Province
              }
              duration
              id
            }
            ServiceCharge
            Tax
            To {
              City
              CityCode
              Code
              Country
              Id
              Name
              Province
            }
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
            City,
            CityCode,
            Code,
            Country,
            Id,
            Name,
            Province
          }
        }
      `
    })
  }

  getAirline() : Observable<any> {
    return this.apollo.query<any>({
      query: gql`
        query {
          getAirline {
            ID,
            Name
          }
        }
      `
    })
  }

  getCarByLocationAndToAndFrom(arg1: string, arg2: string): Observable<any> {
    return this.apollo.query <Query>({
      query: gql`
        query GetCarByLocationAndToAndFrom ($arg1: String!, $arg2: String!) {
          getCarByLocationAndToAndFrom(Location: $arg1, From: $arg2) {
            id
            AvailableAt
            Brand
            Model
            Luggage
            Passenger
            CarVendor {
              Id
              Price
              Vendor {
                Id
                Name
                VendorLocation {
                  Id
                  Location {
                    City
                    Country
                    Id
                    Latitude
                    Longitude
                    Province
                    ZIndex
                  }
                }
              }
            }
          }
        }
      `,
      variables: {
        "arg1": arg1,
        "arg2": arg2,
      }
    })
  }

  getUserByAuthId(authId: string): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getUserByAuthId($authId: String!) {
          getUserByAuthId(AuthId: $authId) {
            Address
            City
            FacebookId
            Gender
            GoogleId
            PostCode
            Title
            BackName
            Email
            FrontName
            ID
            Password
            PhoneNumber
            MainLanguage
          }
        }
      `, variables: {
        "authId": authId
      }
    })
  }

  sendMail(email: string) {
    return this.apollo.query<Query> ({
      query: gql`
        query sendMail($to: String!){
          sendMail(to: $to)
        }
      `, variables: {
        "to": email
      }
    })
  }

  getLimitedLocation(): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query {
          getLimitedLocation{
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



  removeHotel(id: number): Observable<any> {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation removeHotel($id: Int!) {
          removeHotel(id: $id) {
            Id
          }
        }
      `, variables: {
        "id": id
      }
    })
  }

  updateHotel(id: number, name: string, type: string, longitude: number, latitude: number, address: string, rating: number, locationReferId: number): Observable<any> {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation updateHotel($id: Int!, $name: String!, $type: String!, $longitude: Float!, $latitude: Float!, $address: String!, $rating: Int!, $locationReferId: Int!) {
          updateHotel(id: $id, Name: $name, Type: $type, Longitude: $longitude, Latitude: $latitude, Address: $address, Rating: $rating, LocationRefer: $locationReferId, TripAdvisorReviewReferId: 1, TiketComReviewReferId: 1){
            Id
          }
        }
      `, variables: {
        "id": id,
        "name": name,
        "type": type,
        "longitude": longitude,
        "latitude": latitude,
        "address": address,
        "rating": rating,
        "locationReferId": locationReferId
      }
    })
  }

  createHotel(name: string, type: string, longitude: number, latitude: number, address: string, rating: number, locationReferId: number): Observable<any> {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation createHotel($name: String!, $type: String!, $longitude: Float!, $latitude: Float!, $address: String!, $rating: Int!, $locationReferId: Int!) {
          createHotel(Name: $name, Type: $type, Longitude: $longitude, Latitude: $latitude, Address: $address, Rating: $rating, LocationRefer: $locationReferId, TripAdvisorReviewReferId: 1, TiketComReviewReferId: 1) {
            Id
          }
        }
      `, variables: {
        "name": name,
        "type": type,
        "longitude": longitude,
        "latitude": latitude,
        "address": address,
        "rating": rating,
        "locationReferId": locationReferId
      }
    })
  }

  getAllChat(userId: number): Observable<any> {
    return this.apollo.query<Query>({
      query: gql`
        query getAllChat($id: Int!) {
          getAllChat(UserId: $id) {
            Id
            User1
            User2
            Content
          }
        }
      `, variables: {
        "id": userId
      }, fetchPolicy: "no-cache"
    })
  }

  updateChat(id: number, content: string, user1: number, user2: number): Observable<any> {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation updateChat($id: Int!, $content: String!, $user1: Int!, $user2: Int!) {
          updateChat(Id: $id, Content: $content, User1: $user1, User2: $user2){
            Id
          }
        }
      `, variables: {
        "id": id,
        "content": content,
        "user1": user1,
        "user2": user2
      }
    })
  }

  createChat(user1: number, user2: number, content: string): Observable<any> {
    return this.apollo.mutate<any> ({
      mutation: gql`
        mutation createChat($user1: Int!, $user2: Int!, $content: String!) {
          createChat(User1: $user1, User2: $user2, Content: $content) {
            Id
          }
        }
      `, variables: {
        "user1": user1,
        "user2": user2,
        "content": content
      }
    })
  }

  getAdminByEmail(email: string): Observable<any> {
    return this.apollo.query<Query> ({
      query: gql`
        query getAdminByEmail($email: String!) {
          getAdminByEmail(Email: $email) {
            id
            email
            name
            password
          }
        }
      `, variables: {
        "email": email
      }
    })
  }

}

export class Query {
  data: any
}