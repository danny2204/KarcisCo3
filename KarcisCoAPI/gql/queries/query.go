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
		},
		Description: "",
	})
}