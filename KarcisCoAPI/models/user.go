package models

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
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
	GoogleId string
	FacebookId string
	Gender string
	Title string `gorm:"type:varchar(100);"`
	City string `gorm:"type:varchar(100);"`
	Address string `gorm:"type:varchar(100);"`
	PostCode string `gorm:"type:varchar(100);"`
	MainLanguage string
}

func GetAllUser()([]User, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var userList []User
	db.Find(&userList)

	return userList, err
}

func CreateUser(frontName string, backName string, email string, password string, phoneNumber string, title string, city string, address string, postcode string, googleId string, facebookId string, gender string, mainLanguage string) (*User, error){
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
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
		GoogleId: googleId,
		FacebookId: facebookId,
		Gender: gender,
		MainLanguage: mainLanguage,
	}

	db.Save(user)
	return user, err
}

func UpdateUser(id int, frontName string, backName string, email string, password string, phoneNumber string, title string, city string, address string, postcode string, googleId string, facebookId string, gender string, mainLanguage string) (*User, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{"front_name": frontName, "back_name": backName, "email": email, "password": password, "phone_number": phoneNumber, "title": title, "address": address, "city": city, "post_code": postcode, "google_id": googleId, "facebook_id": facebookId, "gender": gender, "main_language": mainLanguage})
	//db.Where(Admin{ID: id}).Assign(Admin{Name: name, Email: email, Password: password}).FirstOrCreate(&admin)
	db.Where("id = ?", id).Find(&user)
	return &user, nil
}

func GetUserByEmailAndPhone(arg string) ([]User, error) {
	db, err := connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)

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
	_, err = GetApiKeyDetail(middleware.ApiKey)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := User{ID: id}
	db.Where("id = ?", id).Find(&user)
	return &user, db.Delete(user).Error
}

func Send(toEmails string) {
	token, err := ioutil.ReadFile("models/credentials.json")
	if err != nil {
		panic(err)
	}

	config, err := google.ConfigFromJSON(token, gmail.MailGoogleComScope)
	if err != nil {
		log.Fatalf("Gagal parse config nya, hati hati kalau credentials.json nya di utak atik")
	}
	client := getClient(config)

	service, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	FromEmail := "dannyld224@gmail.com"
	toEmail := toEmails
	pesan := "HALO"

	var message gmail.Message
	temp := []byte("From: 'me'\r\n" +
		"reply-to:" + FromEmail + "\r\n" +
		"To:  " + toEmail + "\r\n" +
		"Subject:  \r\n" +
		"\r\n" + pesan)
	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)
	_, err = service.Users.Messages.Send("me", &message).Do()
	if err != nil {
		fmt.Println(err)
	}
}

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "models/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func GetUserByAuthId(id string) ([]User, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var user []User
	if db.Where("facebook_id = ? OR google_id = ?", id, id).Find(&user).RecordNotFound(){
		return nil, nil
	}
	return user, err

}
