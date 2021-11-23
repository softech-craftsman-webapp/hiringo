package transaction

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get transactions
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get user details
// @Tags transaction
// @Description Get user details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=[]view.TransactionView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /transactions/my [get]
// @Security JWT
func GetMyTransactions(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	transactions := []model.Transaction{}
	result := db.Where("user_id = ?", claims.User.ID).Find(&transactions).Order("created_at")

	if result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: result.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	// TODO: It can be optimized
	var formatted_transactions []view.TransactionView
	for _, transaction := range transactions {
		formatted_transactions = append(formatted_transactions, view.TransactionModelToView(transaction))
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: formatted_transactions,
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
