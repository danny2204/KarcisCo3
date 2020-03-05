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
			"getAdminByEmail": {
				Type: graphql.NewList(types.GetAdminType()),
				Args: graphql.FieldConfigArgument{
					"Email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.GetAdminByEmail,
				Description: "",
			},
			"getUsers": {
				Type: graphql.NewList(types.GetUserType()),
				Resolve: resolver.GetAllUser,
				Description: "",
			},
			"getUserByAuthId": {
				Type: graphql.NewList(types.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"AuthId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.GetUserByAuthId,
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
			"getFlightById": {
				Type: graphql.NewList(types.GetFlightType()),
				Args: graphql.FieldConfigArgument{
					"ID": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: resolver.GetFlightById,
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
			"getLimitedLocation": {
				Type: graphql.NewList(types.GetLocationType()),
				Resolve: resolver.GetLimitedLocation,
				Description: "",
			},
			"getAllStation": {
				Type: graphql.NewList(types.GetStationType()),
				Resolve: resolver.GetAllStation,
				Description: "",
			},
			"getAllTrip": {
				Type: graphql.NewList(types.GetTripType()),
				Resolve: resolver.GetAllTrip,
				Description: "",
			},
			"getTripByToAndFrom": {
				Type: graphql.NewList(types.GetTripType()),
				Args: graphql.FieldConfigArgument{
					"to": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"from": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.GetTripByToAndFrom,
				Description: "",
			},
			"getCarByLocationAndToAndFrom": {
				Type: graphql.NewList(types.GetCarType()),
				Args: graphql.FieldConfigArgument{
					"Location": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"From": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: resolver.GetCarByLocationAndFromAndTo,
				Description: "",
			},
			"getAllTrain": {
				Type: graphql.NewList(types.GetTrainType()),
				Resolve: resolver.GetAllTrain,
				Description: "",
			},
			"getAllVendorLocation": {
				Type: graphql.NewList(types.GetVendorLocationType()),
				Resolve: resolver.GetAllVendorLocation,
				Description: "",
			},
			"getAllVendor": {
				Type: graphql.NewList(types.GetVendorType()),
				Resolve: resolver.GetAllVendor,
				Description: "",
			},
			"getAllCarVendor": {
				Type: graphql.NewList(types.GetCarVendorType()),
				Resolve: resolver.GetAllCarVendor,
				Description: "",
			},
			"getAllCar": {
				Type: graphql.NewList(types.GetCarType()),
				Resolve: resolver.GetAllCar,
				Description: "",
			},
			"getAllRoom": {
				Type: graphql.NewList(types.GetRoomType()),
				Resolve: resolver.GetAllRoom,
				Description: "",
			},
			"getAllReview": {
				Type: graphql.NewList(types.GetReviewType()),
				Resolve: resolver.GetAllReview,
				Description: "",
			},
			"getAllTripAdvisorReview": {
				Type: graphql.NewList(types.GetTripAdvisorReviewType()),
				Resolve: resolver.GetAllTripAdvisorReview,
				Description: "",
			},
			"getAllTiketComReview": {
				Type: graphql.NewList(types.GetTiketComReviewType()),
				Resolve: resolver.GetAllTiketComReview,
				Description: "",
			},
			"getAllEntertainment": {
				Type: graphql.NewList(types.GetEntertainmentType()),
				Resolve: resolver.GetAllEntertainment,
				Description: "",
			},
			"getEntertainmentByLocation": {
				Type: graphql.NewList(types.GetEntertainmentType()),
				Args: graphql.FieldConfigArgument{
					"Location": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.GetEntertainmentByLocation,
				Description: "",
			},
			"getEntertainmentByType": {
				Type: graphql.NewList(types.GetEntertainmentType()),
				Args: graphql.FieldConfigArgument{
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.GetEntertainmentByType,
				Description: "",
			},
			"getEntertainmentById":  {
				Type: types.GetEntertainmentType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetEntertainmentById,
				Description: "",
			},
			"getAllEntertainmentTicket": {
				Type: graphql.NewList(types.GetEntertainmentTicketType()),
				Resolve: resolver.GetAllEntertainmentTicket,
				Description: "",
			},
			"getEntertainmentByTrending": {
				Type: graphql.NewList(types.GetEntertainmentType()),
				Args: graphql.FieldConfigArgument{
					"Type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.GetEntertainmentByTrending,
				Description: "",
			},
			"getHotelByGeolocation": {
				Type: graphql.NewList(types.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: resolver.GetHotelByGeolocation,
				Description: "",
			},
			"getPromoDetailById": {
				Type: types.GetPromoDetailType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetPromoDetailById,
				Description: "",
			},
			"getPromoById": {
				Type: types.GetPromoType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetPromoById,
				Description: "",
			},
			"getOtherPromo": {
				Type: graphql.NewList(types.GetPromoType()),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetOtherPromo,
				Description: "",
			},
			"getAllPromo": {
				Type: graphql.NewList(types.GetPromoType()),
				Resolve: resolver.GetAllPromo,
				Description: "",
			},
			"getAllPromoDetail": {
				Type: graphql.NewList(types.GetPromoDetailType()),
				Resolve: resolver.GetAllPromoDetail,
				Description: "",
			},
			"getAllBlog": {
				Type: graphql.NewList(types.GetBlogType()),
				Resolve: resolver.GetAllBlog,
				Description: "",
			},
			"getBlogById": {
				Type: types.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetBlogById,
				Description: "",
			},
			"getAllBlogExcept": {
				Type: graphql.NewList(types.GetBlogType()),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetAllBlogExcept,
				Description: "",
			},
			"sendMail": {
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"to": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: resolver.SendMail,
				Description: "",
			},
			"getHotelById": {
				Type: graphql.NewList(types.GetHotelType()),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetHotelById,
				Description: "",
			},
			"getAllChat": {
				Type: graphql.NewList(types.GetChatType()),
				Args: graphql.FieldConfigArgument{
					"UserId": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetAllChat,
				Description: "",
			},
			"getChatById": {
				Type: types.GetChatType(),
				Args: graphql.FieldConfigArgument{
					"Id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: resolver.GetChatById,
				Description: "",
			},
		},
		Description: "",
	})
}