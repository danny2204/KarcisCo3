package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type HotelFacility struct {
	Id int `gorm:primary_key`
	FacilityName string
	FacilityPrice int
	HotelReferId int
	CreatedAt time.Time
	UpdateAt time.Time
	DeletedAt *time.Time `sql:index`
}

func GetAllHotelFacility()([]HotelFacility, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var hotelFacilityList []HotelFacility
	db.Find(&hotelFacilityList)
	return hotelFacilityList, err
}

func CreateHotelFacility(facilityName string, facilityPrice int, hotelPreferId int)(*HotelFacility, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	hotelFacility := &HotelFacility{
		FacilityName: facilityName,
		FacilityPrice: facilityPrice,
		HotelReferId: hotelPreferId,
	}

	db.Save(hotelFacility)
	return hotelFacility, err
}

func UpdateHotelFacility(id int, facilityName string, facilityPrice int, hotelPreferId int) (*HotelFacility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotelFacility HotelFacility
	db.Model(&hotelFacility).Where("id = ?", id).Updates(map[string]interface{}{"facility_name": facilityName, "facility_price": facilityPrice, "hotel_prefer_id": hotelPreferId})
	db.Where("id = ?", id).Find(&hotelFacility)
	return &hotelFacility, nil
}

func RemoveHotelFacility(id int) (*HotelFacility, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	hotelFacility := HotelFacility{Id: id}
	db.Where("id = ?", id).Find(&hotelFacility)
	return &hotelFacility, db.Delete(hotelFacility).Error
}