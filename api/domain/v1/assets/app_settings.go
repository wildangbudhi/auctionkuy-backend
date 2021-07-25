package assets

import "auctionkuy.wildangbudhi.com/domain"

type AppSettings struct {
	Key   *string
	Value *string
}

type AppSettingsRepository interface {
	FetchAppSettings() ([]AppSettings, error, domain.RepositoryErrorType)
}
