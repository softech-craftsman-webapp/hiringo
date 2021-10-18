package transaction

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CreateTransactionRequest struct {
	Amount   float64 `json:"amount" validate:"required,min=0"`
	Currency string  `json:"currency" validate:"required,len=3"`
}

/*
   |--------------------------------------------------------------------------
   | Create Transaction
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Transaction
// @Tags transaction
// @Description Create Transaction
// @Accept  json
// @Produce  json
// @Param user body CreateTransactionRequest true "Amount and Currency"
// @Success 200 {object} view.Response{payload=view.TransactionView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /transactions/new [get]
// @Security JWT
func CreateTransaction(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateTransactionRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: "Bad request",
			Payload: nil,
		})
	}

	transaction := &model.Transaction{
		UserID:   claims.User.ID,
		Amount:   req.Amount,
		Currency: req.Currency,
	}

	result := db.Create(&transaction)
	/*
	   |--------------------------------------------------------------------------
	   | DB relation error
	   |--------------------------------------------------------------------------
	*/
	if result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: result.Error.Error(),
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusCreated, ctx, resp)
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
