package httprest

import (
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/gin-gonic/gin"
)

type AuthHTTPRestHandler struct {
	authUsecase auth.AuthUsecase
}

func NewAuthHTTPRestHandler(router *gin.RouterGroup, authUsecase auth.AuthUsecase) {

	handler := AuthHTTPRestHandler{
		authUsecase: authUsecase,
	}

	router.POST("/login", handler.Authenticate)
	router.GET("/refresh", handler.RefreshAccessToken)
	router.GET("/logout", handler.Logout)

}
