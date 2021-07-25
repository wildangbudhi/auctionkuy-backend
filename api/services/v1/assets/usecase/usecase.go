package usecase

import (
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"auctionkuy.wildangbudhi.com/utils"
)

type assetsUsecase struct {
	serverConfig          *utils.Config
	appSettingsRepository assets.AppSettingsRepository
	banksRepository       assets.BanksRepository
}

func NewAssetsUsecase(serverConfig *utils.Config, appSettingsRepository assets.AppSettingsRepository, banksRepository assets.BanksRepository) assets.AssetsUsecase {
	return &assetsUsecase{
		serverConfig:          serverConfig,
		appSettingsRepository: appSettingsRepository,
		banksRepository:       banksRepository,
	}
}
