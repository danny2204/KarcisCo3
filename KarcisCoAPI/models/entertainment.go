package models

import (
	"fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Entertainment struct {
	Id int `gorm:primary_key`
	Name string
	Description string
	Location Location `gorm:"foreignkey:LocationReferId"`
	LocationReferId int
	Type string
	IsTrending bool
	UploadDate time.Time
	EntertainmentTicket []EntertainmentTicket `gorm:"foreignkey:EntertainmentReferId"`
	EventStartDate time.Time
	EventEndDate time.Time
	CreatedAt time.Time
	UpdateAt time.Time
	DeletedAt *time.Time `sql:index`
}

func GetAllEntertainment()([]Entertainment, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var entertainmentList []Entertainment
	db.Find(&entertainmentList)
	for i,_ := range entertainmentList {
		db.Model(entertainmentList[i]).Related(&entertainmentList[i].Location, "LocationReferId")
		db.Model(entertainmentList[i]).Related(&entertainmentList[i].EntertainmentTicket, "EntertainmentReferId")
	}
	return entertainmentList, err
}

func CreateEntertainment(name string, locationReferId int, types string, eventStartDate time.Time, eventEndDate time.Time, isTrending bool, description string)(*Entertainment, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	entertainment := &Entertainment {
		Name: name,
		Type: types,
		LocationReferId: locationReferId,
		EventStartDate: eventStartDate,
		EventEndDate: eventEndDate,
		IsTrending: isTrending,
		UploadDate: time.Now(),
		Description: description,
	}
	db.Save(entertainment)
	return entertainment, err
}

func UpdateEntertainment(id int, name string, locationReferId int, types string, eventStartDate time.Time, eventEndDate time.Time, isTrending bool, description string)(*Entertainment, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainment Entertainment
	db.Model(&entertainment).Where("id = ?", id).Update(map[string]interface{}{"name": name, "type": types, "location_refer_id": locationReferId, "event_start_date": eventStartDate, "event_end_date": eventEndDate, "is_trending": isTrending, "upload_date": time.Now(), "description": description})
	db.Where("id = ?", id).Find(&entertainment)
	return &entertainment, nil
}

func RemoveEntertainment(id int)(*Entertainment, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	entertainment := Entertainment{Id: id}
	db.Where("id = ?", id).Find(&entertainment)
	return &entertainment, db.Delete(entertainment).Error

}

func GetEntertainmentByLocation(location string)([]Entertainment, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainment []Entertainment
	db.Where("location_refer_id = ?", db.Table("locations").Select("id").Where("city = ?", location).SubQuery()).Find(&entertainment)
	for i,_ := range entertainment {
		db.Model(entertainment[i]).Related(&entertainment[i].Location, "LocationReferId")
		db.Model(entertainment[i]).Related(&entertainment[i].EntertainmentTicket, "EntertainmentReferId")
	}
	return entertainment, err
}

func GetEntertainmentByType(types string)([]Entertainment, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainmentList []Entertainment
	db.Where("type = ?", types).Find(&entertainmentList)
	for i,_ := range entertainmentList {
		db.Model(entertainmentList[i]).Related(&entertainmentList[i].Location, "LocationReferId")
		db.Model(entertainmentList[i]).Related(&entertainmentList[i].EntertainmentTicket, "EntertainmentReferId")
	}
	return entertainmentList, err
}

func GetEntertainmentByTrending(types string)([]Entertainment, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainmentList []Entertainment
	db.Where("type = ? AND is_trending = true", types).Find(&entertainmentList).Limit(3)
	for i,_ := range entertainmentList {
		db.Model(entertainmentList[i]).Related(&entertainmentList[i].Location, "LocationReferId")
		db.Model(entertainmentList[i]).Related(&entertainmentList[i].EntertainmentTicket, "EntertainmentReferId")
	}
	return entertainmentList, err
}

func GetEntertainmentById(id int)(*Entertainment, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainment Entertainment
	db.Where("id = ?", id).Find(&entertainment)
	db.Model(entertainment).Related(&entertainment.Location, "LocationReferId")
	db.Model(entertainment).Related(&entertainment.EntertainmentTicket, "EntertainmentReferId")
	fmt.Println(entertainment)
	return &entertainment, err
}

