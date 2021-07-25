package httprest

import (
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

type refreshAccessTokenRequestQuery struct {
	RefreshToken *string `form:"refresh_token" json:"refresh_token" binding:"required"`
}

type refreshAccessTokenResponseBody struct {
	AccessToken string `json:"access_token"`
	Expiration  int64  `json:"expiration"`
}

func (handler *AuthHTTPRestHandler) RefreshAccessToken(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode

	ctx.Header("Content-Type", "application/json")

	requestQuery := &refreshAccessTokenRequestQuery{}

	err = ctx.BindQuery(requestQuery)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	var newAccessToken string
	var tokenExp int64

	newAccessToken, tokenExp, err, statusCode = handler.authUsecase.RefreshAccessToken(*requestQuery.RefreshToken)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), refreshAccessTokenResponseBody{AccessToken: newAccessToken, Expiration: tokenExp})

}
