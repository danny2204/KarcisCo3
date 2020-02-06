package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Insert struct {
	City      string  `json:"city"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	ZIndex    int
	Province  string `json:"admin"`
	Country   string `json:"country"`
}

type LocationT struct {
	Locs []Insert `json:"locations"`
}

func InsertLocationData() {
	byteValue, err := ioutil.ReadFile("models/data/id.json")
	if err != nil {
		fmt.Println(err)
	}
	var locData LocationT
	err = json.Unmarshal(byteValue, &locData)
	//fmt.Println(er)
	for i, _ := range locData.Locs {
		loc := &Location{
			Longitude: locData.Locs[i].Longitude,
			Latitude:  locData.Locs[i].Latitude,
			City:      locData.Locs[i].City,
			Province:  locData.Locs[i].Province,
			Country:   locData.Locs[i].Country,
			ZIndex:    12,
		}
		//fmt.Println(loc.Longitude, loc.Latitude)
		db.Save(&loc)
	}
}
