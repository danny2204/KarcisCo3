package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Car struct {
	Id			int				`gorm:primary_key`
	Brand		string			`gorm:"type:varchar(100);not null"`
	Model		string			`gorm:"type:varchar(100);not null"`
	CarVendor []CarVendor `gorm:"foreignkey:CarReferId"`
	Passenger	int 			`gorm:int`
	Luggage		int				`gorm:int`
	AvailableAt time.Time
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func GetAllCar()([]Car, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var carList []Car
	db.Find(&carList)
	for i,_ := range carList {
		db.Model(carList[i]).Related(&carList[i].CarVendor, "CarReferId")
		for j,_ := range carList[i].CarVendor {
			db.Model(carList[i].CarVendor[j]).Related(&carList[i].CarVendor[j].Vendor, "VendorReferId")
			db.Model(carList[i].CarVendor[j].Vendor).Related(&carList[i].CarVendor[j].Vendor.VendorLocation, "VendorLocationReferId")
			for k,_ := range carList[i].CarVendor[j].Vendor.VendorLocation {
				db.Model(carList[i].CarVendor[j].Vendor.VendorLocation[k]).Related(&carList[i].CarVendor[j].Vendor.VendorLocation[k].Location, "LocationRefer")
			}
		}
	}

	return carList, err
}

func CreateCar(brand string, model string, passenger int, luggage int, availableAt time.Time) (*Car, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	car := &Car{
		Brand: brand,
		Model: model,
		Passenger: passenger,
		Luggage: luggage,
		AvailableAt: availableAt,
	}

	db.Save(car)
	return car, err
}

func UpdateCar(id int, brand string, model string, passenger int, luggage int, availableAt time.Time) (*Car, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var car Car
	db.Model(&car).Where("id = ?", id).Updates(map[string]interface{}{"brand": brand, "model": model, "passenger": passenger, "luggage": luggage, "available_at": availableAt})
	db.Where("id = ?", id).Find(&car)
	return &car, nil
}

func RemoveCar(id int) (*Car, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	car := Car{Id: id}
	db.Where("id = ?", id).Find(&car)
	return &car, db.Delete(car).Error
}

func GetCarByLocationAndFromAndTo(location string, from string) ([]Car, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var carList []Car
		db.Where("id in (?) AND available_at BETWEEN ? AND ?", db.Table("car_vendors").Select("car_refer_id").Where("vendor_refer_id in (?)", db.Table("vendors").Select("id").Where("id in (?)", db.Table("vendor_locations").Select("vendor_location_refer_id").Where("location_refer in (?)", db.Table("locations").Select("id").Where("city = ?", location).SubQuery()).SubQuery()).SubQuery()).SubQuery(), from + " 00:00:00", from + " 23:59:59").Find(&carList)
	for i,_ := range carList {
		db.Model(carList[i]).Related(&carList[i].CarVendor, "CarReferId")
		for j,_ := range carList[i].CarVendor {
			db.Model(carList[i].CarVendor[j]).Related(&carList[i].CarVendor[j].Vendor, "VendorReferId")
			db.Model(carList[i].CarVendor[j].Vendor).Related(&carList[i].CarVendor[j].Vendor.VendorLocation, "VendorLocationReferId")
			for k,_ := range carList[i].CarVendor[j].Vendor.VendorLocation {
				db.Model(carList[i].CarVendor[j].Vendor.VendorLocation[k]).Related(&carList[i].CarVendor[j].Vendor.VendorLocation[k].Location, "LocationRefer")
			}
		}
	}
	return carList, err
}