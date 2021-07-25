package usecase

import (
	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
)

func (usecase *assetsUsecase) AppAssets() (*assets.AppAssets, error, domain.HTTPStatusCode) {

	var err error

	var appSettings []assets.AppSettings
	var supportedBanks []assets.Banks

	appSettings, err, _ = usecase.appSettingsRepository.FetchAppSettings()

	if err != nil {
		return nil, err, 500
	}

	supportedBanks, err, _ = usecase.banksRepository.FetchBanks()

	if err != nil {
		return nil, err, 500
	}

	for i := 0; i < len(supportedBanks); i++ {
		if supportedBanks[i].IconURL != nil {
			supportedBanks[i].IconURL.SetPrefix(usecase.serverConfig.ObjectURLBase)
		}
	}

	var appAssets *assets.AppAssets = &assets.AppAssets{
		SupportedBanks: supportedBanks,
	}

	for i := 0; i < len(appSettings); i++ {

		if appSettings[i].Key == nil || appSettings[i].Value == nil {
			continue
		}

		if *appSettings[i].Key == "company_bank_id" {
			appAssets.CompanyBankID = appSettings[i].Value
		}

		if *appSettings[i].Key == "company_bank_account_id" {
			appAssets.CompanyBankAccountID = appSettings[i].Value
		}

		if *appSettings[i].Key == "company_bank_account_owner_name" {
			appAssets.CompanyBankAccountOwnerName = appSettings[i].Value
		}

	}

	return appAssets, nil, 200

}
