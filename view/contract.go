package view

import (
	model "hiringo/model"
	"time"
)

type ContractView struct {
	ID                       string    `json:"id"`
	StartTime                time.Time `json:"start_time"`
	EndTime                  time.Time `json:"end_time"`
	SignedByRecruiterTime    time.Time `json:"signed_by_recruiter_time"`
	SignedByProfessionalTime time.Time `json:"signed_by_professional_time"`
	RecruiterID              string    `json:"recruiter_id"`
	ProfessionalID           string    `json:"professional_id"`
	JobID                    string    `json:"job_id"`
}

type ContractEmptyView struct {
	ID string `json:"id"`
}

func ContractModelToView(contract model.Contract) ContractView {
	return ContractView{
		ID:                       contract.ID,
		StartTime:                contract.StartTime,
		EndTime:                  contract.EndTime,
		SignedByRecruiterTime:    contract.SignedByRecruiterTime,
		SignedByProfessionalTime: contract.SignedByProfessionalTime,
		RecruiterID:              contract.RecruiterID,
		ProfessionalID:           contract.ProfessionalID,
		JobID:                    contract.JobID,
	}
}
