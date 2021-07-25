package assets

import "auctionkuy.wildangbudhi.com/domain"

type Banks struct {
	ID      *domain.UUID  `json:"id"`
	Name    *string       `json:"name"`
	IconURL *domain.Image `json:"icon_url"`
}

type BanksRepository interface {
	FetchBanks() ([]Banks, error, domain.RepositoryErrorType)
}
