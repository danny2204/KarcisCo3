package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Station struct {
	Id			int			`gorm:primary_key`
	Code 		string		`gorm:"type:varchar(100);"`
	Name		string 		`gorm:"type:varchar(100);"`
	Location Location `gorm:"foreignkey:LocationReferId"`
	LocationReferId int
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

func GetAllStation()([]Station, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var stationList []Station
	db.Find(&stationList)
	for i,_ := range stationList {
		db.Model(stationList[i]).Related(&stationList[i].Location, "LocationReferId")
	}
	return stationList, err
}

func CreateStation(code string, name string, locationReferId int) (*Station, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	station := &Station{
		Code: code,
		Name: name,
		LocationReferId: locationReferId,
	}
	db.Save(station)
	return station, err
}

func UpdateStation(id int, code string, name string, locationReferId int) (*Station, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var station Station
	db.Model(&station).Where("id = ?", id).Updates(map[string]interface{}{"code": code, "name": name, "location_refer_id": locationReferId})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&station)
	return &station, nil
}

func RemoveStation(id int) (*Station, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

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
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	station := Station{Id: id}
	db.Where("id = ?", id).Find(&station)
	return &station, err
}