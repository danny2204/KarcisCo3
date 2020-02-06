package models

import (
	"time"
)

type Vendor struct{
	Id			int			`gorm:primary_key`
	Region		string		`gorm:"type:varchar(100)"`
	Name		string		`gorm:"type:varchar(100)"`
	//RentLocations

	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

