package auth

import "auctionkuy.wildangbudhi.com/domain"

type AuthUsecase interface {
	Authenticate(appleToken, locale string) (string, string, int64, error, domain.HTTPStatusCode)
	RefreshAccessToken(refreshToken string) (string, int64, error, domain.HTTPStatusCode)
	ValidateAccessToken(token string) (*domain.UUID, error, domain.HTTPStatusCode)
	Logout(token string) (error, domain.HTTPStatusCode)
}
