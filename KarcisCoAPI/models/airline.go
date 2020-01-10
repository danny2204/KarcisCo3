package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Airline struct {
	ID int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name string
	Code string
	Seat int
}

func init() {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&Airline {})
}

func GetAllAirline()([]Airline, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var airlineList []Airline
	db.Find(&airlineList)

	return airlineList, err
}

func InsertAirline(name string, code string, seat int) (*Airline, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	airline := &Airline{
		Name:      name,
		Code: code,
		Seat: seat,
	}

	db.Save(airline)
	return airline, err
}

func UpdateAirline(id int, name string, code string, seat int) (*Airline, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var airline Airline
	db.Model(&airline).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "code": code, "seat": seat})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&airline)
	return &airline, nil
}

func RemoveAirline(id int) (*Airline, error) {
	db, err := connection.ConnectDatabase()

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

	if err != nil {
		return nil, err
	}
	defer db.Close()

	airline := Airline{ID: id}
	db.Where("id = ?", id).Find(&airline)
	return &airline, err
}