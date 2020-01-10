package types

import (
	"github.com/graphql-go/graphql"
)

var adminGetter *graphql.Object
var userGetter *graphql.Object
var airlineGetter *graphql.Object
var flightGetter *graphql.Object

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
				"code" : &graphql.Field {
						Type: graphql.String,
				},
			},
		})
	}
	return airlineGetter
}

func GetFlightType() *graphql.Object {
	if flightGetter == nil {
		flightGetter = graphql.NewObject(graphql.ObjectConfig{
			Name: "flightType",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"destination": &graphql.Field{
					Type: graphql.String,
				},
				"from": &graphql.Field{
					Type: graphql.String,
				},
				"flight_date_from": &graphql.Field{
					Type: graphql.DateTime,
				},
				"flight_date_to": &graphql.Field{
					Type: graphql.DateTime,
				},
				"airline_refer": &graphql.Field{
					Type: graphql.Int,
				},
				"price": &graphql.Field{
					Type: graphql.Int,
				},
				"airline": &graphql.Field{
					Type: GetAirlineType(),
				},
			},
		})
	}
	return flightGetter
}