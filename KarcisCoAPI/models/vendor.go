package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Vendor struct{
	Id			int			`gorm:primary_key`
	VendorLocation []VendorLocation `gorm:"foreignkey:VendorLocationReferId"`
	Name		string		`gorm:"type:varchar(100)"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func GetAllVendor()([]Vendor, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var vendorList []Vendor
	db.Find(&vendorList)
	for i,_ := range vendorList {
		db.Model(vendorList[i]).Related(vendorList[i].VendorLocation, "VendorLocationReferId")
		for j,_ := range vendorList[i].VendorLocation {
			db.Model(vendorList[i].VendorLocation[j]).Related(vendorList[i].VendorLocation[j].Location, "LocationRefer")
		}
	}
	return vendorList, err
}

func CreateVendor(name string) (*Vendor, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	vendor := &Vendor {
		Name: name,
	}

	db.Save(vendor)
	return vendor, err
}

func UpdateVendor(id int, name string) (*Vendor, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var vendor Vendor
	db.Model(&vendor).Where("id = ?", id).Updates(map[string]interface{}{"name": name})
	db.Where("id = ?", id).Find(&vendor)
	return &vendor, nil
}

func RemoveVendor(id int) (*Vendor, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	vendor := Vendor{Id: id}
	db.Where("id = ?", id).Find(&vendor)
	return &vendor, db.Delete(vendor).Error
}