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

func GetUserType() *graphql.Object {
	if userGetter == nil {
		userGetter = graphql.NewObject(graphql.ObjectConfig{
			Name:        "userType",
			Fields:      graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"frontName": &graphql.Field{
					Type: graphql.String,
				},
				"backName": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"password": &graphql.Field{
					Type: graphql.String,
				},
				"phoneNumber": &graphql.Field{
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
				"id" : &graphql.Field {
					Type: graphql.Int,
				},
				"name" : &graphql.Field{
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
			Name: "AirportType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
				},
				"city_code": &graphql.Field{
					Type: graphql.String,
				},
				"province": &graphql.Field{
					Type: graphql.String,
				},
				"country": &graphql.Field{
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
				"facility_name": &graphql.Field{
					Type: graphql.String,
				},
				"facility_price": &graphql.Field{
					Type: graphql.Int,
				},
				"flight_id": &graphql.Field{
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"to": &graphql.Field{
					Type: GetAirportType(),
				},
				"ToRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"from": &graphql.Field{
					Type: GetAirportType(),
				},
				"facilities": &graphql.Field{
					Type: graphql.NewList(GetFacilityType()),
				},
				"FromRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"airline": &graphql.Field{
					Type: GetAirlineType(),
				},
				"airline_refer": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"tax": &graphql.Field{
					Type: graphql.Int,
				},
				"service_charge": &graphql.Field{
					Type: graphql.Int,
				},
				"departure": &graphql.Field{
					Type: graphql.DateTime,
				},
				"arrival": &graphql.Field{
					Type: graphql.DateTime,
				},
				"duration": &graphql.Field{
					Type: graphql.Int,
				},
				"routes": &graphql.Field{
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"region": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		})
	}
	return vendorGetter
}

func GetCarType() *graphql.Object {
	if carGetter == nil {
		carGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "carType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"brand": &graphql.Field{
					Type: graphql.String,
				},
				"model": &graphql.Field{
					Type: graphql.String,
				},
				"vendor": &graphql.Field{
					Type: GetVendorType(),
				},
				"vendor_refer": &graphql.Field{
					Type: graphql.Int,
				},
				"passenger": &graphql.Field{
					Type: graphql.Int,
				},
				"luggage": &graphql.Field{
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"room_size": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"location": &graphql.Field{
					Type: GetLocationType(),
				},
				"LocationRefer": &graphql.Field{
					Type: graphql.Int,
				},
				"address": &graphql.Field{
					Type: graphql.String,
				},
				"rating": &graphql.Field{
					Type: graphql.Int,
				},
				"imagePath": &graphql.Field{
					Type: graphql.String,
				},
				"HotelFacilities": &graphql.Field{
					Type: graphql.NewList(GetHotelFacilityType()),
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"code": &graphql.Field{
					Type: graphql.String,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type: graphql.String,
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"class": &graphql.Field{
					Type: graphql.String,
				},
				"subclass": &graphql.Field{
					Type: graphql.String,
				},
				"code": &graphql.Field{
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
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"train": &graphql.Field{
					Type: GetTrainType(),
				},
				"train_refer": &graphql.Field{
					Type: graphql.Int,
				},
				"from": &graphql.Field{
					Type: GetStationType(),
				},
				"from_refer": &graphql.Field{
					Type: graphql.Int,
				},
				"to": &graphql.Field{
					Type: GetStationType(),
				},
				"to_refer": &graphql.Field{
					Type: graphql.Int,
				},
				"departure": &graphql.Field{
					Type: graphql.DateTime,
				},
				"arrival": &graphql.Field{
					Type: graphql.DateTime,
				},
				"duration": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"tax": &graphql.Field{
					Type: graphql.Int,
				},
				"service": &graphql.Field{
					Type: graphql.Int,
				},
			},
		})
	}
	return tripGetter
}

