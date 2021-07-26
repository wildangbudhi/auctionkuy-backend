package httpobject

import (
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/gin-gonic/gin"
)

type assetsHTTPObjectHandler struct {
	assetsUsecase  assets.AssetsUsecase
	authMiddleware auth.AuthMiddlewareDelivery
}

func NewAssetHTTPObjectHandler(router *gin.RouterGroup, assetsUsecase assets.AssetsUsecase, authMiddleware auth.AuthMiddlewareDelivery) {

	handler := assetsHTTPObjectHandler{
		assetsUsecase:  assetsUsecase,
		authMiddleware: authMiddleware,
	}

	router.GET("/banks/:object-name", handler.BankObject)

}
