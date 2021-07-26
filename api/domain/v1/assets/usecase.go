package assets

import (
	"auctionkuy.wildangbudhi.com/domain"
	"github.com/minio/minio-go/v7"
)

type AssetsUsecase interface {
	AppAssets() (*AppAssets, error, domain.HTTPStatusCode)
	BankObject(objectName string) (*minio.Object, *minio.ObjectInfo, error, domain.HTTPStatusCode)
}
