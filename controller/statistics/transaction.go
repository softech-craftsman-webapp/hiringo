package statistics

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
   | Transaction statistics
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Transaction statistics
// @Tags statistics
// @Description Transaction statistics
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.TransactionStatistics}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /statistics/transaction [get]
// @Security JWT
func TransactionStatistics(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()

	// get user transactions
	userTransactions := []model.Transaction{}
	result := db.Where("user_id = ?", claims.User.ID).Find(&userTransactions)

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

	// transactions
	transactions := []model.Transaction{}
	resultTransactions := db.Find(&transactions)

	if resultTransactions.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultTransactions.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	// get latest transaction
	latestTransaction := model.Transaction{
		UserID: claims.User.ID,
	}

	resultLatestTransaction := db.Where("user_id = ?", claims.User.ID).Order("created_at desc").First(&latestTransaction)
	if resultLatestTransaction.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultLatestTransaction.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	// result
	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.TransactionStatistics{
			LatestTransaction: &view.TransactionView{
				ID:       latestTransaction.ID,
				UserID:   latestTransaction.UserID,
				Amount:   latestTransaction.Amount,
				Currency: latestTransaction.Currency,
			},
			UserTransactionCount: len(userTransactions),
			Total:                len(transactions),
			Time:                 latestTransaction.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
