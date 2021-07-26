package assets

import (
	"auctionkuy.wildangbudhi.com/domain"
)

type AssetsUsecase interface {
	AppAssets() (*AppAssets, error, domain.HTTPStatusCode)
	BankObject(objectName string) ([]byte, string, error, domain.HTTPStatusCode)
}
