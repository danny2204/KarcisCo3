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
	torefer := p.Args["toRefer"].(int)
	fromrefer := p.Args["fromRefer"].(int)
	airlinerefer := p.Args["airlineRefer"].(int)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	servicecharge := p.Args["serviceCharge"].(int)
	departure := p.Args["departure"].(time.Time)
	arrival := p.Args["arrival"].(time.Time)
	duration := p.Args["duration"].(int)

	flights, err := models.CreateFlight(torefer, fromrefer, airlinerefer, price, tax, servicecharge, departure, arrival, duration)
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

func GetFlightByToAndFrom(p graphql.ResolveParams)(interface{}, error) {
	to := p.Args["to"].(string)
	from := p.Args["from"].(string)
	flights, err := models.GetFlightByToAndFrom(to, from)
	return flights, err
}

func UpdateFlight(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	torefer := p.Args["toRefer"].(uint)
	fromrefer := p.Args["fromRefer"].(uint)
	airlinerefer := p.Args["airlineRefer"].(int)
	price := p.Args["price"].(int)
	tax := p.Args["tax"].(int)
	servicecharge := p.Args["serviceCharge"].(int)
	departure := p.Args["departure"].(time.Time)
	arrival := p.Args["arrival"].(time.Time)
	duration := p.Args["duration"].(uint)

	flights, err := models.UpdateFlight(id, torefer, fromrefer, airlinerefer, price, tax, servicecharge, departure, arrival, duration)
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
	title := p.Args["title"].(string)
	city := p.Args["city"].(string)
	address := p.Args["address"].(string)
	postcode := p.Args["postCode"].(string)

	users, err := models.CreateUser(frontName, backName, email, password, phoneNumber, title, city, address, postcode)
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

	airline, err := models.InsertAirline(name)
	return airline, err
}

func UpdateAirline(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)

	admins, err := models.UpdateAirline(id, name)
	return admins, err
}

func RemoveAirport(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	airport, err := models.RemoveAirport(id)
	return airport, err
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

func CreateAirport(p graphql.ResolveParams)(interface{}, error) {
	code := p.Args["code"].(string)
	name := p.Args["name"].(string)
	city := p.Args["city"].(string)
	citycode := p.Args["cityCode"].(string)
	province := p.Args["province"].(string)
	country := p.Args["country"].(string)

	airport, err := models.InsertAirport(code, name, city, citycode, province, country)
	return airport, err
}

func GetAllAirport(p graphql.ResolveParams)(interface{}, error) {
	airport, err := models.GetAllAirport()
	return airport, err
}

func GetAllFacility(p graphql.ResolveParams)(interface{}, error) {
	facility, err := models.GetAllFacility()
	return facility, err
}

func GetAllRoute(p graphql.ResolveParams)(interface{}, error) {
	Route, err := models.GetAllRoute()
	return Route, err
}

func CreateFacility(p graphql.ResolveParams)(interface{}, error) {
	facilityName := p.Args["facilityName"].(string)
	facilityPrice := p.Args["facilityPrice"].(int)
	flightId := p.Args["flightId"].(int)

	facility, err := models.CreateFacility(facilityName, facilityPrice, flightId)
	return facility, err
}

func CreateRoute(p graphql.ResolveParams)(interface{}, error) {
	toRefer := p.Args["routeToRefer"].(int)
	fromRefer := p.Args["routeFromRefer"].(int)
	duration := p.Args["duration"].(int)
	flightId := p.Args["flightRouteId"].(int)

	route, err := models.CreateRoute(toRefer, fromRefer, duration, flightId)
	return route, err
}

func UpdateFacility(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	facilityName := p.Args["facilityName"].(string)
	facilityPrice := p.Args["facilityPrice"].(int)
	flightId := p.Args["flightId"].(int)

	facility, err := models.UpdateFacility(id, facilityName, facilityPrice, flightId)
	return facility, err
}

func UpdateRoute(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	toRefer := p.Args["routeToRefer"].(int)
	fromRefer := p.Args["routeFromRefer"].(int)
	duration := p.Args["duration"].(int)
	flightId := p.Args["flightId"].(int)

	route, err := models.UpdateRoute(id, toRefer, fromRefer, duration, flightId)
	return route, err
}

func RemoveFacility(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)

	facility, err := models.RemoveFacility(id)
	return facility, err
}

func RemoveRoute(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	route, err := models.RemoveRoute(id)
	return route, err
}

func CreateHotelFaclity(p graphql.ResolveParams)(interface{}, error) {
	facilityName := p.Args["facilityName"].(string)
	facilityPrice := p.Args["facilityPrice"].(int)
	hotelPreferId := p.Args["hotelPreferId"].(int)

	hotelFacility, err := models.CreateHotelFacility(facilityName, facilityPrice, hotelPreferId)
	return hotelFacility, err
}

func UpdateHotelFacility(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	facilityName := p.Args["facilityName"].(string)
	facilityPrice := p.Args["facilityPrice"].(int)
	hotelPreferId := p.Args["hotelPreferId"].(int)

	hotelFacility, err := models.UpdateHotelFacility(id, facilityName, facilityPrice, hotelPreferId)
	return hotelFacility, err
}

func RemoveHotelFacility(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)

	hotelFacility, err := models.RemoveHotelFacility(id)
	return hotelFacility, err
}

func GetAllHotelFacility(p graphql.ResolveParams)(interface{}, error) {
	hotelFacility, err := models.GetAllHotelFacility()
	return hotelFacility, err
}

func GetAllHotel(p graphql.ResolveParams)(interface{}, error) {
	hotel, err := models.GetAllHotel()
	return hotel, err
}

func CreateHotel(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["name"].(string)
	locationRefer := p.Args["locationRefer"].(int)
	address := p.Args["address"].(string)
	rating := p.Args["rating"].(int)
	imagePath := p.Args["imagePath"].(string)

	hotel, err := models.CreateHotel(name, locationRefer, address, rating, imagePath)
	return hotel, err
}

func UpdateHotel(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	name := p.Args["name"].(string)
	locationRefer := p.Args["locationRefer"].(int)
	address := p.Args["address"].(string)
	rating := p.Args["rating"].(int)
	imagePath := p.Args["imagePath"].(string)

	hotel, err := models.UpdateHotel(id, name, locationRefer, address, rating, imagePath)
	return hotel, err
}

func RemoveHotel(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	hotel, err := models.RemoveHotel(id)
	return hotel, err
}

func GetAllLocation(p graphql.ResolveParams)(interface{}, error) {
	location, err := models.GetAllLocation()
	return location, err
}

func CreateLocation(p graphql.ResolveParams)(interface{}, error) {
	longitude := p.Args["longitude"].(float64)
	latitude := p.Args["latitude"].(float64)
	city := p.Args["city"].(string)
	province := p.Args["province"].(string)
	country := p.Args["country"].(string)
	zindex := p.Args["zindex"].(int)

	location, err := models.CreateLocation(longitude, latitude, city, province, country, zindex)
	return location, err
}

func UpdateLocation(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	longitude := p.Args["longitude"].(float64)
	latitude := p.Args["latitude"].(float64)
	city := p.Args["city"].(string)
	province := p.Args["province"].(string)
	country := p.Args["country"].(string)
	zindex := p.Args["zindex"].(int)

	location, err := models.UpdateLocation(id, longitude, latitude, city, province, country, zindex)
	return location, err
}

func RemoveLocation(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	location, err := models.RemoveLocation(id)
	return location, err
}

func GetHotelByCity(p graphql.ResolveParams)(interface{}, error) {
	city := p.Args["city"].(string)
	cities, err := models.GetHotelByCity(city)
	return cities, err
}