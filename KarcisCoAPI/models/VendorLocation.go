package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type VendorLocation struct {
	Id int `gorm:"primary_key"`
	Location Location `gorm:"foreignkey: LocationRefer"`
	LocationRefer int
	VendorLocationReferId int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func GetAllVendorLocation()([]VendorLocation, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var vendorLocationList []VendorLocation
	db.Find(&vendorLocationList)
	for i,_ := range vendorLocationList {
		db.Model(vendorLocationList[i]).Related(vendorLocationList[i].Location, "LocationRefer")
	}
	return vendorLocationList, err
}

func CreateVendorLocation(locationRefer int, vendorLocationReferId int) (*VendorLocation, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	vendorLocation := &VendorLocation{
		LocationRefer: locationRefer,
		VendorLocationReferId: vendorLocationReferId,
	}

	db.Save(vendorLocation)
	return vendorLocation, err
}

func UpdateVendorLocation(id int, vendorLocationReferId int, locationRefer int) (*VendorLocation, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var vendorLocation VendorLocation
	db.Model(&vendorLocation).Where("id = ?", id).Updates(map[string]interface{}{"vendor_location_refer_id": vendorLocationReferId, "location_refer": locationRefer})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&vendorLocation)
	return &vendorLocation, nil
}

func RemoveVendorLocation(id int) (*VendorLocation, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	vendorLocation := VendorLocation{Id: id}
	db.Where("id = ?", id).Find(&vendorLocation)
	return &vendorLocation, db.Delete(vendorLocation).Error
}