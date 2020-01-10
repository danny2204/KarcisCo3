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
					"destination": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"flightDateFrom": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"flightDateTo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"airlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
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
					"destination": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"flightDateFrom": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"flightDateTo": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.DateTime),
					},
					"airlineRefer": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"price": &graphql.ArgumentConfig{
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
			"getFlightById": &graphql.Field{
				Type: types.GetFlightType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetFlightById,
				Description: "",
			},
			"createAirline": &graphql.Field{
				Type: types.GetAirlineType(),
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"code": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"seat": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
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
					"code": &graphql.ArgumentConfig{
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
		},
		Description: "",
	})
}
