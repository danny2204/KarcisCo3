package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
)

func init() {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//db.DropTableIfExists(&Flight{})
	//db.DropTableIfExists(&Airline{})
	//db.DropTableIfExists(&Airport{})
	//db.DropTableIfExists(&Route{})
	//
	//db.DropTableIfExists(&Car{})
	//db.DropTableIfExists(&Vendor{})
	//
	//db.DropTableIfExists(&Hotel{})
	//db.DropTableIfExists(&Location{})
	//db.DropTableIfExists(&Room{})
	//
	//db.DropTableIfExists(&Trip{})
	//db.DropTableIfExists(&Station{})
	//db.DropTableIfExists(&Train{})
	//
	//db.DropTableIfExists(&User{})
	db.AutoMigrate(&ApiKey{})
	db.AutoMigrate(&Airport{})
	db.AutoMigrate(&Airline{})
	db.AutoMigrate(&Car{})
	db.AutoMigrate(&Vendor{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Train{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Admin{})
	db.AutoMigrate(&TiketComReview{})
	db.AutoMigrate(&TripAdvisorReview{})
	db.AutoMigrate(&Entertainment{})
	db.AutoMigrate(&Promo{})
	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&Chat{})
	db.AutoMigrate(&PromoDetail{}).AddForeignKey("promo_refer_id", "promos(id)", "CASCADE", "CASCADE")

	InsertLocationData()
	db.AutoMigrate(&EntertainmentTicket{}).AddForeignKey("entertainment_refer_id", "entertainments(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Review{}).AddForeignKey("tiket_com_review_refer_id", "tiket_com_reviews(id)", "CASCADE", "CASCADE").AddForeignKey("trip_advisor_review_refer_id", "trip_advisor_reviews(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&CarVendor{}).AddForeignKey("vendor_refer_id", "vendors(id)", "CASCADE", "CASCADE").AddForeignKey("car_refer_id", "cars(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&VendorLocation{}).AddForeignKey("location_refer", "locations(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Hotel{}).AddForeignKey("location_refer", "locations(id)", "CASCADE", "CASCADE").AddForeignKey("trip_advisor_review_refer_id", "trip_advisor_reviews(id)", "CASCADE", "CASCADE").AddForeignKey("tiket_com_review_refer_id", "tiket_com_reviews(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&HotelFacility{}).AddForeignKey("hotel_refer_id", "hotels(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Facility{}).AddForeignKey("flight_id", "flights(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Route{}).AddForeignKey("flight_route_id", "flights(id)", "CASCADE", "CASCADE").AddForeignKey("route_from_refer", "airports(id)", "CASCADE", "CASCADE").AddForeignKey("route_to_refer", "airports(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Flight{}).AddForeignKey("airline_refer", "airlines(id)", "CASCADE", "CASCADE").AddForeignKey("to_refer", "airports(id)", "CASCADE", "CASCADE").AddForeignKey("from_refer", "airports(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Station{}).AddForeignKey("location_refer_id", "locations(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Trip{}).AddForeignKey("train_refer", "trains(id)", "CASCADE", "CASCADE").AddForeignKey("from_refer", "stations(id)", "CASCADE", "CASCADE").AddForeignKey("to_refer", "stations(id)", "CASCADE", "CASCADE")
	//.AddForeignKey("room_refer", "rooms(id)", "CASCADE", "CASCADE")

}
