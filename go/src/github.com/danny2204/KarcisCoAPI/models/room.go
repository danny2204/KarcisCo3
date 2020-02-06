package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Room struct{
	Id			int				`gorm:primary_key`
	Name		string			`gorm:varchar(100);not null`
	//Image
	//Service
	RoomSize	int	 			`gorm:int`
	Price		int				`gorm:int`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func GetAllRoom()([]Room, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var roomList []Room
	db.Find(&roomList)
	//for i,_ := range hotelList {
	//	db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
	//}
	return roomList, err
}

func CreateRoom(name string, roomsize int, price int) (*Room, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	room := &Room{
		Name: name,
		RoomSize: roomsize,
		Price: price,
	}

	db.Save(room)
	return room, err
}

func UpdateRoom(id int, name string, roomsize int, price int) (*Room, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var room Room
	db.Model(&room).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "room_size": roomsize, "price": price})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&room)
	return &room, nil
}

func RemoveRoom(id int) (*Room, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	room := Room{Id: id}
	db.Where("id = ?", id).Find(&room)
	return &room, db.Delete(room).Error
}

func GetRoomById(id int)(*Room, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	room := Room{Id: id}
	db.Where("id = ?", id).Find(&room)
	return &room, err
}