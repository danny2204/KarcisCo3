package models

import "github.com/danny2204/KarcisCoAPI/connection"

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

	db.AutoMigrate(&Airport{})
	db.AutoMigrate(&Airline{})
	db.AutoMigrate(&Vendor{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Train{})
	db.AutoMigrate(&Station{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Location{})
	db.AutoMigrate(&Admin{})

	db.AutoMigrate(&Hotel{}).AddForeignKey("location_refer", "locations(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&HotelFacility{}).AddForeignKey("hotel_refer_id", "hotels(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Facility{}).AddForeignKey("flight_id", "flights(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Route{}).AddForeignKey("flight_route_id", "flights(id)", "CASCADE", "CASCADE").AddForeignKey("route_from_refer", "airports(id)", "CASCADE", "CASCADE").AddForeignKey("route_to_refer", "airports(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Flight{}).AddForeignKey("airline_refer", "airlines(id)", "CASCADE", "CASCADE").AddForeignKey("to_refer", "airports(id)", "CASCADE", "CASCADE").AddForeignKey("from_refer", "airports(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Car{}).AddForeignKey("vendor_refer", "vendors(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&Trip{}).AddForeignKey("train_refer", "trains(id)", "CASCADE", "CASCADE").AddForeignKey("from_refer", "stations(id)", "CASCADE", "CASCADE").AddForeignKey("to_refer", "stations(id)", "CASCADE", "CASCADE")
	//.AddForeignKey("room_refer", "rooms(id)", "CASCADE", "CASCADE")

	 InsertLocationData()

}
