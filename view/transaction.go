package view

type TransactionView struct {
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type TransactionEmptyView struct {
	ID string `json:"id"`
}
