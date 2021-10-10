package view

type UserDetailView struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	Email      string `json:"email"`
	Bio        string `json:"bio"`
	LocationID string `json:"location_id"`
	Telephone  string `json:"telephone"`
}

type UserDetailEmptyView struct {
	ID string `json:"id"`
}
