package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Room struct{
	Id			int				`gorm:primary_key`
	Name		string			`gorm:varchar(100);not null`
	RoomSize	int	 			`gorm:int`
	Price		int				`gorm:int`
	MaxPassenger int
	FreeBreakfast bool
	FreeWifi bool
	Bed string `gorm:varchar(500);not null`
	RoomLeft int
	RoomReferId int
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func GetAllRoom()([]Room, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var roomList []Room
	db.Find(&roomList)
	return roomList, err
}

func CreateRoom(name string, roomSize int, price int, maxPassenger int, freeBreakfast bool, freeWifi bool, bed string, roomLeft int, roomReferId int) (*Room, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	room := &Room{
		Name: name,
		RoomSize: roomSize,
		Price: price,
		MaxPassenger: maxPassenger,
		FreeBreakfast: freeBreakfast,
		FreeWifi: freeWifi,
		Bed: bed,
		RoomLeft: roomLeft,
		RoomReferId: roomReferId,
	}

	db.Save(room)
	return room, err
}

func UpdateRoom(id int, name string, roomSize int, price int, maxPassenger int, freeBreakfast bool, freeWifi bool, bed string, roomLeft int, hotelReferId int) (*Room, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var room Room
	db.Model(&room).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "room_size": roomSize, "price": price, "max_passenger": maxPassenger, "free_breakfast": freeBreakfast, "free_wifi": freeWifi, "bed": bed, "room_left": roomLeft, "hotel_refer_id": hotelReferId})
	db.Where("id = ?", id).Find(&room)
	return &room, nil
}

func RemoveRoom(id int) (*Room, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	room := Room{Id: id}
	db.Where("id = ?", id).Find(&room)
	return &room, db.Delete(room).Error
}
