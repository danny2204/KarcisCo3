package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Station struct {
	Id			int			`gorm:primary_key`
	Code 		string		`gorm:"type:varchar(100);not null"`
	Name		string 		`gorm:"type:varchar(100);not null"`
	City		string 		`gorm:"type:varchar(100);not null"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

//func init(){
//	db, err := connection.ConnectDatabase()
//
//	if err!=nil{
//		panic(err)
//	}
//
//	defer db.Close()
//
//	////1
//	//db.Create(&Station{
//	//	Code:      "JAKK",
//	//	Name:      "Jakarta Kota",
//	//	City:      "Jakarta",
//	//})
//	//
//	////2
//	//db.Create(&Station{
//	//	Code:      "SMT",
//	//	Name:      "Semarangtawang",
//	//	City:      "Semarang",
//	//})
//	//
//	////3
//	//db.Create(&Station{
//	//	Code:      "PSE",
//	//	Name:      "Pasar Senen",
//	//	City:      "Jakarta",
//	//})
//	//
//	////4
//	//db.Create(&Station{
//	//	Code:      "GMR",
//	//	Name:      "Gambir",
//	//	City:      "Jakarta",
//	//})
//	//
//	//
//	////5
//	//db.Create(&Station{
//	//	Code:      "CMI",
//	//	Name:      "Cimahi",
//	//	City:      "Bandung",
//	//})
//	//
//	////6
//	//db.Create(&Station{
//	//	Code:      "KAC",
//	//	Name:      "Kiaracondong",
//	//	City:      "Bandung",
//	//})
//	//
//	////7
//	//db.Create(&Station{
//	//	Code:      "BD",
//	//	Name:      "Bandung",
//	//	City:      "Bandung",
//	//})
//
//}

func GetAllStation()([]Station, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var stationList []Station
	db.Find(&stationList)
	//for i,_ := range hotelList {
	//	db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
	//}
	return stationList, err
}

func CreateStation(code string, name string, city string) (*Station, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	station := &Station{
		Code: code,
		Name: name,
		City: city,
	}

	db.Save(station)
	return station, err
}

func UpdateStation(id int, code string, name string, city string) (*Station, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var station Station
	db.Model(&station).Where("id = ?", id).Updates(map[string]interface{}{"code": code, "name": name, "city": city})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&station)
	return &station, nil
}

func RemoveStation(id int) (*Station, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	station := Station{Id: id}
	db.Where("id = ?", id).Find(&station)
	return &station, db.Delete(station).Error
}

func GetStationById(id int)(*Station, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	station := Station{Id: id}
	db.Where("id = ?", id).Find(&station)
	return &station, err
}