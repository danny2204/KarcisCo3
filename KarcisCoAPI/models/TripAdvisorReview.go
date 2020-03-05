package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
)

type TripAdvisorReview struct {
	Id int `gorm:"primary_key"`
	AverageScore float64
	LocationScore float64
	CleanlinessScore float64
	ServiceScore float64
	ValueScore float64
	Reviews []Review `gorm:"foreignkey:TripAdvisorReviewReferId"`
}

func GetAllTripAdvisorReview()([]TripAdvisorReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	var tripAdvisorReviewList []TripAdvisorReview
	db.Find(&tripAdvisorReviewList)
	for i,_ := range tripAdvisorReviewList {
		db.Model(tripAdvisorReviewList[i]).Related(&tripAdvisorReviewList[i].Reviews, "TripAdvisorReviewReferId")
	}
	return tripAdvisorReviewList, err
}

func CreateTripAdvisorReview(averageScore float64, locationScore float64, cleanlinessScore float64, serviceStore float64, valueScore float64)(*TripAdvisorReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tripAdvisorReview := &TripAdvisorReview{
		AverageScore: averageScore,
		LocationScore: locationScore,
		CleanlinessScore: cleanlinessScore,
		ServiceScore: serviceStore,
		ValueScore: valueScore,
	}
	db.Save(&tripAdvisorReview)
	return tripAdvisorReview, err
}

func UpdateTripAdvisorReview(id int, averageScore float64, locationScore float64, cleanlinessScore float64, serviceStore float64, valueScore float64)(*TripAdvisorReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	var tripAdvisorReview TripAdvisorReview
	db.Model(&tripAdvisorReview).Where("id = ?", id).Update(map[string]interface{}{"avrage_score": averageScore, "location_score": locationScore, "cleanliness_score": cleanlinessScore, "service_score": serviceStore, "value_score": valueScore})
	db.Where("id = ?", id).Find(&tripAdvisorReview)
	return &tripAdvisorReview, nil
}

func RemoveTripAdvisorReview(id int) (*TripAdvisorReview, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	tripAdvisorReview := TripAdvisorReview{Id: id}
	db.Where("id = ?", id).Find(&tripAdvisorReview)
	return &tripAdvisorReview, db.Delete(tripAdvisorReview).Error
}
