package auth

import "auctionkuy.wildangbudhi.com/domain"

type AuthUsecase interface {
	Authenticate(appleToken, locale string) (string, string, error, domain.HTTPStatusCode)
}
