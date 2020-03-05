package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Trip struct{
	Id					int 		`gorm:primary_key`
	Train				Train		`gorm:"foreignkey:TrainRefer"`
	TrainRefer  		int
	From				Station		`gorm:"foreignkey:FromRefer"`
	FromRefer			int
	To					Station		`gorm:"foreignkey:ToRefer"`
	ToRefer				int
	Departure			time.Time
	Arrival				time.Time
	Duration			int
	Price				int
	Tax					int
	Service				int

	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			*time.Time	`sql:index`
}

func GetAllTrip()([]Trip, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var tripList []Trip
	db.Find(&tripList)
	for i,_ := range tripList {
		db.Model(tripList[i]).Related(&tripList[i].Train, "TrainRefer")
		db.Model(tripList[i]).Related(&tripList[i].To, "ToRefer")
		db.Model(tripList[i]).Related(&tripList[i].From, "FromRefer")
		db.Model(tripList[i].To).Related(&tripList[i].To.Location, "LocationReferId")
		db.Model(tripList[i].From).Related(&tripList[i].From.Location, "LocationReferId")
	}
	return tripList, err
}

func CreateTrip(trainrefer int, fromrefer int, torefer int, departure time.Time, arrival time.Time, duration int, price int, tax int, service int) (*Trip, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	trip := &Trip{
		TrainRefer: trainrefer,
		FromRefer: fromrefer,
		ToRefer: torefer,
		Departure: departure,
		Arrival: arrival,
		Duration: duration,
		Price: price,
		Tax: tax,
		Service: service,
	}

	db.Save(trip)
	return trip, err
}

func UpdateTrip(id int, trainrefer int, fromrefer int, torefer int, departure time.Time, arrival time.Time, duration int, price int, tax int, service int) (*Trip, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var trip Trip
	db.Model(&trip).Where("id = ?", id).Updates(map[string]interface{}{"train_refer": trainrefer, "from_refer": fromrefer, "to+refer": torefer, "departure": departure, "arrival": arrival, "duration": duration, "price": price, "tax": tax, "service": service})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&trip)
	return &trip, nil
}

func RemoveTrip(id int) (*Trip, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	trip := Trip{Id: id}
	db.Where("id = ?", id).Find(&trip)
	return &trip, db.Delete(trip).Error
}

func GetTripById(id int)(*Trip, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	trip := Trip{Id: id}
	db.Where("id = ?", id).Find(&trip)
	return &trip, err
}

func GetTripByToAndFrom(to string, from string)([]Trip, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var trips []Trip
	db.Where("to_refer in (?) AND from_refer in (?)", db.Table("stations").Select("id").Where("id in (?)", db.Table("locations").Select("id").Where("city = ?", from).SubQuery()).SubQuery(), db.Table("stations").Select("id").Where("id in (?)", db.Table("locations").Select("id").Where("city = ?", to).SubQuery()).SubQuery()).Find(&trips)
	for i,_ := range trips {
		db.Model(trips[i]).Related(&trips[i].From, "FromRefer")
		db.Model(trips[i]).Related(&trips[i].To, "ToRefer")
		db.Model(trips[i]).Related(&trips[i].Train, "TrainRefer")

		db.Model(trips[i].From).Related(&trips[i].From.Location, "LocationReferId")
		db.Model(trips[i].To).Related(&trips[i].To.Location, "LocationReferId")
	}
	return trips, err
}