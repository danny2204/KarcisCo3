package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Location struct{
	Id        int     `gorm:primary_key`
	Longitude float64 `gorm:"type:decimal(10,4)"`
	Latitude  float64 `gorm:"type:decimal(10,4)"`
	ZIndex    int
	City      string `gorm:"type:varchar(255);"`
	Province  string `gorm:"type:varchar(255);"`
	Country   string `gorm:"type:varchar(255);"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

func GetAllLocation()([]Location, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var locationList []Location
	db.Find(&locationList)
	//for i,_ := range hotelList {
	//	db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
	//}
	return locationList, err
}

func CreateLocation(longitude float64, latitude float64, city string, province string, country string, zindex int) (*Location, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	location := &Location{
		Longitude: longitude,
		Latitude: latitude,
		City: city,
		Province: province,
		Country: country,
		ZIndex: zindex,
	}

	db.Save(location)
	return location, err
}

func UpdateLocation(id int, longitude float64, latitude float64, city string, province string, country string, zindex int) (*Location, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var location Location
	db.Model(&location).Where("id = ?", id).Updates(map[string]interface{}{"longitude": longitude, "latitude": latitude, "city": city, "province": province, "country": country, "z_index": zindex})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&location)
	return &location, nil
}

func RemoveLocation(id int) (*Location, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	location := Location{Id: id}
	db.Where("id = ?", id).Find(&location)
	return &location, db.Delete(location).Error
}

func GetLocationById(id int)(*Location, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	location := Location{Id: id}
	db.Where("id = ?", id).Find(&location)
	return &location, err
}