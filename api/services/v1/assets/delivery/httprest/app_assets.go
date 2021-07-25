package httprest

import (
	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"github.com/gin-gonic/gin"
)

type appAssetsResponseBody struct {
	*assets.AppAssets
}

func (handler *assetsHTTPRestHandler) AppAssets(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var appAssets *assets.AppAssets

	ctx.Header("Content-Type", "application/json")

	appAssets, err, statusCode = handler.assetsUsecase.AppAssets()

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), appAssetsResponseBody{AppAssets: appAssets})

}
