package view

type UserDetailView struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Email     string  `json:"email"`
	Bio       string  `json:"bio"`
	Telephone string  `json:"telephone"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UserDetailEmptyView struct {
	ID string `json:"id"`
}

type UserRatingDetailView struct {
	ID         string  `json:"id"`
	Rating     float64 `json:"rating"`
	TotalRates int     `json:"total_rates"`
}
