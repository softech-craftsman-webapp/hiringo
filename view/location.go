package view

type LocationView struct {
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
