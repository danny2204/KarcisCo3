package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Car struct {
	Id			int				`gorm:primary_key`
	Brand		string			`gorm:"type:varchar(100);not null"`
	Model		string			`gorm:"type:varchar(100);not null"`
	Price		int				`gorm:int`
	Vendor Vendor			`gorm:"foreignkey:VendorRefer"`
	VendorRefer int
	Passenger	int 			`gorm:int`
	Luggage		int				`gorm:int`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func GetCar()([]Car, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var carList []Car
	db.Find(&carList)

	return carList, err
}

func InsertCar(brand string, model string, price int, vendorId int, passenger int, luggage int) (*Car, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	car := &Car{
		Brand: brand,
		Model: model,
		Price: price,
		VendorRefer: vendorId,
		Passenger: passenger,
		Luggage: luggage,
	}

	db.Save(car)
	return car, err
}

func UpdateCar(id int, brand string, model string, price int, vendorId int, passenger int, luggage int) (*Car, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var car Car
	db.Model(&car).Where("id = ?", id).Updates(map[string]interface{}{"brand": brand, "model": model, "price": price, "vendor_refer": vendorId, "passenger": passenger, "luggage": luggage})
	db.Where("id = ?", id).Find(&car)
	return &car, nil
}

func RemoveCar(id int) (*Car, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	car := Car{Id: id}
	db.Where("id = ?", id).Find(&car)
	return &car, db.Delete(car).Error
}

func GetCarById(id int) (*Car, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	car := Car{Id: id}
	db.Where("id = ?", id).Find(&car)
	return &car, err
}