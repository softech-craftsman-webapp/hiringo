package view

type LocationView struct {
	PlaceID     int    `json:"place_id"`
	OsmType     string `json:"osm_type"`
	OsmID       int    `json:"osm_id"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	DisplayName string `json:"display_name"`
	Address     struct {
		HouseNumber  string `json:"house_number"`
		Road         string `json:"road"`
		Suburb       string `json:"suburb"`
		CityDistrict string `json:"city_district"`
		Region       string `json:"region"`
		Postcode     string `json:"postcode"`
		Country      string `json:"country"`
		CountryCode  string `json:"country_code"`
	} `json:"address"`
	Boundingbox []string `json:"boundingbox"`
}

type CoordinatesView struct {
	Ip  string  `json:"ip"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
