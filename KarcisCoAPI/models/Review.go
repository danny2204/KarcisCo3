package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Review struct {
	Id int `gorm:"primary_key"`
	Score float64
	Name string
	UploadDate time.Time
	Conclusion string `gorm:varchar(100); not null`
	ReviewString string `gorm:varchar(500);not null`
	TiketComReviewReferId int
	TripAdvisorReviewReferId int
}

func GetAllReview()([]Review, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var reviewList []Review
	db.Find(&reviewList)
	return reviewList, err
}

func CreateReview(score float64, uploadDate time.Time, conclusion string, reviewString string, tiketComReviewReferId int, tripAdvisorReviewReferId int, name string)(*Review, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	review := &Review {
		Score: score,
		UploadDate: uploadDate,
		Conclusion: conclusion,
		ReviewString: reviewString,
		TiketComReviewReferId: tiketComReviewReferId,
		TripAdvisorReviewReferId: tripAdvisorReviewReferId,
		Name: name,
	}
	db.Save(review)
	return review, err
}

func UpdateReview(id int, score float64, uploadDate time.Time, conclusion string, reviewString string, tiketComReviewReferId int, tripAdvisorReviewReferId int, name string)(*Review, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var review Review
	db.Model(&review).Where("id = >", id).Update(map[string]interface{}{"score": score, "upload_date": uploadDate, "conclusion": conclusion, "review_string": reviewString, "tiket_com_review_refer_id": tiketComReviewReferId, "trip_advisor_review_refer_id": tripAdvisorReviewReferId, "name": name})
	db.Where("id = ?", id).Find(&review)
	return &review, nil
}

func RemoveReview(id int)(*Review, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	review := Review{Id: id}
	db.Where("id = ?", id).Find(&review)
	return &review, err
}
