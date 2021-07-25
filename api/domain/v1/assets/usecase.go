package assets

import "auctionkuy.wildangbudhi.com/domain"

type AssetsUsecase interface {
	AppAssets() (*AppAssets, error, domain.HTTPStatusCode)
}
