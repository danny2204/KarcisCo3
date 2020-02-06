package models

import (
	"fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type Hotel struct{
	Id 				int				`gorm: primary_key`
	Name 			string			`gorm:"type:varchar(100);not null"`
	Location		Location		`gorm:"foreignkey:LocationRefer"`
	LocationRefer	int
	Address			string			`gorm:"type:varchar(100);not null"`
	Rating			int
	ImagePath string
	HotelFacilities []HotelFacility `gorm:"foreignkey:HotelReferId"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`
}

func GetAllHotel()([]Hotel, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var hotelList []Hotel
	db.Find(&hotelList)
	for i,_ := range hotelList {
		db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
		db.Model(hotelList[i]).Related(&hotelList[i].HotelFacilities, "HotelReferId")
	}
	fmt.Println(hotelList)
	return hotelList, err
}

func CreateHotel(name string, locationrefer int, address string, rating int, imagePath string) (*Hotel, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	hotel := &Hotel{
		Name: name,
		LocationRefer: locationrefer,
		Address: address,
		Rating: rating,
		ImagePath: imagePath,
	}

	db.Save(hotel)
	return hotel, err
}

func UpdateHotel(id int, name string, locationrefer int, address string, rating int, imagePath string) (*Hotel, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotel Hotel
	db.Model(&hotel).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "location_refer": locationrefer, "address": address, "rating": rating, "image_path": imagePath})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&hotel)
	return &hotel, nil
}

func RemoveHotel(id int) (*Hotel, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	hotel := Hotel{Id: id}
	db.Where("id = ?", id).Find(&hotel)
	return &hotel, db.Delete(hotel).Error
}

func GetHotelByCity(city string)([]Hotel, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotels []Hotel
	db.Where("location_refer in (?)", db.Table("locations").Select("id").Where("city = ?", city).SubQuery()).Find(&hotels)

	for i,_ := range hotels{
		db.Model(hotels[i]).Related(&hotels[i].Location, "LocationRefer")
		db.Model(hotels[i]).Related(&hotels[i].HotelFacilities, "HotelReferId")
	}
	return hotels, err
}