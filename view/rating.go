package view

type RatingView struct {
	ID            string `json:"id"`
	SubmittedByID string `json:"submitted_by_id"`
	ContractID    string `json:"contract_id"`
	UserID        string `json:"user_id"`
	Points        int    `json:"points"`
	Comment       string `json:"comment"`
}

type RatingEmptyView struct {
	ID string `json:"id"`
}
