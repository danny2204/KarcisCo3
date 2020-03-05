package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
)

type CarVendor struct {
	Id int `gorm:"primary_key"`
	Vendor Vendor `gorm:"foreignkey:VendorReferId"`
	VendorReferId int
	CarReferId int
	Price int
}

func GetAllCarVendor()([]CarVendor, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var carVendorList []CarVendor
	db.Find(&carVendorList)
	for i,_ := range carVendorList {
		db.Model(carVendorList[i]).Related(&carVendorList[i].Vendor, "VendorReferId")
		db.Model(carVendorList[i].Vendor).Related(&carVendorList[i].Vendor.VendorLocation, "VendorLocationReferId")
		for j,_ := range carVendorList[i].Vendor.VendorLocation {
			db.Model(carVendorList[i].Vendor.VendorLocation[j]).Related(&carVendorList[i].Vendor.VendorLocation[j].Location, "LocationRefer")
		}
	}
	return carVendorList, err
}

func CreateCarVendor(vendorReferId int, carReferId int, price int) (*CarVendor, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	carVendor := &CarVendor{
		VendorReferId: vendorReferId,
		CarReferId: carReferId,
		Price: price,
	}
	db.Save(carVendor)
	return carVendor, err
}

func UpdateCarVendor(id int, vendorReferId int, carReferId int, price int) (*CarVendor, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var carVendor CarVendor
	db.Model(&carVendor).Where("id = ?", id).Updates(map[string]interface{}{"vendor_refer_id": vendorReferId, "car_refer_id": carReferId, "price": price})
	db.Where("id = ?", id).Find(&carVendor)
	return &carVendor, nil
}

func RemoveCarVendor(id int) (*CarVendor, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	carVendor := CarVendor{Id: id}
	db.Where("id = ?", id).Find(&carVendor)
	return &carVendor, db.Delete(carVendor).Error
}