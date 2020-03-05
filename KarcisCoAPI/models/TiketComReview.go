package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
)

type TiketComReview struct {
	Id int `gorm:"primary_key"`
	AverageScore float64
	LocationScore float64
	CleanlinessScore float64
	ServiceScore float64
	ValueScore float64
	Reviews []Review `gorm:"foreignkey:TiketComReviewReferId"`
}

func GetAllTiketComReview()([]TiketComReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var tiketComReview []TiketComReview
	db.Find(&tiketComReview)
	for i,_ := range tiketComReview {
		db.Model(tiketComReview[i]).Related(&tiketComReview[i].Reviews, "TiketComReviewReferId")
	}
	return tiketComReview, err
}

func CreateTiketComReview(averageScore float64, locationScore float64, cleanlinessScore float64, serviceStore float64, valueScore float64)(*TiketComReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tiketComReview := &TiketComReview{
		AverageScore: averageScore,
		LocationScore: locationScore,
		CleanlinessScore: cleanlinessScore,
		ServiceScore: serviceStore,
		ValueScore: valueScore,
	}
	db.Save(&tiketComReview)
	return tiketComReview, err
}

func UpdateTiketComReview(id int, averageScore float64, locationScore float64, cleanlinessScore float64, serviceStore float64, valueScore float64)(*TiketComReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	var tiketComReview TiketComReview
	db.Model(&tiketComReview).Where("id = ?", id).Update(map[string]interface{}{"average_score": averageScore, "location_score": locationScore, "cleanliness_score": cleanlinessScore, "service_score": serviceStore, "value_score": valueScore})
	db.Where("id = ?", id).Find(&tiketComReview)
	return &tiketComReview, nil
}

func RemoveTiketComReview(id int) (*TiketComReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	tiketComReview := TiketComReview{Id: id}
	db.Where("id = ?", id).Find(&tiketComReview)
	return &tiketComReview, db.Delete(tiketComReview).Error
}
