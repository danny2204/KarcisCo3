package models

import (
	"github.com/danny2204/KarcisCoAPI/connection"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"time"
)

type Promo struct {
	Id int `gorm:"primary_key"`
	Name string
	PromoStart time.Time
	PromoEnd time.Time
	Platform string
	PromoDetail string
	PromoDetailList []PromoDetail `gorm:"foreign_key:PromoReferId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:index`
}

func GetAllPromo()([]Promo, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var promoList []Promo
	db.Find(&promoList)

	for i,_ := range promoList {
		db.Model(promoList[i]).Related(&promoList[i].PromoDetailList, "PromoReferId")
	}

	return promoList, err
}

func CreatePromo(name string, promoStart time.Time, promoEnd time.Time, platform string, promoDetail string)(*Promo, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	promo := &Promo{
		Name: name,
		PromoStart: promoStart,
		PromoEnd: promoEnd,
		Platform: platform,
		PromoDetail: promoDetail,
	}
	db.Save(promo)
	return promo, err
}

func GetPromoById(id int)(*Promo, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var promo Promo
	db.Where("id = ?", id).Find(&promo)
	db.Model(promo).Related(&promo.PromoDetailList, "PromoReferId")
	return &promo, err
}

func GetOtherPromo(id int)([]Promo, error) {
	db, err = connection.ConnectDatabase()
	_, err = GetApiKeyDetail(middleware.ApiKey)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var promoList []Promo
	db.Where("id != ?", id).Limit(3).Find(&promoList)
	for i,_ := range promoList {
		db.Model(promoList[i]).Related(&promoList[i].PromoDetailList, "PromoReferId")
	}
	return promoList, err
}

