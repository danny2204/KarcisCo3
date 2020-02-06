package models

import (
	_ "fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Airport struct {
	Id			int			`gorm:primary_key`
	Code		string		`gorm: "type:char(3);not null"`
	Name		string		`gorm: "type:varchar(100);not null"`
	City		string		`gorm: "type:varchar(100);not null"`
	CityCode	string		`gorm: "type:char(4);not null"`
	Province	string		`gorm: "type:varchar(100);not null"`
	Country		string		`gorm: "type:varchar(100);not null"`
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`
}

func GetAllAirport()([]Airport, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var airportList []Airport
	db.Find(&airportList)

	return airportList, err
}

func InsertAirport(code string, name string, city string, citycode string, province string, country string) (*Airport, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	airport := &Airport{
		Code: code,
		Name:      name,
		City: city,
		CityCode: citycode,
		Province: province,
		Country: country,
	}

	db.Save(airport)
	return airport, err
}

func UpdateAirport(id int, name string, code string, city string, citycode string, province string, country string) (*Airport, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var airport Airport
	db.Model(&airport).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "city": city, "code": code, "city_code": citycode, "province": province, "country": country})
	db.Where("id = ?", id).Find(&airport)
	return &airport, nil
}

func RemoveAirport(id int) (*Airport, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airport := Airport{Id: id}
	db.Where("id = ?", id).Find(&airport)
	return &airport, db.Delete(airport).Error
}

func GetAirportbyId(id int) (*Airport, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airport := Airport{Id: id}
	db.Where("id = ?", id).Find(&airport)
	return &airport, err
}