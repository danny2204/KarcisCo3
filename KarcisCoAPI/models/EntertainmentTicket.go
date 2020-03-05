package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type EntertainmentTicket struct {
	Id int `gorm:"primary_key"`
	Type string
	Price int
	Description string
	EntertainmentReferId int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func GetAllEntertainmentTicket()([]EntertainmentTicket, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainmentTicketList []EntertainmentTicket
	db.Find(&entertainmentTicketList)
	return entertainmentTicketList, err
}

func CreateEntertainmentTicket(types string, price int, entertainmentReferId int, description string)(*EntertainmentTicket, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	entertainmentTicket := &EntertainmentTicket{
		Type: types,
		Price: price,
		EntertainmentReferId: entertainmentReferId,
		Description: description,
	}

	db.Save(entertainmentTicket)
	return entertainmentTicket, err
}

func UpdateEntertainmentTicket(id int, types string, price int, entertainmentReferId int, description string)(*EntertainmentTicket, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var entertainment EntertainmentTicket
	db.Model(&entertainment).Where("id = ?", id).Update(map[string]interface{}{"type": types, "price": price, "entertainment_refer_id": entertainmentReferId, "description": description})
	db.Where("id = ?", id).Find(&entertainment)
	return &entertainment, nil
}

func RemoveEntertainmentTicket(id int)(*EntertainmentTicket, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	entertainmentTicket := EntertainmentTicket{Id: id}
	db.Where("id = ?", id).Find(&entertainmentTicket)
	return &entertainmentTicket, db.Delete(entertainmentTicket).Error
}

//func GetEntertainmentTicketById(id int)(*EntertainmentTicket, error) {
//	db, err := connection.ConnectDatabase()
//
//	if err != nil {
//		return nil, err
//	}
//	defer db.Close()
//
//}
