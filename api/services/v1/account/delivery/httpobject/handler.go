package httpobject

import (
	"auctionkuy.wildangbudhi.com/domain/v1/account"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/gin-gonic/gin"
)

type accountHTTPObjectHandler struct {
	accountUsecase account.AccountUsecase
	authMiddleware auth.AuthMiddlewareDelivery
}

func NewAccountHTTPObjectHandler(router *gin.RouterGroup, accountUsecase account.AccountUsecase, authMiddleware auth.AuthMiddlewareDelivery) {

	handler := accountHTTPObjectHandler{
		accountUsecase: accountUsecase,
		authMiddleware: authMiddleware,
	}

	router.GET("/profile/avatar/:userid", handler.authMiddleware.ValidateAccessToken, handler.GetProfileAvatar)
	router.PUT("/profile/avatar", handler.authMiddleware.ValidateAccessToken, handler.UpdateProfileAvatar)

}
