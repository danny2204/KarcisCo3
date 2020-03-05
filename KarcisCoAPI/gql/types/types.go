package types

import (
	"github.com/graphql-go/graphql"
)

var adminGetter *graphql.Object
var userGetter *graphql.Object
var airlineGetter *graphql.Object
var flightGetter *graphql.Object
var airportGetter *graphql.Object
var routeGetter *graphql.Object
var carGetter *graphql.Object
var vendorGetter *graphql.Object
var roomGetter *graphql.Object
var locationGetter *graphql.Object
var hotelGetter *graphql.Object
var stationGetter *graphql.Object
var trainGetter *graphql.Object
var tripGetter *graphql.Object
var facilityGetter *graphql.Object
var hotelFacilityGetter *graphql.Object
var vendorLocationGetter *graphql.Object
var carVendorGetter *graphql.Object
var reviewGetter *graphql.Object
var tripAdvisorReviewGetter *graphql.Object
var tiketComReviewGetter *graphql.Object
var entertainmentGetter *graphql.Object
var entertainmentTicketGetter *graphql.Object
var blogGetter *graphql.Object
var promoGetter *graphql.Object
var promoDetailGetter *graphql.Object
var chatGetter *graphql.Object

func GetUserType() *graphql.Object {
	if userGetter == nil {
		userGetter = graphql.NewObject(graphql.ObjectConfig{
			Name:        "userType",
			Fields:      graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"FrontName": &graphql.Field{
					Type: graphql.String,
				},
				"BackName": &graphql.Field{
					Type: graphql.String,
				},
				"Email": &graphql.Field{
					Type: graphql.String,
				},
				"Password": &graphql.Field{
					Type: graphql.String,
				},
				"PhoneNumber": &graphql.Field{
					Type: graphql.String,
				},
				"GoogleId": &graphql.Field{
					Type: graphql.String,
				},
				"FacebookId": &graphql.Field{
					Type: graphql.String,
				},
				"Gender": &graphql.Field{
					Type: graphql.String,
				},
				"Title": &graphql.Field{
					Type: graphql.String,
				},
				"City": &graphql.Field{
					Type: graphql.String,
				},
				"Address": &graphql.Field{
					Type: graphql.String,
				},
				"PostCode": &graphql.Field{
					Type: graphql.String,
				},
				"MainLanguage": &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "",
		})
	}
	return userGetter
}

func GetAdminType() *graphql.Object {
	if adminGetter == nil {
		adminGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "adminType",
			Fields: graphql.Fields{
				"id" : &graphql.Field{
					Type: graphql.Int,
				},
				"name" : &graphql.Field{
					Type: graphql.String,
				},
				"email" : &graphql.Field{
					Type: graphql.String,
				},
				"password" : &graphql.Field{
					Type: graphql.String,
				},
			},
			Description: "",
		})
	}
	return adminGetter
}

func GetAirlineType() *graphql.Object {
	if airlineGetter == nil {
		airlineGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "airlineType",
			Fields: graphql.Fields{
				"ID" : &graphql.Field {
					Type: graphql.Int,
				},
				"Name" : &graphql.Field{
						Type: graphql.String,
				},
			},
		})
	}
	return airlineGetter
}

func GetAirportType() *graphql.Object {
	if airportGetter == nil {
		airportGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "airportType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Code": &graphql.Field{
					Type: graphql.String,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"City": &graphql.Field{
					Type: graphql.String,
				},
				"CityCode": &graphql.Field{
					Type: graphql.String,
				},
				"Province": &graphql.Field{
					Type: graphql.String,
				},
				"Country": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return airportGetter
}

func GetHotelFacilityType() *graphql.Object {
	if hotelFacilityGetter == nil {
		hotelFacilityGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "hotelFacilityType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"FacilityName": &graphql.Field{
					Type: graphql.String,
				},
				"FacilityPrice": &graphql.Field{
					Type: graphql.Int,
				},
				"HotelReferId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return hotelFacilityGetter
}

func GetRouteType() *graphql.Object {
	if routeGetter == nil {
		routeGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "routeType",
			Fields: graphql.Fields{
				"id": &graphql.Field{Type: graphql.Int,
				},
				"AirportFrom": &graphql.Field{
					Type: GetAirportType(),
				},
				"AirportTo": &graphql.Field{
					Type: GetAirportType(),
				},
				"RouteFromRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"RouteToRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"duration": &graphql.Field{
					Type: graphql.Int,
				},
				"FlightRouteId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return routeGetter
}

func GetFacilityType() *graphql.Object {
	if facilityGetter == nil {
		facilityGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "facilityType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"FacilityName": &graphql.Field{
					Type: graphql.String,
				},
				"FacilityPrice": &graphql.Field{
					Type: graphql.Int,
				},
				"FlightId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return facilityGetter
}

func GetFlightType() *graphql.Object {
	if flightGetter == nil {
		flightGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "flightType",
			Fields: graphql.Fields{
				"ID": &graphql.Field{
					Type: graphql.Int,
				},
				"To": &graphql.Field{
					Type: GetAirportType(),
				},
				"ToRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"From": &graphql.Field{
					Type: GetAirportType(),
				},
				"Facilities": &graphql.Field{
					Type: graphql.NewList(GetFacilityType()),
				},
				"FromRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"Airline": &graphql.Field{
					Type: GetAirlineType(),
				},
				"AirlineRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"Price": &graphql.Field{
					Type: graphql.Int,
				},
				"Tax": &graphql.Field{
					Type: graphql.Int,
				},
				"ServiceCharge": &graphql.Field{
					Type: graphql.Int,
				},
				"Departure": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Arrival": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Duration": &graphql.Field{
					Type: graphql.Int,
				},
				"Routes": &graphql.Field{
					Type: graphql.NewList(GetRouteType()),
				},
			},
		})
	}
	return flightGetter
}

func GetVendorType() *graphql.Object {
	if vendorGetter == nil {
		vendorGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "vendorType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"VendorLocation": &graphql.Field{
					Type: graphql.NewList(GetVendorLocationType()),
				},
			},
		})
	}
	return vendorGetter
}

func GetCarVendorType() *graphql.Object {
	if carVendorGetter == nil {
		carVendorGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "carVendorType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Vendor": &graphql.Field{
					Type: GetVendorType(),
				},
				"VendorReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"CarReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"Price": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return carVendorGetter
}

func GetCarType() *graphql.Object {
	if carGetter == nil {
		carGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "carType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"Brand": &graphql.Field{
					Type: graphql.String,
				},
				"Model": &graphql.Field{
					Type: graphql.String,
				},
				"CarVendor": &graphql.Field{
					Type: graphql.NewList(GetCarVendorType()),
				},
				"AvailableAt": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Passenger": &graphql.Field{
					Type: graphql.Int,
				},
				"Luggage": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
 	return carGetter
}

func GetRoomType() *graphql.Object {
	if roomGetter == nil {
		roomGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "roomType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"RoomSize": &graphql.Field{
					Type: graphql.Int,
				},
				"Price": &graphql.Field{
					Type: graphql.Int,
				},
				"MaxPassenger": &graphql.Field{
					Type: graphql.Int,
				},
				"FreeBreakfast": &graphql.Field{
					Type: graphql.Boolean,
				},
				"FreeWifi": &graphql.Field{
					Type: graphql.Boolean,
				},
				"Bed": &graphql.Field{
					Type: graphql.String,
				},
				"RoomLeft": &graphql.Field{
					Type: graphql.Int,
				},
				"RoomReferId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return roomGetter
}

func GetLocationType() *graphql.Object {
	if locationGetter == nil {
		locationGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "locationType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Longitude": &graphql.Field{
					Type: graphql.Float,
				},
				"Latitude": &graphql.Field{
					Type: graphql.Float,
				},
				"City": &graphql.Field{
					Type: graphql.String,
				},
				"Province": &graphql.Field{
					Type: graphql.String,
				},
				"Country": &graphql.Field{
					Type: graphql.String,
				},
				"ZIndex": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return locationGetter
}

func GetHotelType() *graphql.Object {
	if hotelGetter == nil {
		hotelGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "hotelType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Location": &graphql.Field{
					Type: GetLocationType(),
				},
				"LocationRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"Address": &graphql.Field{
					Type: graphql.String,
				},
				"Rating": &graphql.Field{
					Type: graphql.Float,
				},
				"Longitude": &graphql.Field{
					Type: graphql.Float,
				},
				"Latitude": &graphql.Field{
					Type: graphql.Float,
				},
				"Type": &graphql.Field{
					Type: graphql.String,
				},
				"TripAdvisorReview": &graphql.Field{
					Type: GetTripAdvisorReviewType(),
				},
				"TripAdvisorReviewReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"TiketComReview": &graphql.Field{
					Type: GetTiketComReviewType(),
				},
				"TiketComReviewReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"HotelFacilities": &graphql.Field{
					Type: graphql.NewList(GetHotelFacilityType()),
				},
				"Room": &graphql.Field{
					Type: graphql.NewList(GetRoomType()),
				},
			},
		})
	}
	return hotelGetter
}

func GetStationType() *graphql.Object {
	if stationGetter == nil {
		stationGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "stationType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Code": &graphql.Field{
					Type: graphql.String,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Location": &graphql.Field{
					Type: GetLocationType(),
				},
				"LocationReferId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return stationGetter
}

func GetTrainType() *graphql.Object {
	if trainGetter == nil {
		trainGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "trainType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Class": &graphql.Field{
					Type: graphql.String,
				},
				"Subclass": &graphql.Field{
					Type: graphql.String,
				},
				"Code": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return trainGetter
}

func GetTripType() *graphql.Object {
	if tripGetter == nil {
		tripGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "tripType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Train": &graphql.Field{
					Type: GetTrainType(),
				},
				"TrainRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"From": &graphql.Field{
					Type: GetStationType(),
				},
				"FromRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"To": &graphql.Field{
					Type: GetStationType(),
				},
				"ToRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"Departure": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Arrival": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Duration": &graphql.Field{
					Type: graphql.Int,
				},
				"Price": &graphql.Field{
					Type: graphql.Int,
				},
				"Tax": &graphql.Field{
					Type: graphql.Int,
				},
				"Service": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return tripGetter
}

func GetVendorLocationType() *graphql.Object {
	if vendorLocationGetter == nil {
		vendorLocationGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "vendorLocationType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Location": &graphql.Field{
					Type: GetLocationType(),
				},
				"LocationRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"VendorLocationReferId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return vendorLocationGetter
}

func GetReviewType() *graphql.Object {
	if reviewGetter == nil {
		reviewGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "reviewType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Score": &graphql.Field{
					Type: graphql.Float,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"UploadDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Conclusion": &graphql.Field{
					Type: graphql.String,
				},
				"ReviewString": &graphql.Field{
					Type: graphql.String,
				},
				"TiketComReviewReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"TripAdvisorReviewReferId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return reviewGetter
}

func GetTripAdvisorReviewType() *graphql.Object {
	if tripAdvisorReviewGetter == nil {
		tripAdvisorReviewGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "tripAdvisorReviewType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"AverageScore": &graphql.Field{
					Type: graphql.Float,
				},
				"LocationScore": &graphql.Field{
					Type: graphql.Float,
				},
				"CleanlinessScore": &graphql.Field{
					Type: graphql.Float,
				},
				"ServiceScore": &graphql.Field{
					Type: graphql.Float,
				},
				"ValueScore": &graphql.Field{
					Type: graphql.Float,
				},
				"Reviews": &graphql.Field{
					Type: graphql.NewList(GetReviewType()),
				},
			},
		})
	}
	return tripAdvisorReviewGetter
}

func GetTiketComReviewType() *graphql.Object {
	if tiketComReviewGetter == nil {
		tiketComReviewGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "ticketComReviewType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"AverageScore": &graphql.Field{
					Type: graphql.Float,
				},
				"LocationScore": &graphql.Field{
					Type: graphql.Float,
				},
				"CleanlinessScore": &graphql.Field{
					Type: graphql.Float,
				},
				"ServiceScore": &graphql.Field{
					Type: graphql.Float,
				},
				"ValueScore": &graphql.Field{
					Type: graphql.Float,
				},
				"Reviews": &graphql.Field{
					Type: graphql.NewList(GetReviewType()),
				},
			},
		})
	}
	return tiketComReviewGetter
}

func GetEntertainmentType() *graphql.Object {
	if entertainmentGetter == nil {
		entertainmentGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "entertainmentType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"Type": &graphql.Field{
					Type: graphql.String,
				},
				"LocationReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"Location": &graphql.Field{
					Type: GetLocationType(),
				},
				"EventStartDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"EventEndDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"IsTrending": &graphql.Field{
					Type: graphql.Boolean,
				},
				"EntertainmentTicket": &graphql.Field{
					Type: graphql.NewList(GetEntertainmentTicketType()),
				},
				"UploadDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Description": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return entertainmentGetter
}

func GetEntertainmentTicketType() *graphql.Object {
	if entertainmentTicketGetter == nil {
		entertainmentTicketGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "entertainmentTicketType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Type": &graphql.Field{
					Type: graphql.String,
				},
				"Price": &graphql.Field{
					Type: graphql.Int,
				},
				"EntertainmentReferId": &graphql.Field{
					Type: graphql.Int,
				},
				"Description": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return entertainmentTicketGetter
}

func GetBlogType() *graphql.Object {
	if blogGetter == nil {
		blogGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "blogType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"BlogTitle": &graphql.Field{
					Type: graphql.String,
				},
				"BlogDetail": &graphql.Field{
					Type: graphql.String,
				},
				"UploadDate": &graphql.Field{
					Type: graphql.DateTime,
				},
				"View": &graphql.Field{
					Type: graphql.Int,
				},
				"ImagePath": &graphql.Field{
					Type: graphql.String,
				},
				"Category": &graphql.Field{
					Type: graphql.String,
				},
				"AuthorName": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return blogGetter
}

func GetPromoType() *graphql.Object {
	if promoGetter == nil {
		promoGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "promoType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"Name": &graphql.Field{
					Type: graphql.String,
				},
				"PromoStart": &graphql.Field{
					Type: graphql.DateTime,
				},
				"PromoEnd": &graphql.Field{
					Type: graphql.DateTime,
				},
				"Platform": &graphql.Field{
					Type: graphql.String,
				},
				"PromoDetail": &graphql.Field{
					Type: graphql.String,
				},
				"PromoDetailList": &graphql.Field{
					Type: graphql.NewList(GetPromoDetailType()),
				},
			},
		})
	}
	return promoGetter
}

func GetPromoDetailType() *graphql.Object {
	if promoDetailGetter == nil {
		promoDetailGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "promoDetailType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"DiscountAmount": &graphql.Field{
					Type: graphql.Int,
				},
				"KodePromo": &graphql.Field{
					Type: graphql.String,
				},
				"PromoReferId": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return promoDetailGetter
}

func GetChatType() *graphql.Object {
	if chatGetter == nil {
		chatGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "chatType",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.Int,
				},
				"User1": &graphql.Field{
					Type: graphql.Int,
				},
				"User2": &graphql.Field{
					Type: graphql.Int,
				},
				"Content": &graphql.Field{
					Type: graphql.String,
				},
				"UpdatedAt": &graphql.Field{
					Type: graphql.DateTime,
				},
			},
		})
	}
	return chatGetter
}
