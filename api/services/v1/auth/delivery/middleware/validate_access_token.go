package middleware

import (
	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

type authRequestHeader struct {
	Authorization string `header:"Authorization" json:"Authorization" binding:"required"`
}

func (handler *authMiddlewareDelivery) ValidateAccessToken(ctx *gin.Context) {

	var err error
	var userID string

	requestHeader := &authRequestHeader{}

	err = ctx.BindHeader(requestHeader)

	if err != nil {
		ctx.AbortWithStatusJSON(401, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	userID, err, _ = handler.authUsecase.ValidateAccessToken(requestHeader.Authorization)

	if err != nil {
		ctx.AbortWithStatusJSON(401, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.Set("AUTH_HEADER", userID)

	ctx.Next()

}
