package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
)

type Train struct {
	Id					int 		`gorm:primary_key`
	Name				string 		`gorm:"type:varchar(100);not null"`
	Class				string		`gorm:"type:varchar(100);not null"`
	Subclass			string		`gorm:"type:varchar(100);not null"`
	Code				string		`gorm:"type:varchar(100);not null"`
	//Image
}

func GetAllTrain()([]Train, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var trainList []Train
	db.Find(&trainList)
	//for i,_ := range hotelList {
	//	db.Model(hotelList[i]).Related(&hotelList[i].Location, "LocationRefer")
	//}
	return trainList, err
}

func CreateTrain(name string, class string, subclass string, code string) (*Train, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	train := &Train{
		Name: name,
		Class: class,
		Subclass: subclass,
		Code: code,
	}

	db.Save(train)
	return train, err
}

func UpdateTrain(id int, name string, class string, subclass string, code string) (*Train, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var train Train
	db.Model(&train).Where("id = ?", id).Updates(map[string]interface{}{"name": name, "class": class, "subclass": subclass, "code": code})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&train)
	return &train, nil
}

func RemoveTrain(id int) (*Train, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	train := Train{Id: id}
	db.Where("id = ?", id).Find(&train)
	return &train, db.Delete(train).Error
}

func GetTrainById(id int)(*Train, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	train := Train{Id: id}
	db.Where("id = ?", id).Find(&train)
	return &train, err
}