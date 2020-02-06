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
					"toRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"fromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"airlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"serviceCharge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.CreateFlight,
				Description: "",
			},
			"updateFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"toRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"fromRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"airlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"tax": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"serviceCharge": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"departure": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"arrival": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.UpdateFlight,
				Description: "",
			},
			"removeFlight": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
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
					"frontName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"backName": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"phoneNumber": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"postCode": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.CreateUser,
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
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"locationRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"imagePath": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.CreateHotel,
				Description: "",
			},
			"updateHotel": &graphql.Field{
				Type: types.GetHotelType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"locationRefer": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"imagePath": &graphql.ArgumentConfig{
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
		},
		Description: "",
	})
}
