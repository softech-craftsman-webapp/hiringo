package view

type CategoryStatistics struct {
	PopularItem *CategoryView `json:"popular_item"`
	Total       int           `json:"total"`
	Time        string        `json:"time"`
}

type JobStatistics struct {
	Total        int      `json:"total"`
	UserJobCount int      `json:"user_job_count"`
	PopularItem  *JobView `json:"latest_job"`
	Time         string   `json:"time"`
}

type RatingStatistics struct {
	Total           int         `json:"total"`
	UserRatingCount int         `json:"user_rating_count"`
	LatestRating    *RatingView `json:"latest_rating"`
	Time            string      `json:"time"`
}

type TransactionStatistics struct {
	Total                int              `json:"total"`
	UserTransactionCount int              `json:"user_transaction_count"`
	LatestTransaction    *TransactionView `json:"latest_transaction"`
	Time                 string           `json:"time"`
}
