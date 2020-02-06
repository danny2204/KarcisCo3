package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Trip struct{
	Id					int 		`gorm:primary_key`
	Train				Train		`gorm:"foreignkey:TrainRefer"`
	TrainRefer  		uint
	From				Station		`gorm:"foreignkey:FromRefer"`
	FromRefer			uint
	To					Station		`gorm:"foreignkey:ToRefer"`
	ToRefer				uint
	Departure			time.Time
	Arrival				time.Time
	Duration			uint
	Price				uint
	Tax					uint
	Service				uint

	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			*time.Time	`sql:index`
}

func GetAllTrip()([]Trip, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var tripList []Trip
	db.Find(&tripList)
	//for i,_ := range hotelList {
	//	db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
	//}
	return tripList, err
}

func CreateTrip(trainrefer uint, fromrefer uint, torefer uint, departure time.Time, arrival time.Time, duration uint, price uint, tax uint, service uint) (*Trip, error){
	db, err = connection.ConnectDatabase()
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

func UpdateTrip(id int, trainrefer uint, fromrefer uint, torefer uint, departure time.Time, arrival time.Time, duration uint, price uint, tax uint, service uint) (*Trip, error) {
	db, err := connection.ConnectDatabase()

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

	if err != nil {
		return nil, err
	}
	defer db.Close()

	trip := Trip{Id: id}
	db.Where("id = ?", id).Find(&trip)
	return &trip, err
}