package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"time"
)

type User struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	FrontName      string     `gorm:"type:varchar(100);not null"`
	BackName string `gorm:"type:varchar(100);not null"`
	Email     string     `gorm:"type:varchar(100);not null;primary_key"`
	Password  string     `gorm:"type:varchar(100);not null;primary_key"`
	PhoneNumber string `gorm:"type:varchar(50);not null""`
	Title string `gorm:"type:varchar(100);"`
	City string `gorm:"type:varchar(100);"`
	Address string `gorm:"type:varchar(100);"`
	PostCode string `gorm:"type:varchar(100);"`
}

func GetAllUser()([]User, error) {
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var userList []User
	db.Find(&userList)

	return userList, err
}

func CreateUser(frontName string, backName string, email string, password string, phoneNumber string, title string, city string, address string, postcode string) (*User, error){
	db, err = connection.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	user := &User{
		FrontName: frontName,
		BackName: backName,
		Email:     email,
		Password:  password,
		PhoneNumber: phoneNumber,
		Title: title,
		City: city,
		Address: address,
		PostCode: postcode,
	}

	db.Save(user)
	return user, err
}

func UpdateUser(id int, frontName string, backName string, email string, password string, phoneNumber string, title string, city string, address string, postcode string) (*User, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"front_name": frontName, "back_name": backName, "email": email, "password": password, "phone_number": phoneNumber, "title": title, "address": address, "city": city, "post_code": postcode})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&user)
	return &user, nil
}

func GetUserByEmailAndPhone(arg string) ([]User, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	var user []User
	if db.Where("email = ?", arg).Or("phone_number = ?", arg).Find(&user).RecordNotFound(){
		return nil, nil
	}
	return user, err
}

func RemoveUser(id int) (*User, error) {
	db, err := connection.ConnectDatabase()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := User{ID: id}
	db.Where("id = ?", id).Find(&user)
	return &user, db.Delete(user).Error
}