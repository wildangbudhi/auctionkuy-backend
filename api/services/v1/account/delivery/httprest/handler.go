package httprest

import (
	"auctionkuy.wildangbudhi.com/domain/v1/account"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/gin-gonic/gin"
)

type accountHTTPRestHandler struct {
	accountUsecase account.AccountUsecase
	authMiddleware auth.AuthMiddlewareDelivery
}

func NewAccountHTTPRestHandler(router *gin.RouterGroup, accountUsecase account.AccountUsecase, authMiddleware auth.AuthMiddlewareDelivery) {

	handler := accountHTTPRestHandler{
		accountUsecase: accountUsecase,
		authMiddleware: authMiddleware,
	}

	router.GET("/profile", handler.authMiddleware.ValidateAccessToken, handler.Profile)
	router.PATCH("/profile", handler.authMiddleware.ValidateAccessToken, handler.UpdateProfile)
	router.PUT("/profile/avatar", handler.authMiddleware.ValidateAccessToken, handler.UpdateProfileAvatar)

}
