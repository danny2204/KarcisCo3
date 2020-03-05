package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
)

type PromoDetail struct {
	Id int `gorm:"primary_key"`
	DiscountAmount int
	KodePromo string
	PromoReferId int
}

func GetAllPromoDetail()([]PromoDetail, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var promoDetails []PromoDetail
	db.Find(&promoDetails)
	return promoDetails, err
}

func CreatePromoDetail(discountAmount int, kodePromo string, promoReferId int)(*PromoDetail, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	promoDetail := &PromoDetail {
		DiscountAmount: discountAmount,
		KodePromo: kodePromo,
		PromoReferId: promoReferId,
	}
	db.Save(promoDetail)
	return promoDetail, err
}

func GetPromoDetailById(id int)(*PromoDetail, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var promoDetail PromoDetail
	db.Where("id = ?", id).Find(&promoDetail)
	return &promoDetail, err
}
