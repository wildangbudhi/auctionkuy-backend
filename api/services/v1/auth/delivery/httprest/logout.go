package httprest

import (
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

type logoutRequestQuery struct {
	AccessToken *string `form:"access_token" json:"access_token" binding:"required"`
}

type logoutResponseBody struct {
	Status string `json:"status"`
}

func (handler *AuthHTTPRestHandler) Logout(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode

	ctx.Header("Content-Type", "application/json")

	requestQuery := &logoutRequestQuery{}

	err = ctx.BindQuery(requestQuery)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err, statusCode = handler.authUsecase.Logout(*requestQuery.AccessToken)

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), logoutResponseBody{Status: "Success"})

}
