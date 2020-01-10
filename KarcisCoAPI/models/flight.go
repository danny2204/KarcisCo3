package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Flight struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Destination string
	From string
	FlightDateFrom time.Time
	FlightDateTo time.Time
	Airline Airline `gorm:"foreignkey:AirlineRefer"`
	AirlineRefer int
	Price int
}

func init() {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Flight{}).AddForeignKey("airline_refer", "airlines(id)", "CASCADE", "CASCADE")
}

func GetAllFlight()([]Flight, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var flightList []Flight
	db.Find(&flightList)
	for i,_ := range flightList {
		db.Model(flightList[i]).Related(&flightList[i].Airline, "AirlineRefer")
	}

	return flightList, err
}

func CreateFlight(destination string, from string, flightDateFrom time.Time, flightDateTo time.Time, airlineID int, price int) (*Flight, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	flight := &Flight{
		Destination: destination,
		From: from,
		FlightDateFrom: flightDateFrom,
		FlightDateTo: flightDateTo,
		AirlineRefer: airlineID,
		Price: price,
	}

	db.Save(flight)
	return flight, err
}

func UpdateFlight(id int, destination string, from string, flightDateFrom time.Time, flightDateTo time.Time, airlineID int, price int64) (*Flight, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var flight Flight
	db.Model(&flight).Where("id = ?", id).Updates(map[string]interface{}{"destination": destination, "from": from, "flight_date_from": flightDateFrom, "flight_date_to": flightDateTo, "airline_refer": airlineID, "price": price})
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
