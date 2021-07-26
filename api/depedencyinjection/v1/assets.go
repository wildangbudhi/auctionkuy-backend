package v1

import (
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/services/v1/assets/delivery/httpobject"
	"auctionkuy.wildangbudhi.com/services/v1/assets/delivery/httprest"
	"auctionkuy.wildangbudhi.com/services/v1/assets/repository/minioobject"
	"auctionkuy.wildangbudhi.com/services/v1/assets/repository/mysql"
	"auctionkuy.wildangbudhi.com/services/v1/assets/usecase"
	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func getAssetsUsecase(server *utils.Server) assets.AssetsUsecase {

	var appSettingsRepository assets.AppSettingsRepository
	var banksRepository assets.BanksRepository
	var banksObjectRepository assets.BanksObjectRepository

	var assetsUsecase assets.AssetsUsecase

	appSettingsRepository = mysql.NewAppSettingsRepository(server.DB)
	banksRepository = mysql.NewBanksRepository(server.DB)
	banksObjectRepository = minioobject.NewBanksObjectRepository(server.ObjectStorage)

	assetsUsecase = usecase.NewAssetsUsecase(&server.Config, appSettingsRepository, banksRepository, banksObjectRepository)

	return assetsUsecase

}

func AssetsHTTPDI(server *utils.Server) {

	var route *gin.RouterGroup = server.Router.Group("/v1/assets")

	var assetsUsecase assets.AssetsUsecase
	var authMiddleware auth.AuthMiddlewareDelivery

	assetsUsecase = getAssetsUsecase(server)
	authMiddleware = AuthMiddlewareDI(server)

	httprest.NewAssetsHTTPRestHandler(route, assetsUsecase, authMiddleware)
	httpobject.NewAssetHTTPObjectHandler(route, assetsUsecase, authMiddleware)
}
