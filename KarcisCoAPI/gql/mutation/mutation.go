package mutation

import (
	"github.com/danny2204/KarcisCoAPI/gql/resolver"
	"github.com/danny2204/KarcisCoAPI/gql/types"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"ToRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"FromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"AirlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"ServiceCharge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"Arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"Duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.CreateFlight,
				Description: "",
			},
			"updateFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"ToRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"FromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"AirlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"ServiceCharge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"Departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"Arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"Duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.UpdateFlight,
				Description: "",
			},
			"removeFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.RemoveFlight,
				Description: "",
			},
			//"getFlightById": &graphql.Field{
			//	Type: types.GetFlightType(),
			//	Args: graphql.FieldConfigArgument{
			//		"id": &graphql.ArgumentConfig{
			//			Type: graphql.Int,
			//		},
			//	},
			//	Resolve: resolver.GetFlightById,
			//	Description: "",
			//},
			"createAirline": &graphql.Field{
				Type: types.GetAirlineType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.CreateAirline,
				Description: "",
			},
			"updateAirline": &graphql.Field{
				Type: types.GetAirlineType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.UpdateAirline,
				Description: "",
			},
			"removeAirline": &graphql.Field{
				Type: types.GetAirlineType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.RemoveAirline,
				Description: "",
			},
			"getAirlineById": &graphql.Field{
				Type: types.GetAirlineType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.GetAirlineById,
				Description: "",
			},
			"createAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolver.CreateAdmin,
				Description: "",
			},
			"updateAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve:     resolver.UpdateAdmin,
				Description: "",
			},
			"removeAdmin": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve:	resolver.RemoveAdmin,
				Description: "",
			},
			"getAdminById": &graphql.Field{
				Type: types.GetAdminType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.GetAdminById,
				Description: "",
			},
			"createUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"FrontName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"BackName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"PhoneNumber": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"City": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"PostCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"GoogleId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"FacebookId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Gender": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"MainLanguage": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateUser,
				Description: "",
			},
			"updateUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"FrontName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"BackName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"PhoneNumber": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"City": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"Address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"PostCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"GoogleId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"FacebookId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Gender": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"MainLanguage": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateUser,
				Description: "",
			},
			"removeUser": &graphql.Field{
				Type: types.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.RemoveUser,
				Description: "",
			},
			"createAirport": &graphql.Field{
				Type: types.GetAirportType(),
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"cityCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.CreateAirport,
				Description: "",
			},
			"createFacility": &graphql.Field{
				Type: types.GetFacilityType(),
				Args: graphql.FieldConfigArgument{
					"facilityName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"facilityPrice": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"flightId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.CreateFacility,
				Description: "",
			},
			"updateFacility": &graphql.Field{
				Type: types.GetFacilityType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"facilityName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"facilityPrice": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"flightId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.UpdateFacility,
				Description: "",
			},
			"removeFacility": &graphql.Field{
				Type: types.GetFacilityType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.RemoveFacility,
				Description: "",
			},
			"createRoute": &graphql.Field{
				Type: types.GetRouteType(),
				Args: graphql.FieldConfigArgument{
					"routeToRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"routeFromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"flightRouteId": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.CreateRoute,
				Description: "",
			},
			"updateRoute": &graphql.Field{
				Type: types.GetRouteType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"routeToRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"routeFromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.UpdateRoute,
				Description: "",
			},
			"removeRoute": &graphql.Field{
				Type: types.GetRouteType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveRoute,
				Description: "",
			},
			"removeAirport": &graphql.Field{
				Type: types.GetAirportType(),
				Args: graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveAirport,
				Description: "",
			},
			"createHotelFacility": &graphql.Field{
				Type: types.GetHotelFacilityType(),
				Args: graphql.FieldConfigArgument{
					"facilityName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"facilityPrice": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"hotelPreferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateHotelFaclity,
				Description: "",
			},
			"updateHotelFacility": &graphql.Field{
				Type: types.GetHotelFacilityType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"facilityName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"facilityPrice": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"hotelPreferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateHotelFacility,
				Description: "",
			},
			"removeHotelFacility": &graphql.Field{
				Type: types.GetHotelFacilityType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveHotelFacility,
				Description: "",
			},
			"createHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"LocationRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Rating": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"Latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"TripAdvisorReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"TiketComReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateHotel,
				Description: "",
			},
			"updateHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"LocationRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Rating": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"Latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"TripAdvisorReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"TiketComReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateHotel,
				Description: "",
			},
			"removeHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveHotel,
				Description: "",
			},
			"createLocation": &graphql.Field{
				Type: types.GetLocationType(),
				Args: graphql.FieldConfigArgument{
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"county": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"zIndex": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateLocation,
				Description: "",
			},
			"updateLocation": &graphql.Field{
				Type: types.GetLocationType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"country": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"zIndex": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateLocation,
				Description: "",
			},
			"removeLocation": &graphql.Field{
				Type: types.GetLocationType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveLocation,
				Description: "",
			},
			"createStation": &graphql.Field{
				Type: types.GetStationType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"locationReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateStation,
				Description: "",
			},
			"updateStation": &graphql.Field{
				Type: types.GetStationType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"LocationReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateStation,
				Description: "",
			},
			"removeStation": &graphql.Field{
				Type: types.GetStationType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveStation,
				Description: "",
			},
			"createTrip": &graphql.Field{
				Type: types.GetTripType(),
				Args: graphql.FieldConfigArgument{
					"TrainRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"FromRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"ToRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Departure": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Arrival": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Tax": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Service": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateTrip,
				Description: "",
			},
			"updateTrip": &graphql.Field{
				Type: types.GetTripType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"TrainRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"FromRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"ToRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Departure": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Arrival": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Tax": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Service": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateTrip,
				Description: "",
			},
			"removeTrip": &graphql.Field{
				Type: types.GetStationType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveTrip,
				Description: "",
			},
			"createTrain": &graphql.Field{
				Type: types.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Class": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Subclass": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateTrain,
				Description: "",
			},
			"updateTrain": &graphql.Field{
				Type: types.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Class": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Subclass": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Code": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateTrain,
				Description: "",
			},
			"removeTrain": &graphql.Field{
				Type: types.GetTrainType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveTrain,
				Description: "",
			},
			"createCar": &graphql.Field{
				Type: types.GetCarType(),
				Args: graphql.FieldConfigArgument{
					"Brand": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Passenger": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Luggage": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"AvailableAt": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
				},
				Resolve: resolver.CreateCar,
				Description: "",
			},
			"updateCar": &graphql.Field{
				Type: types.GetCarType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Brand": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Model": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Passenger": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Luggage": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"AvailableAt": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
				},
				Resolve: resolver.UpdateCar,
				Description: "",
			},
			"removeCar": &graphql.Field{
				Type: types.GetCarType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveCar,
				Description: "",
			},
			"createCarVendor": &graphql.Field{
				Type: types.GetCarVendorType(),
				Args: graphql.FieldConfigArgument{
					"VendorReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"CarReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateCarVendor,
				Description: "",
			},
			"updateCarVendor": &graphql.Field{
				Type: types.GetCarVendorType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"VendorReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"CarReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateCarVendor,
				Description: "",
			},
			"removeCarVendor": &graphql.Field{
				Type: types.GetCarVendorType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveCarVendor,
				Description: "",
			},
			"createVendor": &graphql.Field{
				Type: types.GetVendorType(),
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateVendor,
				Description: "",
			},
			"updateVendor": &graphql.Field{
				Type: types.GetVendorType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateVendor,
				Description: "",
			},
			"removeVendor": &graphql.Field{
				Type: types.GetVendorType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveVendor,
			},
			"createVendorLocation": &graphql.Field{
				Type: types.GetVendorLocationType(),
				Args: graphql.FieldConfigArgument{
					"LocationRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"VendorLocationReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateVendorLocation,
				Description: "",
			},
			"updateVendorLocation": &graphql.Field{
				Type: types.GetVendorLocationType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"LocationRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"VendorLocationReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateVendorLocation,
				Description: "",
			},
			"removeVendorLocation": &graphql.Field{
				Type: types.GetVendorLocationType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveVendorLocation,
				Description: "",
			},
			"createRoom": &graphql.Field{
				Type: types.GetRoomType(),
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"RoomSize": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"MaxPassenger": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"FreeBreakfast": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"FreeWifi": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"Bed": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"RoomLeft": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"RoomReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateRoom,
				Description: "",
			},
			"updateRoom": &graphql.Field{
				Type: types.GetRoomType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"RoomSize": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"MaxPassenger": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"FreeBreakfast": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"FreeWifi": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"Bed": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"RoomLeft": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"RoomReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateRoom,
				Description: "",
			},
			"removeRoom": &graphql.Field{
				Type: types.GetRoomType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveRoom,
				Description: "",
			},
			"createReview": &graphql.Field{
				Type: types.GetReviewType(),
				Args: graphql.FieldConfigArgument{
					"Score": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"UploadDate": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Conclusion": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ReviewString": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"TiketComReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"TripAdvisorReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreateReview,
				Description: "",
			},
			"updateReview": &graphql.Field{
				Type: types.GetReviewType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Score": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"UploadDate": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Conclusion": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ReviewString": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"TiketComReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"TripAdvisorReviewReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.UpdateReview,
				Description: "",
			},
			"removeReview": &graphql.Field{
				Type: types.GetReviewType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.RemoveReview,
				Description: "",
			},
			"createTripAdvisorReview": &graphql.Field{
				Type: types.GetTripAdvisorReviewType(),
				Args: graphql.FieldConfigArgument{
					"AverageScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"LocationScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"CleanlinessScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ServiceScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ValueScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: resolver.CreateTripAdvisorReview,
				Description: "",
			},
			"updateTripAdvisorReview": &graphql.Field{
				Type: types.GetTripAdvisorReviewType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"AverageScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"LocationScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"CleanlinessScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ServiceScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ValueScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: resolver.UpdateTripAdvisorReview,
				Description: "",
			},
			"removeTripAdvisorReview": &graphql.Field{
				Type: types.GetTripAdvisorReviewType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveTripAdvisorReview,
				Description: "",
			},
			"createTiketComReview": &graphql.Field{
				Type: types.GetTiketComReviewType(),
				Args: graphql.FieldConfigArgument{
					"AverageScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"LocationScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"CleanlinessScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ServiceScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ValueScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: resolver.CreateTiketComReview,
				Description: "",
			},
			"updateTiketComReview": &graphql.Field{
				Type: types.GetTiketComReviewType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"AverageScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"LocationScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"CleanlinessScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ServiceScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ValueScore": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: resolver.UpdateTiketComReview,
				Description: "",
			},
			"removeTiketComReview": &graphql.Field{
				Type: types.GetTiketComReviewType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveTiketComReview,
				Description: "",
			},
			"createEntertainment": &graphql.Field{
				Type: types.GetEntertainmentType(),
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"LocationReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"EventStartDate": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"EventEndDate": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"IsTrending": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"Description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateEntertainment,
				Description: "",
			},
			"updateEntertainment": &graphql.Field{
				Type: types.GetEntertainmentType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"LocationReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"EventStartDate": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"EventEndDate": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"IsTrending": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"Description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateEntertainment,
				Description: "",
			},
			"removeEntertainment": &graphql.Field{
				Type: types.GetEntertainmentType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveEntertainment,
				Description: "",
			},
			"createEntertainmentTicket": &graphql.Field{
				Type: types.GetEntertainmentTicketType(),
				Args: graphql.FieldConfigArgument{
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"EntertainmentReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateEntertainmentTicket,
				Description: "",
			},
			"updateEntertainmentTicket": &graphql.Field{
				Type: types.GetEntertainmentTicketType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Price": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"EntertainmentReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateEntertainmentTicket,
				Description: "",
			},
			"removeEntertainmentTicket": &graphql.Field{
				Type: types.GetEntertainmentTicketType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveEntertainmentTicket,
				Description: "",
			},
			"createPromo": &graphql.Field{
				Type: types.GetPromoType(),
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"PromoStart": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"PromoEnd": &graphql.ArgumentConfig{
						Type: graphql.DateTime,
					},
					"Platform": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"PromoDetail": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreatePromo,
				Description: "",
			},
			"createPromoDetail": &graphql.Field{
				Type: types.GetPromoDetailType(),
				Args: graphql.FieldConfigArgument{
					"DiscountAmount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"KodePromo": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"PromoReferId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.CreatePromoDetail,
				Description: "",
			},
			"createBlog": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"BlogTitle": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"BlogDetail": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"View": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"AuthorName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ImagePath": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateBlog,
				Description: "",
			},
			"updateBlog": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"BlogTitle": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"BlogDetail": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"View": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"AuthorName": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"Category": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"ImagePath": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateBlog,
				Description: "",
			},
			"removeBlog": &graphql.Field{
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.RemoveBlog,
				Description: "",
			},
			"createChat": &graphql.Field{
				Type: types.GetChatType(),
				Args: graphql.FieldConfigArgument{
					"User1": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"User2": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateChat,
				Description: "",
			},
			"updateChat": &graphql.Field{
				Type: types.GetChatType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"User1": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"User2": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"Content": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.UpdateChat,
				Description: "",
			},
		},
		Description: "",
	})
}
