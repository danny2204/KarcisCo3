package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Facility struct {
	Id int `gorm:primary_key`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	FacilityName string
	FacilityPrice int `gorm:type:int`
	FlightId int
}

func GetAllFacility()([]Facility, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var facilityList []Facility
	db.Find(&facilityList)
	return facilityList, err
}

func CreateFacility(facilityName string, facilityPrice int, flightId int)(*Facility, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	facility := &Facility{
		FacilityName: facilityName,
		FacilityPrice: facilityPrice,
		FlightId: flightId,
	}
	db.Save(facility)
	return facility, err
}

func UpdateFacility(id int, facilityName string, facilityPrice int, flightId int) (*Facility, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var facility Facility
	db.Model(&facility).Where("id = ?", id).Updates(map[string]interface{}{"facility_name": facilityName, "facility_price": facilityPrice, "flight_id": flightId})
	db.Where("id = ?", id).Find(&facility)
	return &facility, nil
}

func RemoveFacility(id int) (*Facility, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	facility := Facility{Id: id}
	db.Where("id = ?", id).Find(&facility)
	return &facility, db.Delete(facility).Error
}