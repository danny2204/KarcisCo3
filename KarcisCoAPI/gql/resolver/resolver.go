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
	torefer := p.Args["ToRefer"].(int)
	fromrefer := p.Args["FromRefer"].(int)
	airlinerefer := p.Args["AirlineRefer"].(int)
	price := p.Args["Price"].(int)
	tax := p.Args["Tax"].(int)
	servicecharge := p.Args["ServiceCharge"].(int)
	departure := p.Args["Departure"].(time.Time)
	arrival := p.Args["Arrival"].(time.Time)
	duration := p.Args["Duration"].(int)

	flights, err := models.CreateFlight(torefer, fromrefer, airlinerefer, price, tax, servicecharge, departure, arrival, duration)
	return flights, err
}

func RemoveFlight(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["ID"].(int)
	flights, err := models.RemoveFlight(id)
	return flights, err
}

func GetFlightById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["ID"].(int)
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
	id := p.Args["ID"].(int)
	torefer := p.Args["ToRefer"].(int)
	fromrefer := p.Args["FromRefer"].(int)
	airlinerefer := p.Args["AirlineRefer"].(int)
	price := p.Args["Price"].(int)
	tax := p.Args["Tax"].(int)
	servicecharge := p.Args["ServiceCharge"].(int)
	departure := p.Args["Departure"].(time.Time)
	arrival := p.Args["Arrival"].(time.Time)
	duration := p.Args["Duration"].(int)

	flights, err := models.UpdateFlight(id, torefer, fromrefer, airlinerefer, price, tax, servicecharge, departure, arrival, duration)
	return flights, err
}

//user
func GetAllUser(p graphql.ResolveParams)(interface{}, error) {
	users, err := models.GetAllUser()
	return users, err
}

func CreateUser(p graphql.ResolveParams)(interface{}, error) {
	frontName := p.Args["FrontName"].(string)
	backName := p.Args["BackName"].(string)
	email := p.Args["Email"].(string)
	password := p.Args["Password"].(string)
	phoneNumber := p.Args["PhoneNumber"].(string)
	title := p.Args["Title"].(string)
	city := p.Args["City"].(string)
	address := p.Args["Address"].(string)
	postcode := p.Args["PostCode"].(string)
	facebookId := p.Args["FacebookId"].(string)
	googleId := p.Args["GoogleId"].(string)
	gender := p.Args["Gender"].(string)
	mainLanguage := p.Args["MainLanguage"].(string)

	users, err := models.CreateUser(frontName, backName, email, password, phoneNumber, title, city, address, postcode,  googleId, facebookId, gender, mainLanguage)
	return users, err
}

func UpdateUser(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["ID"].(int)
	frontName := p.Args["FrontName"].(string)
	backName := p.Args["BackName"].(string)
	email := p.Args["Email"].(string)
	password := p.Args["Password"].(string)
	phoneNumber := p.Args["PhoneNumber"].(string)
	title := p.Args["Title"].(string)
	city := p.Args["City"].(string)
	address := p.Args["Address"].(string)
	postcode := p.Args["PostCode"].(string)
	facebookId := p.Args["FacebookId"].(string)
	googleId := p.Args["GoogleId"].(string)
	gender := p.Args["Gender"].(string)
	mainLanguage := p.Args["MainLanguage"].(string)

	user, err := models.UpdateUser(id, frontName, backName, email, password, phoneNumber, title, city, address, postcode, googleId, facebookId, gender, mainLanguage)
	return user, err

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
	name := p.Args["Name"].(string)

	airline, err := models.InsertAirline(name)
	return airline, err
}

func UpdateAirline(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["ID"].(int)
	name := p.Args["Name"].(string)

	admins, err := models.UpdateAirline(id, name)
	return admins, err
}

func RemoveAirline(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["ID"].(int)

	airlines, err := models.RemoveAirline(id)
	return airlines, err
}

func GetAirlineById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["ID"].(int)

	airlines, err := models.GetAirlineById(id)
	return airlines, err
}

//airport
func RemoveAirport(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	airport, err := models.RemoveAirport(id)
	return airport, err
}

func CreateAirport(p graphql.ResolveParams)(interface{}, error) {
	code := p.Args["Code"].(string)
	name := p.Args["Name"].(string)
	city := p.Args["City"].(string)
	citycode := p.Args["CityCode"].(string)
	province := p.Args["Province"].(string)
	country := p.Args["Country"].(string)

	airport, err := models.InsertAirport(code, name, city, citycode, province, country)
	return airport, err
}

func GetAllAirport(p graphql.ResolveParams)(interface{}, error) {
	airport, err := models.GetAllAirport()
	return airport, err
}

//facility
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

//route
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

//hotel
func GetAllHotel(p graphql.ResolveParams)(interface{}, error) {
	hotel, err := models.GetAllHotel()
	return hotel, err
}

func CreateHotel(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["Name"].(string)
	locationRefer := p.Args["LocationRefer"].(int)
	address := p.Args["Address"].(string)
	rating := p.Args["Rating"].(int)
	longitude := p.Args["Longitude"].(float64)
	latitude := p.Args["Latitude"].(float64)
	types := p.Args["Type"].(string)
	tripAdvisorReviewReferId := p.Args["TripAdvisorReviewReferId"].(int)
	tiketComReviewReferId := p.Args["TiketComReviewReferId"].(int)

	hotel, err := models.CreateHotel(name, locationRefer, address, rating, latitude, longitude, tripAdvisorReviewReferId, tiketComReviewReferId, types)
	return hotel, err
}

func UpdateHotel(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	name := p.Args["Name"].(string)
	locationRefer := p.Args["LocationRefer"].(int)
	address := p.Args["Address"].(string)
	rating := p.Args["Rating"].(int)
	longitude := p.Args["Longitude"].(float64)
	latitude := p.Args["Latitude"].(float64)
	types := p.Args["Type"].(string)
	tripAdvisorReviewReferId := p.Args["TripAdvisorReviewReferId"].(int)
	tiketComReviewReferId := p.Args["TiketComReviewReferId"].(int)

	hotel, err := models.UpdateHotel(id, name, locationRefer, address, rating, latitude, longitude, tripAdvisorReviewReferId, tiketComReviewReferId, types)
	return hotel, err
}

func RemoveHotel(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["id"].(int)
	hotel, err := models.RemoveHotel(id)
	return hotel, err
}

func GetHotelByGeolocation(p graphql.ResolveParams)(interface{}, error) {
	longitude := p.Args["longitude"].(float64)
	latitude := p.Args["latitude"].(float64)

	hotel, err := models.GetHotelByGeolocation(longitude, latitude)
	return hotel, err
}

//location
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

func GetLimitedLocation(p graphql.ResolveParams)(interface{}, error) {
	location, err := models.GetLimitedLocation()
	return location, err
}

func GetHotelByCity(p graphql.ResolveParams)(interface{}, error) {
	city := p.Args["city"].(string)
	cities, err := models.GetHotelByCity(city)
	return cities, err
}

//station
func GetAllStation(p graphql.ResolveParams)(interface{}, error) {
	station, err := models.GetAllStation()
	return station, err
}

func CreateStation(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["Name"].(string)
	code := p.Args["Code"].(string)
	locationReferId := p.Args["LocationReferId"].(int)
	station, err := models.CreateStation(code, name, locationReferId)
	return station, err
}

func UpdateStation(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	name := p.Args["Name"].(string)
	code := p.Args["Code"].(string)
	locationReferId := p.Args["LocationReferId"].(int)

	stations, err := models.UpdateStation(id, code, name, locationReferId)
	return stations, err
}

func RemoveStation(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	station, err := models.RemoveStation(id)
	return station, err
}

//trip
func GetAllTrip(p graphql.ResolveParams)(interface{}, error) {
	trips, error := models.GetAllTrip()
	return trips, error
}

func CreateTrip(p graphql.ResolveParams)(interface{}, error) {
	trainRefer := p.Args["TrainRefer"].(int)
	fromRefer := p.Args["FromRefer"].(int)
	toRefer := p.Args["ToRefer"].(int)
	departure := p.Args["Departure"].(time.Time)
	arrival := p.Args["Arrival"].(time.Time)
	duration := p.Args["Duration"].(int)
	price := p.Args["Price"].(int)
	tax := p.Args["Tax"].(int)
	service := p.Args["Service"].(int)

	trips, err := models.CreateTrip(trainRefer, fromRefer, toRefer,  departure, arrival, duration, price, tax, service)
	return trips, err
}

func UpdateTrip(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	trainRefer := p.Args["TrainRefer"].(int)
	fromRefer := p.Args["FromRefer"].(int)
	toRefer := p.Args["ToRefer"].(int)
	departure := p.Args["Departure"].(time.Time)
	arrival := p.Args["Arrival"].(time.Time)
	duration := p.Args["Duration"].(int)
	price := p.Args["Price"].(int)
	tax := p.Args["Tax"].(int)
	service := p.Args["Service"].(int)

	trips, err := models.UpdateTrip(id, trainRefer, fromRefer, toRefer, departure, arrival, duration, price, tax, service)
	return trips, err
}

func RemoveTrip(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	trip, err := models.RemoveTrip(id)
	return trip, err
}

func GetTripByToAndFrom(p graphql.ResolveParams)(interface{}, error) {
	to := p.Args["to"].(string)
	from := p.Args["from"].(string)
	trip, err := models.GetTripByToAndFrom(to, from)
	return trip, err
}

//train
func GetAllTrain(p graphql.ResolveParams)(interface{}, error) {
	trains, error := models.GetAllTrain()
	return trains, error
}

func CreateTrain(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["Name"].(string)
	class := p.Args["Class"].(string)
	subClass := p.Args["Subclass"].(string)
	code := p.Args["Code"].(string)

	trains, err := models.CreateTrain(name, class, subClass, code)
	return trains, err
}

func UpdateTrain(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	name := p.Args["Name"].(string)
	class := p.Args["Class"].(string)
	subClass := p.Args["Subclass"].(string)
	code := p.Args["Code"].(string)

	trains, err := models.UpdateTrain(id, name, class, subClass, code)
	return trains, err
}

func RemoveTrain(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	train, err := models.RemoveTrain(id)
	return train, err
}

//vendor location
func GetAllVendorLocation(p graphql.ResolveParams)(interface{}, error) {
	vendorLocation, err := models.GetAllVendorLocation()
	return vendorLocation, err
}

func CreateVendorLocation(p graphql.ResolveParams) (interface{}, error) {
	locationRefer := p.Args["LocationRefer"].(int)
	vendorLocationReferId := p.Args["VendorLocationReferId"].(int)

	vendorLocation, err := models.CreateVendorLocation(locationRefer, vendorLocationReferId)
	return vendorLocation, err
}

func UpdateVendorLocation(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	locationRefer := p.Args["LocationRefer"].(int)
	vendorLocationReferId := p.Args["VendorLocationReferId"].(int)

	vendorLocation, err := models.UpdateVendorLocation(id, vendorLocationReferId, locationRefer)
	return vendorLocation, err
}

func RemoveVendorLocation(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	vendorLocation, err := models.RemoveVendorLocation(id)
	return vendorLocation, err
}

//vendor
func GetAllVendor(p graphql.ResolveParams)(interface{}, error) {
	vendor, err := models.GetAllVendor()
	return vendor, err
}

func CreateVendor(p graphql.ResolveParams) (interface{}, error) {
	name := p.Args["Name"].(string)
	vendor, err := models.CreateVendor(name)
	return vendor, err
}

func UpdateVendor(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	name := p.Args["Name"].(string)

	vendor, err := models.UpdateVendor(id, name)
	return vendor, err
}

func RemoveVendor(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	vendor, err := models.RemoveVendor(id)
	return vendor, err
}

//car vendor
func GetAllCarVendor(p graphql.ResolveParams)(interface{}, error) {
	carVendor, err := models.GetAllCarVendor()
	return carVendor, err
}

func CreateCarVendor(p graphql.ResolveParams) (interface{}, error) {
	vendorReferId := p.Args["VendorReferId"].(int)
	carReferId := p.Args["CarReferId"].(int)
	price := p.Args["Price"].(int)

	carVendor, err := models.CreateCarVendor(vendorReferId, carReferId, price)
	return carVendor, err
}

func UpdateCarVendor(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	vendorReferId := p.Args["VendorReferId"].(int)
	carReferId := p.Args["CarreferId"].(int)
	price := p.Args["Price"].(int)

	carVendor, err := models.UpdateCarVendor(id, vendorReferId, carReferId, price)
	return carVendor, err
}

func RemoveCarVendor(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	carVendor, err := models.RemoveCarVendor(id)
	return carVendor, err
}

func GetAllCar(p graphql.ResolveParams)(interface{}, error) {
	car, err := models.GetAllCar()
	return car, err
}

func CreateCar(p graphql.ResolveParams) (interface{}, error) {
	brand := p.Args["Brand"].(string)
	model := p.Args["Model"].(string)
	passenger := p.Args["Passenger"].(int)
	luggage := p.Args["Luggage"].(int)
	availableAt := p.Args["AvailableAt"].(time.Time)

	car, err := models.CreateCar(brand, model, passenger, luggage, availableAt)
	return car, err
}

func UpdateCar(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	brand := p.Args["Brand"].(string)
	model := p.Args["Model"].(string)
	passenger := p.Args["Passenger"].(int)
	luggage := p.Args["Luggage"].(int)
	availableAt := p.Args["AvailableAt"].(time.Time)

	car, err := models.UpdateCar(id, brand, model, passenger, luggage, availableAt)
	return car, err
}

func RemoveCar(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	car, err := models.RemoveCar(id)
	return car, err
}

func GetCarByLocationAndFromAndTo(p graphql.ResolveParams)(interface{}, error) {
	location := p.Args["Location"].(string)
	from := p.Args["From"].(string)

	car, err := models.GetCarByLocationAndFromAndTo(location, from)
	return car, err
}

func GetAllRoom(p graphql.ResolveParams)(interface{}, error) {
	roomList, err := models.GetAllRoom()
	return roomList, err
}

func CreateRoom(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["Name"].(string)
	roomSize := p.Args["RoomSize"].(int)
	price := p.Args["Price"].(int)
	maxPassenger := p.Args["MaxPassenger"].(int)
	freeBreakfast := p.Args["FreeBreakfast"].(bool)
	freeWifi := p.Args["FreeWifi"].(bool)
	bed := p.Args["Bed"].(string)
	roomLeft := p.Args["RoomLeft"].(int)
	roomReferId := p.Args["RoomReferId"].(int)

	room, err := models.CreateRoom(name, roomSize, price, maxPassenger, freeBreakfast, freeWifi, bed, roomLeft, roomReferId)
	return room, err
}

func UpdateRoom(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	name := p.Args["Name"].(string)
	roomSize := p.Args["RoomSize"].(int)
	price := p.Args["Price"].(int)
	maxPassenger := p.Args["MaxPassenger"].(int)
	freeBreakfast := p.Args["FreeBreakfast"].(bool)
	freeWifi := p.Args["FreeWifi"].(bool)
	bed := p.Args["Bed"].(string)
	roomLeft := p.Args["RoomLeft"].(int)
	roomReferId := p.Args["RoomReferId"].(int)

	room, err := models.UpdateRoom(id, name, roomSize, price, maxPassenger, freeBreakfast, freeWifi, bed, roomLeft, roomReferId)
	return room, err
}

func RemoveRoom(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	room, err := models.RemoveRoom(id)
	return room, err
}

func GetAllReview(p graphql.ResolveParams)(interface{}, error) {
	reviews, err := models.GetAllReview()
	return reviews, err
}

func CreateReview(p graphql.ResolveParams)(interface{}, error) {
	score := p.Args["Score"].(float64)
	name := p.Args["Name"].(string)
	uploadDate := p.Args["UploadDate"].(time.Time)
	conclusion := p.Args["Conclusion"].(string)
	reviewString := p.Args["ReviewString"].(string)
	tiketComReviewReferId := p.Args["TiketComReviewReferId"].(int)
	tripAdvisorReviewReferId := p.Args["TripAdvisorReviewReferId"].(int)

	reviews, err := models.CreateReview(score, uploadDate, conclusion, reviewString, tiketComReviewReferId, tripAdvisorReviewReferId, name)
	return reviews, err
}

func UpdateReview(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	score := p.Args["Score"].(float64)
	name := p.Args["Name"].(string)
	uploadDate := p.Args["UploadDate"].(time.Time)
	conclusion := p.Args["Conclusion"].(string)
	reviewString := p.Args["ReviewString"].(string)
	tiketComReviewReferId := p.Args["TiketComReviewReferId"].(int)
	tripAdvisorReviewReferId := p.Args["TripAdvisorReviewReferId"].(int)

	reviews, err := models.UpdateReview(id, score, uploadDate, conclusion, reviewString, tiketComReviewReferId, tripAdvisorReviewReferId, name)
	return reviews, err
}

func RemoveReview(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	review, err := models.RemoveReview(id)
	return review, err
}

func GetAllTripAdvisorReview(p graphql.ResolveParams)(interface{}, error) {
	tripAdvisorReview, err := models.GetAllTripAdvisorReview()
	return tripAdvisorReview, err
}

func CreateTripAdvisorReview(p graphql.ResolveParams)(interface{}, error) {
	averageScore := p.Args["AverageScore"].(float64)
	locationScore := p.Args["LocationScore"].(float64)
	cleanlinessScore := p.Args["CleanlinessScore"].(float64)
	serviceScore := p.Args["ServiceScore"].(float64)
	valueScore := p.Args["ValueScore"].(float64)

	tripAdvisorReview, err := models.CreateTripAdvisorReview(averageScore, locationScore, cleanlinessScore, serviceScore, valueScore)
	return tripAdvisorReview, err
}

func UpdateTripAdvisorReview(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	averageScore := p.Args["AverageScore"].(float64)
	locationScore := p.Args["LocationScore"].(float64)
	cleanlinessScore := p.Args["CleanlinessScore"].(float64)
	serviceScore := p.Args["ServiceScore"].(float64)
	valueScore := p.Args["ValueScore"].(float64)

	tripAdvisorReview, err := models.UpdateTripAdvisorReview(id, averageScore, locationScore, cleanlinessScore, serviceScore, valueScore)
	return tripAdvisorReview, err
}

func RemoveTripAdvisorReview(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	tripAdvisorReview, err := models.RemoveTripAdvisorReview(id)
	return tripAdvisorReview, err
}

func GetAllTiketComReview(p graphql.ResolveParams)(interface{}, error) {
	tiketComReview, err := models.GetAllTiketComReview()
	return tiketComReview, err
}

func CreateTiketComReview(p graphql.ResolveParams)(interface{}, error) {
	averageScore := p.Args["AverageScore"].(float64)
	locationScore := p.Args["LocationScore"].(float64)
	cleanlinessScore := p.Args["CleanlinessScore"].(float64)
	serviceScore := p.Args["ServiceScore"].(float64)
	valueScore := p.Args["ValueScore"].(float64)

	tikerComReview, err := models.CreateTiketComReview(averageScore, locationScore, cleanlinessScore, serviceScore, valueScore)
	return tikerComReview, err
}

func UpdateTiketComReview(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	averageScore := p.Args["AverageScore"].(float64)
	locationScore := p.Args["LocationScore"].(float64)
	cleanlinessScore := p.Args["CleanlinessScore"].(float64)
	serviceScore := p.Args["ServiceScore"].(float64)
	valueScore := p.Args["ValueScore"].(float64)

	tikerComReview, err := models.UpdateTiketComReview(id, averageScore, locationScore, cleanlinessScore, serviceScore, valueScore)
	return tikerComReview, err
}

func RemoveTiketComReview(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	tiketComReview, err := models.RemoveTiketComReview(id)
	return tiketComReview, err
}

//entertainment
func GetAllEntertainment(p graphql.ResolveParams)(interface{}, error) {
	entertainment, err := models.GetAllEntertainment()
	return entertainment, err
}

func CreateEntertainment(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["Name"].(string)
	locationReferId := p.Args["LocationReferId"].(int)
	types := p.Args["Type"].(string)
	eventStartDate := p.Args["EventStartDate"].(time.Time)
	eventEndDate := p.Args["EventEndDate"].(time.Time)
	isTrending := p.Args["IsTrending"].(bool)
	description := p.Args["Description"].(string)

	entertainment, err := models.CreateEntertainment(name, locationReferId, types, eventStartDate, eventEndDate, isTrending, description)
	return entertainment, err
}

func UpdateEntertainment(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	name := p.Args["Name"].(string)
	locationReferId := p.Args["LocationReferId"].(int)
	types := p.Args["Type"].(string)
	eventStartDate := p.Args["EventStartDate"].(time.Time)
	eventEndDate := p.Args["EventEndDate"].(time.Time)
	isTrending := p.Args["IsTrending"].(bool)
	description := p.Args["Description"].(string)

	entertainment, err := models.UpdateEntertainment(id, name, locationReferId, types, eventStartDate, eventEndDate, isTrending, description)
	return entertainment, err
}

func RemoveEntertainment(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	entertainment, err := models.RemoveEntertainment(id)
	return entertainment, err
}

func GetEntertainmentByLocation(p graphql.ResolveParams)(interface{}, error) {
	location := p.Args["Location"].(string)
	entertainment, err := models.GetEntertainmentByLocation(location)
	return entertainment, err
}

func GetEntertainmentByType(p graphql.ResolveParams)(interface{}, error) {
	types := p.Args["Type"].(string)
	entertainment, err := models.GetEntertainmentByType(types)
	return entertainment, err
}

func GetEntertainmentByTrending(p graphql.ResolveParams)(interface{}, error) {
	types := p.Args["Type"].(string)
	entertainment, err := models.GetEntertainmentByTrending(types)
	return entertainment, err
}

func GetEntertainmentById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	entertainment, err := models.GetEntertainmentById(id)
	return entertainment, err
}

//entertainment ticket
func GetAllEntertainmentTicket(p graphql.ResolveParams)(interface{}, error) {
	entertainmentTicket, err := models.GetAllEntertainmentTicket()
	return entertainmentTicket, err
}

func CreateEntertainmentTicket(p graphql.ResolveParams)(interface{}, error) {
	types := p.Args["Type"].(string)
	price := p.Args["Price"].(int)
	description := p.Args["Description"].(string)
	entertainmentReferId := p.Args["EntertainmentReferId"].(int)

	entertainmentTicket, err := models.CreateEntertainmentTicket(types, price, entertainmentReferId, description)
	return entertainmentTicket, err
}

func UpdateEntertainmentTicket(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	types := p.Args["Type"].(string)
	price := p.Args["Price"].(int)
	description := p.Args["Description"].(string)
	entertainmentReferId := p.Args["EntertainmentReferId"].(int)

	entertainmentTicket, err := models.UpdateEntertainmentTicket(id, types, price, entertainmentReferId, description)
	return entertainmentTicket, err
}

func RemoveEntertainmentTicket(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	entertainmentTicket, err := models.RemoveEntertainmentTicket(id)
	return entertainmentTicket, err
}

//promo
func GetAllPromo(p graphql.ResolveParams)(interface{}, error) {
	promo, err := models.GetAllPromo()
	return promo, err
}

func CreatePromo(p graphql.ResolveParams)(interface{}, error) {
	name := p.Args["Name"].(string)
	promoStart := p.Args["PromoStart"].(time.Time)
	promoEnd := p.Args["PromoEnd"].(time.Time)
	platform := p.Args["Platform"].(string)
	promoDetail := p.Args["PromoDetail"].(string)

	promo, err := models.CreatePromo(name, promoStart, promoEnd, platform, promoDetail)
	return promo, err
}

func GetPromoById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	promo, err := models.GetPromoById(id)
	return promo, err
}

func GetOtherPromo(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	promo, err := models.GetOtherPromo(id)
	return promo, err
}

//promo detail
func GetAllPromoDetail(p graphql.ResolveParams)(interface{}, error) {
	promoDetial, err := models.GetAllPromoDetail()
	return promoDetial, err
}

func CreatePromoDetail(p graphql.ResolveParams)(interface{}, error) {
	discountAmount := p.Args["DiscountAmount"].(int)
	kodePromo := p.Args["KodePromo"].(string)
	promoReferId := p.Args["PromoReferId"].(int)

	promoDetail, err := models.CreatePromoDetail(discountAmount, kodePromo, promoReferId)
	return promoDetail, err
}

func GetPromoDetailById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	promoDetail, err := models.GetPromoDetailById(id)
	return promoDetail, err
}

//blog
func GetAllBlog(p graphql.ResolveParams)(interface{}, error) {
	blog, err := models.GetAllBlog()
	return blog, err
}

func CreateBlog(p graphql.ResolveParams)(interface{}, error) {
	blogTitle := p.Args["BlogTitle"].(string)
	blogDetail := p.Args["BlogDetail"].(string)
	view := p.Args["View"].(int)
	authorName := p.Args["AuthorName"].(string)
	category := p.Args["Category"].(string)
	imagePath := p.Args["ImagePath"].(string)

	blog, err := models.CreateBlog(blogTitle, blogDetail, view, imagePath, authorName, category)
	return blog, err
}

func UpdateBlog(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	blogTitle := p.Args["BlogTitle"].(string)
	blogDetail := p.Args["BlogDetail"].(string)
	view := p.Args["View"].(int)
	authorName := p.Args["AuthorName"].(string)
	category := p.Args["Category"].(string)
	imagePath := p.Args["ImagePath"].(string)

	blog, err := models.UpdateBlog(id, blogTitle, blogDetail, view, imagePath, authorName, category)
	return blog, err
}

func RemoveBlog(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	blog, err := models.RemoveBlog(id)
	return blog, err
}

func GetBlogById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	blog, err := models.GetBlogById(id)
	return blog, err
}

func GetAllBlogExcept(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	blogs, err := models.GetAllBlogExcept(id)
	return blogs, err
}


//send mail
func SendMail(p graphql.ResolveParams)(interface{}, error) {
	email := p.Args["to"].(string)
	models.Send(email)
	return nil, nil
}

func GetHotelById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	hotel, err := models.GetHotelById(id)
	return hotel, err
}

func GetAllChat(p graphql.ResolveParams)(interface{}, error) {

	userId := p.Args["UserId"].(int)

	chat, err := models.GetAllChat(userId)
	return chat, err
}

func CreateChat(p graphql.ResolveParams)(interface{}, error) {
	user1 := p.Args["User1"].(int)
	user2 := p.Args["User2"].(int)
	content := p.Args["Content"].(string)

	chat, err := models.CreateChat(user1, user2, content)
	return chat, err
}

func UpdateChat(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)
	user1 := p.Args["User1"].(int)
	user2 := p.Args["User2"].(int)
	content := p.Args["Content"].(string)

	chat, err := models.UpdateChat(id, user1, user2, content)
	return chat, err
}

func GetChatById(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["Id"].(int)

	chat, err := models.GetChatById(id)
	return chat, err
}

func GetAdminByEmail(p graphql.ResolveParams)(interface{}, error) {
	email := p.Args["Email"].(string)

	admin, err := models.GetAdminByEmail(email)
	return admin, err
}

func GetUserByAuthId(p graphql.ResolveParams)(interface{}, error) {
	id := p.Args["AuthId"].(string)
	users, err := models.GetUserByAuthId(id)
	return users, err
}
