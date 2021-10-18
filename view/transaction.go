package view

import "hiringo/model"

type TransactionView struct {
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type TransactionEmptyView struct {
	ID string `json:"id"`
}

func TransactionModelToView(transaction model.Transaction) TransactionView {
	return TransactionView{
		ID:       transaction.ID,
		UserID:   transaction.UserID,
		Amount:   transaction.Amount,
		Currency: transaction.Currency,
	}
}
