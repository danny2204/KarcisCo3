package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Airline struct {
	ID int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name string `gorm:varchar(100);not null`
}

func GetAllAirline()([]Airline, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var airlineList []Airline
	db.Find(&airlineList)

	return airlineList, err
}

func InsertAirline(name string) (*Airline, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	airline := &Airline{
		Name:      name,
	}

	db.Save(airline)
	return airline, err
}

func UpdateAirline(id int, name string) (*Airline, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var airline Airline
	db.Model(&airline).Where("id = ?", id).Updates(map[string]interface{}{"name": name})
	db.Where("id = ?", id).Find(&airline)
	return &airline, nil
}

func RemoveAirline(id int) (*Airline, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airline := Airline{ID: id}
	db.Where("id = ?", id).Find(&airline)
	return &airline, db.Delete(airline).Error
}

func GetAirlineById(id int) (*Airline, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airline := Airline{ID: id}
	db.Where("id = ?", id).Find(&airline)
	return &airline, err
}