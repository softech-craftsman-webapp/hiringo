package view

import "hiringo/model"

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

type ContractRatingsView struct {
	JobID            string       `json:"job_id"`
	ContractID       string       `json:"contract_id"`
	RecruiterID      string       `json:"recruiter_id"`
	ProfessionalID   string       `json:"professional_id"`
	IsRatingFinished bool         `json:"is_rating_finished"`
	Ratings          []RatingView `json:"ratings"`
}

func RatingModelToView(rating model.Rating) RatingView {
	return RatingView{
		ID:            rating.ID,
		UserID:        rating.UserID,
		SubmittedByID: rating.SubmittedByID,
		ContractID:    rating.ContractID,
		Points:        rating.Points,
		Comment:       rating.Comment,
	}
}
