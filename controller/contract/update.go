package contract

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UpdateContractRequest struct {
	StartTime string `json:"start_time" validate:"required"`
	EndTime   string `json:"end_time" validate:"required"`
}

/*
   |--------------------------------------------------------------------------
   | Update Contract
   | @JWT via Acess Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Update Contract
// @Tags contract
// @Description Update Contract
// @Accept  json
// @Produce  json
// @Param id path string true "Contract id"
// @Param user body UpdateContractRequest true "Contract details"
// @Success 200 {object} view.Response{payload=view.ContractEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/{id} [put]
// @Security JWT
func UpdateContractDetail(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(UpdateContractRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: config.GetMessageFromError(err.Error()),
			Payload: nil,
		})
	}

	// Time Validation Start Time
	startTime, err := time.Parse("2006-01-02 15:04", req.StartTime)
	if err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "StartTime validation failed, example 2006-01-02 15:04",
			Payload: nil,
		})
	}

	endTime, err := time.Parse("2006-01-02 15:04", req.EndTime)
	if err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "EndTime validation failed, example 2006-01-02 15:04",
			Payload: nil,
		})
	}

	contract := &model.Contract{
		ID: ctx.Param("id"),
	}

	result := db.First(&contract, "id =? AND professional_id = ?", contract.ID, claims.User.ID)

	/*
	   |--------------------------------------------------------------------------
	   | Main Error
	   |--------------------------------------------------------------------------
	*/
	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: result.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	/*
	   |--------------------------------------------------------------------------
	   | Check rquired fields
	   |--------------------------------------------------------------------------
	*/
	resultContractUpdate := db.Model(&contract).Updates(model.Contract{
		StartTime: startTime,
		EndTime:   endTime,
	})

	if resultContractUpdate.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: resultContractUpdate.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.CategoryEmptyView{
			ID: contract.ID,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
