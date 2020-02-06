package models

import (
	"fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Flight struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	To Airport `gorm:"foreignkey:ToRefer"`
	ToRefer int
	From Airport `gorm:"foreignkey:FromRefer"`
	FromRefer int
	Airline Airline `gorm:"foreignkey:AirlineRefer"`
	AirlineRefer int
	Price int `gorm:"type:int"`
	Tax int `gorm:"type:int"`
	ServiceCharge int `gorm:type:int`
	Routes []Route `gorm:"foreingkey:FlightRouteId"`
	Facilities []Facility `gorm:"foreignkey:FlightId"`
	Departure time.Time
	Arrival time.Time
	Duration int
}

func GetAllFlight()([]Flight, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var flightList []Flight
	db.Find(&flightList)
	//db.Preload("Facilities").Find(&flightList)
	for i,_ := range flightList {
		db.Model(flightList[i]).Related(&flightList[i].Airline, "AirlineRefer")
		db.Model(flightList[i]).Related(&flightList[i].Facilities , "FlightId")
		db.Model(flightList[i]).Related(&flightList[i].Routes, "FlightRouteId")
		db.Model(flightList[i]).Related(&flightList[i].To, "ToRefer")
		db.Model(flightList[i]).Related(&flightList[i].From, "FromRefer")
		for j,_ := range flightList[i].Routes {
			db.Model(flightList[i].Routes[j]).Related(&flightList[i].Routes[j].AirportFrom, "RouteFromRefer").Related(&flightList[i].Routes[j].AirportTo, "RouteToRefer")
		}
	}
	fmt.Println(flightList)
	return flightList, err
}

func CreateFlight(torefer int, fromrefer int, airlinerefer int, price int, tax int, servicecharge int, /*route Route[],*/ departure time.Time, arrival time.Time, duration int) (*Flight, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	flight := &Flight{
		ToRefer: torefer,
		FromRefer: fromrefer,
		AirlineRefer: airlinerefer,
		Price: price,
		Tax: tax,
		ServiceCharge: servicecharge,
		Departure: departure,
		Arrival: arrival,
		Duration: duration,
	}
	fmt.Println(flight)
	db.Save(flight)
	return flight, err
}

func UpdateFlight(id int, torefer uint, fromrefer uint, airlinerefer int, price int, tax int, servicecharge int, /*route Route[],*/ departure time.Time, arrival time.Time, duration uint) (*Flight, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var flight Flight
	db.Model(&flight).Where("id = ?", id).Updates(map[string]interface{}{"to_refer": torefer, "from_refer": fromrefer, "airline_refer": airlinerefer, "price": price, "tax": tax, "service_charge": servicecharge, "departure": departure, "arrival": arrival, "duration": duration})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&flight)
	return &flight, nil
}

func RemoveFlight(id int) (*Flight, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	flight := Flight{ID: id}
	db.Where("id = ?", id).Find(&flight)
	return &flight, db.Delete(flight).Error
}

func GetFlightById(id int)(*Flight, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	flight := Flight{ID: id}
	db.Where("id = ?", id).Find(&flight)
	return &flight, err
}

func GetFlightByToAndFrom(to string, from string)([]Flight, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var flight []Flight
	db.Where("to_refer in (?) AND from_refer in(?)", db.Table("airports").Select("id").Where("city = ?", to).SubQuery(), db.Table("airports").Select("id").Where("city = ?", from).SubQuery()).Find(&flight)
	for i,_ := range flight {
		//db.Model(flight[i]).Related(&flight[i].To, "ToRefer").Related(&flight[i].From, "FromRefer").Related(&flight[i].Airline, "AirlineRefer")
		db.Model(flight[i]).Related(&flight[i].Airline, "AirlineRefer")
		db.Model(flight[i]).Related(&flight[i].Facilities , "FlightId")
		db.Model(flight[i]).Related(&flight[i].Routes, "FlightRouteId")
		db.Model(flight[i]).Related(&flight[i].To, "ToRefer")
		db.Model(flight[i]).Related(&flight[i].From, "FromRefer")
		for j,_ := range flight[i].Routes {
			db.Model(flight[i].Routes[j]).Related(&flight[i].Routes[j].AirportFrom, "RouteFromRefer").Related(&flight[i].Routes[j].AirportTo, "RouteToRefer")
		}
	}
	return flight, err
}
