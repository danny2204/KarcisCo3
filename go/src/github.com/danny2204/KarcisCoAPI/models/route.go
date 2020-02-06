package models

import (
	"fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Route struct {
	Id int `gorm:primary_key`
	AirportFrom Airport `gorm:"foreignkey:RouteFromRefer"`
	RouteFromRefer int
	AirportTo Airport `gorm:"foreignkey:RouteToRefer"`
	RouteToRefer int
	Duration int
	FlightRouteId int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:index`
}

func GetAllRoute()([]Route, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var routeList []Route
	db.Find(&routeList)
	for i,_ := range routeList {
		db.Model(routeList[i]).Related(&routeList[i].AirportTo, "RouteToRefer").Related(&routeList[i].AirportFrom, "RouteFromRefer")
	}
	fmt.Println(routeList)
	return routeList, err
}

func CreateRoute(toRefer int, fromRefer int, duration int, flightId int) (*Route, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	route := &Route{
		RouteToRefer: toRefer,
		RouteFromRefer: fromRefer,
		Duration: duration,
		FlightRouteId: flightId,
	}

	db.Save(route)
	return route, err
}

func UpdateRoute(id int, toRefer int, fromRefer int, duration int, flightId int) (*Route, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var route Route
	db.Model(&route).Where("id = ?", id).Updates(map[string]interface{}{"route_from_refer": fromRefer, "route_to_refer": toRefer, "duration": duration, "flight_id": flightId})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&route)
	return &route, nil
}

func RemoveRoute(id int) (*Route, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	route := Route{Id: id}
	db.Where("id = ?", id).Find(&route)
	return &route, db.Delete(route).Error
}

func GetRouteById(id int)(*Route, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	route := Route{Id: id}
	db.Where("id = ?", id).Find(&route)
	return &route, err
}