package httprest

import (
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/gin-gonic/gin"
)

type assetsHTTPRestHandler struct {
	assetsUsecase  assets.AssetsUsecase
	authMiddleware auth.AuthMiddlewareDelivery
}

func NewAssetsHTTPRestHandler(router *gin.RouterGroup, assetsUsecase assets.AssetsUsecase, authMiddleware auth.AuthMiddlewareDelivery) {

	handler := assetsHTTPRestHandler{
		assetsUsecase:  assetsUsecase,
		authMiddleware: authMiddleware,
	}

	router.GET("/app", handler.authMiddleware.ValidateAccessToken, handler.AppAssets)

}
