package assets

import (
	"auctionkuy.wildangbudhi.com/domain"
	"github.com/minio/minio-go/v7"
)

type Banks struct {
	ID      *domain.UUID  `json:"id"`
	Name    *string       `json:"name"`
	IconURL *domain.Image `json:"icon_url"`
}

type BanksRepository interface {
	FetchBanks() ([]Banks, error, domain.RepositoryErrorType)
}

type BanksObjectRepository interface {
	GetBanksLogo(objectName string) (*minio.Object, error)
}
