package resolver

import (
	"github.com/danny2204/KarcisCoAPI/models"
	"github.com/graphql-go/graphql"
	"time"
)

//flight
func GetAllFlight(p graphql.ResolveParams)(interface{}, error) {
	flights, err := models.GetAllFlight()
	return flights, err
}

func CreateFlight(p graphql.ResolveParams)(interface{}, error) {
	destination := p.Args["destination"].(string)
	from := p.Args["from"].(string)
	flightDateFrom := p.Args["flightDateFrom"].(time.Time)
	flightDateTo := p.Args["flightDateTo"].(time.Time)
	airlineRefer := p.Args["airlineRefer"].(int)
	price := p.Args["price"].(int)

	flights, err := models.CreateFlight(destination, from, flightDateFrom, flightDateTo, airlineRefer, price)
	return flights, err
}

func RemoveFlight(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	flights, err := models.RemoveFlight(id)
	return flights, err
}

func GetFlightById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	flights, err := models.GetFlightById(id)
	return flights, err
}

func UpdateFlight(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	destination := p.Args["destination"].(string)
	from := p.Args["from"].(string)
	flightDateFrom := p.Args["flightDateFrom"].(time.Time)
	flightDateTo := p.Args["flightDateTo"].(time.Time)
	airlineRefer := p.Args["airlineRefer"].(int)
	price := p.Args["price"].(int64)

	flights, err := models.UpdateFlight(id, destination, from, flightDateFrom, flightDateTo, airlineRefer, price)
	return flights, err
}

//user
func GetAllUser(p graphql.ResolveParams)(interface{}, error) {
	users, err := models.GetAllUser()
	return users, err
}

func CreateUser(p graphql.ResolveParams)(interface{}, error) {
	frontName := p.Args["frontName"].(string)
	backName := p.Args["backName"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)
	phoneNumber := p.Args["phoneNumber"].(string)

	users, err := models.CreateUser(frontName, backName, email, password, phoneNumber)
	return users, err
}

func GetUserByEmailAndPhone(p graphql.ResolveParams)(interface{}, error) {
	arg := p.Args["arg"].(string)
	users, err := models.GetUserByEmailAndPhone(arg)
	return users, err
}

func RemoveUser(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	users, err := models.RemoveUser(id)
	return users, err
}

//admin
func GetAllAdmin(p graphql.ResolveParams)(interface {}, error) {
	admins, err := models.GetAllAdmin()
	return admins, err
}

func CreateAdmin(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["name"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	admins, err := models.CreateAdmin(name, email, password)
	return admins, err
}

func UpdateAdmin(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	admins, err := models.UpdateAdmin(id, name, email, password)
	return admins, err
}

func RemoveAdmin(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)

	admins, err := models.RemoveAdmin(id)
	return admins, err
}

func GetAdminById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)

	admins, err := models.GetAdminById(id)
	return admins, err
}

//airline
func GetAllAirline(p graphql.ResolveParams)(interface{}, error) {
	airlines, err := models.GetAllAirline()
	return airlines, err
}

func CreateAirline(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["name"].(string)
	code := p.Args["code"].(string)
	seat := p.Args["seat"].(int)

	airline, err := models.InsertAirline(name, code, seat)
	return airline, err
}

func UpdateAirline(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	code := p.Args["code"].(string)
	seat := p.Args["seat"].(int)

	admins, err := models.UpdateAirline(id, name, code, seat)
	return admins, err
}

func RemoveAirline(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)

	airlines, err := models.RemoveAirline(id)
	return airlines, err
}

func GetAirlineById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)

	airlines, err := models.GetAirlineById(id)
	return airlines, err
}
