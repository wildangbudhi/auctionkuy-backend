package account

import "auctionkuy.wildangbudhi.com/domain"

type AccountUsecase interface {
	Profile(authUserID *domain.UUID) (*Users, error, domain.HTTPStatusCode)
}
