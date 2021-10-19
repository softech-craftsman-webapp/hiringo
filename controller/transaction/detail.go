package transaction

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get Transaction Detail
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get transaction Details
// @Tags transaction
// @Description Get transaction Details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.TransactionView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /categories/{id} [get]
// @Security JWT
func GetTransactionDetail(ctx echo.Context) error {
	db := config.GetDB()

	transaction := model.Transaction{
		ID: ctx.Param("id"),
	}
	result := db.First(&transaction)

	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Transaction not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.TransactionView{
			ID:       transaction.ID,
			UserID:   transaction.UserID,
			Amount:   transaction.Amount,
			Currency: transaction.Currency,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
