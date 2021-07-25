package httprest

import (
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

type authtenticateRequestBody struct {
	AppleToken *string `json:"apple_token" binding:"required"`
	Locale     *string `json:"locale" binding:"required"`
}

type authtenticateResponseBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expiration   int64  `json:"expiration"`
}

func (handler *AuthHTTPRestHandler) Authenticate(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode

	ctx.Header("Content-Type", "application/json")

	requestBodyData := &authtenticateRequestBody{}

	err = ctx.BindJSON(requestBodyData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	var aksesToken, refreshToken string
	var tokenExp int64

	aksesToken, refreshToken, tokenExp, err, statusCode = handler.authUsecase.Authenticate(
		*requestBodyData.AppleToken,
		*requestBodyData.Locale,
	)

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), authtenticateResponseBody{AccessToken: aksesToken, RefreshToken: refreshToken, Expiration: tokenExp})

}
