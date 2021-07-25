package v1

import (
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/services/v1/assets/delivery/httprest"
	"auctionkuy.wildangbudhi.com/services/v1/assets/repository/mysql"
	"auctionkuy.wildangbudhi.com/services/v1/assets/usecase"
	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func AssetsHTTPRestDI(server *utils.Server) {

	var route *gin.RouterGroup = server.Router.Group("/v1/assets")

	var appSettingsRepository assets.AppSettingsRepository
	var banksRepository assets.BanksRepository

	var assetsUsecase assets.AssetsUsecase
	var authMiddleware auth.AuthMiddlewareDelivery

	appSettingsRepository = mysql.NewAppSettingsRepository(server.DB)
	banksRepository = mysql.NewBanksRepository(server.DB)

	assetsUsecase = usecase.NewAssetsUsecase(&server.Config, appSettingsRepository, banksRepository)
	authMiddleware = AuthMiddlewareDI(server)

	httprest.NewAssetsHTTPRestHandler(route, assetsUsecase, authMiddleware)
}
