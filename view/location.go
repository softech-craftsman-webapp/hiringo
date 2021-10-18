package view

import "hiringo/model"

type LocationView struct {
	ID            string  `json:"id"`
	UserID        string  `json:"user_id"`
	Address       string  `json:"address"`
	Street        string  `json:"street"`
	HouseNumber   string  `json:"house_number"`
	Suburb        string  `json:"suburb"`
	Postcode      string  `json:"postcode"`
	State         string  `json:"state"`
	StateCode     string  `json:"state_code"`
	StateDistrict string  `json:"state_district"`
	County        string  `json:"county"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"country_code"`
	City          string  `json:"city"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type LocationEmptyView struct {
	ID string `json:"id"`
}

func LocationModelToView(location model.Location) LocationView {
	return LocationView{
		ID:            location.ID,
		UserID:        location.UserID,
		Address:       location.Address,
		Street:        location.Street,
		HouseNumber:   location.HouseNumber,
		Suburb:        location.Suburb,
		Postcode:      location.Postcode,
		State:         location.State,
		StateCode:     location.StateCode,
		StateDistrict: location.StateDistrict,
		County:        location.County,
		Country:       location.Country,
		CountryCode:   location.CountryCode,
		City:          location.City,
		Latitude:      location.Latitude,
		Longitude:     location.Longitude,
	}
}
