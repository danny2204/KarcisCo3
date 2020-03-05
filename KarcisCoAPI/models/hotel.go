package models

import (
	"fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Hotel struct{
	Id 				int				`gorm: primary_key`
	Name 			string			`gorm:"type:varchar(100);not null"`
	Location		Location		`gorm:"foreignkey:LocationRefer"`
	LocationRefer	int
	Address			string			`gorm:"type:varchar(100);not null"`
	Rating			int
	Type string
	Longitude float64
	Latitude float64
	TripAdvisorReview TripAdvisorReview `gorm:"foreignkey:TripAdvisorReviewReferId"`
	TripAdvisorReviewReferId int
	TiketComReview TiketComReview  `gorm:"foreignkey:TiketComReviewReferId"`
	TiketComReviewReferId int
	HotelFacilities []HotelFacility `gorm:"foreignkey:HotelReferId"`
	Room []Room `gorm:foreignkey:RoomReferId`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`
}

func GetAllHotel()([]Hotel, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var hotelList []Hotel
	db.Find(&hotelList)
	for i,_ := range hotelList {
		db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
		db.Model(hotelList[i]).Related(&hotelList[i].HotelFacilities, "HotelReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].Room, "RoomReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TripAdvisorReview, "TripAdvisorReviewReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TiketComReview, "TiketComReviewReferId")
		db.Model(hotelList[i].TiketComReview).Related(&hotelList[i].TiketComReview.Reviews, "TiketComReviewReferId")
		db.Model(hotelList[i].TripAdvisorReview).Related(&hotelList[i].TripAdvisorReview.Reviews, "TripAdvisorReviewReferId")
		for j,_ := range hotelList[i].TiketComReview.Reviews {
			db.Model(hotelList[i].TiketComReview).Related(&hotelList[i].TiketComReview.Reviews[j], "TiketComReviewReferId")
		}
		for j,_ := range hotelList[i].TripAdvisorReview.Reviews {
			db.Model(hotelList[i].TripAdvisorReview).Related(&hotelList[i].TripAdvisorReview.Reviews[j], "TripAdvisorReviewReferId")
		}
	}
	fmt.Println(hotelList)
	return hotelList, err
}

func CreateHotel(name string, locationRefer int, address string, rating int, latitude float64, longitude float64, tripAdvisorReviewReferId int, tiketComReviewReferId int, types string) (*Hotel, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	hotel := &Hotel{
		Name: name,
		LocationRefer: locationRefer,
		Address: address,
		Rating: rating,
		Latitude: latitude,
		Longitude: longitude,
		TripAdvisorReviewReferId: tripAdvisorReviewReferId,
		TiketComReviewReferId: tiketComReviewReferId,
		Type: types,
	}

	db.Save(hotel)
	return hotel, err
}

func UpdateHotel(id int, name string, locationRefer int, address string, rating int, latitude float64, longitude float64, tripAdvisorReviewReferId int, tiketComReviewReferId int, types string) (*Hotel, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotel Hotel
	db.Model(&hotel).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "location_refer": locationRefer, "address": address, "rating": rating, "latitude": latitude, "longitude": longitude, "trip_advisor_review_refer_id": tripAdvisorReviewReferId, "tiket_com_review_refer_id": tiketComReviewReferId, "type": types})
	db.Where("id = ?", id).Find(&hotel)
	return &hotel, nil
}

func RemoveHotel(id int) (*Hotel, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

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
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotelList []Hotel
	db.Where("location_refer in (?)", db.Table("locations").Select("id").Where("city = ?", city).SubQuery()).Find(&hotelList)

	for i,_ := range hotelList{
		db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
		db.Model(hotelList[i]).Related(&hotelList[i].HotelFacilities, "HotelReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].Room, "RoomReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TripAdvisorReview, "TripAdvisorReviewReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TiketComReview, "TiketComReviewReferId")
		db.Model(hotelList[i].TiketComReview).Related(&hotelList[i].TiketComReview.Reviews, "TiketComReviewReferId")
		db.Model(hotelList[i].TripAdvisorReview).Related(&hotelList[i].TripAdvisorReview.Reviews, "TripAdvisorReviewReferId")
	}
	return hotelList, err
}

func GetHotelById(id int)([]Hotel, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotelList []Hotel
	db.Where("id = ?", id).Find(&hotelList)
	for i,_ := range hotelList {
		db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
		db.Model(hotelList[i]).Related(&hotelList[i].HotelFacilities, "HotelReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].Room, "RoomReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TripAdvisorReview, "TripAdvisorReviewReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TiketComReview, "TiketComReviewReferId")
		db.Model(hotelList[i].TiketComReview).Related(&hotelList[i].TiketComReview.Reviews, "TiketComReviewReferId")
		db.Model(hotelList[i].TripAdvisorReview).Related(&hotelList[i].TripAdvisorReview.Reviews, "TripAdvisorReviewReferId")
		for j,_ := range hotelList[i].TiketComReview.Reviews {
			db.Model(hotelList[i].TiketComReview).Related(&hotelList[i].TiketComReview.Reviews[j], "TiketComReviewReferId")
		}
		for j,_ := range hotelList[i].TripAdvisorReview.Reviews {
			db.Model(hotelList[i].TripAdvisorReview).Related(&hotelList[i].TripAdvisorReview.Reviews[j], "TripAdvisorReviewReferId")
		}
	}
	return hotelList, err
}

func GetHotelByGeolocation(longitude float64, latitude float64)([]Hotel, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var hotelList []Hotel
	db.Where("id in (?)", db.Raw("SELECT id FROM ( SELECT id, (3959 * acos(cos(radians(?)) * cos(radians(latitude)) * cos(radians(longitude) - radians(?)) + sin(radians(?)) * sin(radians(latitude )))) AS distance FROM hotels ORDER BY distance) AS nearestHotel", latitude, longitude, latitude).SubQuery()).Limit(8).Find(&hotelList)
	for i,_ := range hotelList{
		db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
		db.Model(hotelList[i]).Related(&hotelList[i].HotelFacilities, "HotelReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].Room, "RoomReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TripAdvisorReview, "TripAdvisorReviewReferId")
		db.Model(hotelList[i]).Related(&hotelList[i].TiketComReview, "TiketComReviewReferId")
		db.Model(hotelList[i].TiketComReview).Related(&hotelList[i].TiketComReview.Reviews, "TiketComReviewReferId")
		db.Model(hotelList[i].TripAdvisorReview).Related(&hotelList[i].TripAdvisorReview.Reviews, "TripAdvisorReviewReferId")
	}
	return hotelList, err
}