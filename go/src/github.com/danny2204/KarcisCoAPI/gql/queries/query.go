package queries

import (
	"github.com/danny2204/KarcisCoAPI/gql/resolver"
	"github.com/danny2204/KarcisCoAPI/gql/types"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "RootQuery",
		Fields:      graphql.Fields{
			"getAdmins": {
				Type: graphql.NewList(types.GetAdminType()),
				Resolve: resolver.GetAllAdmin,
				Description: "",
			},
			"getUsers": {
				Type: graphql.NewList(types.GetUserType()),
				Resolve: resolver.GetAllUser,
				Description: "",
			},
			"getUserByEmailAndPhone":{
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"arg": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.GetUserByEmailAndPhone,
				Description: "",
			},
			"getFlightByToAndFrom": {
				Type: graphql.NewList(types.GetFlightType()),
				Args: graphql.FieldConfigArgument{
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.GetFlightByToAndFrom,
				Description: "",
			},
			"getHotelByCity": {
				Type: graphql.NewList(types.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.GetHotelByCity,
				Description: "",
			},
			"getAllFlight": {
				Type: graphql.NewList(types.GetFlightType()),
				Resolve: resolver.GetAllFlight,
				Description: "",
			},
			"getAirline": {
				Type: graphql.NewList(types.GetAirlineType()),
				Resolve: resolver.GetAllAirline,
				Description: "",
			},
			"getAllAirport": {
				Type: graphql.NewList(types.GetAirportType()),
				Resolve: resolver.GetAllAirport,
				Description: "",
			},
			"getAllFacility": {
				Type: graphql.NewList(types.GetFacilityType()),
				Resolve: resolver.GetAllFacility,
				Description: "",
			},
			"getAllRoute": {
				Type: graphql.NewList(types.GetRouteType()),
				Resolve: resolver.GetAllRoute,
				Description: "",
			},
			"getAllHotelFacility": {
				Type: graphql.NewList(types.GetHotelFacilityType()),
				Resolve: resolver.GetAllHotelFacility,
				Description: "",
			},
			"getAllHotel": {
				Type: graphql.NewList(types.GetHotelType()),
				Resolve: resolver.GetAllHotel,
				Description: "",
			},
			"getAllLocation": {
				Type: graphql.NewList(types.GetLocationType()),
				Resolve: resolver.GetAllLocation,
				Description: "",
			},
		},
		Description: "",
	})
}