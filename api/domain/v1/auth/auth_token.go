package auth

import "auctionkuy.wildangbudhi.com/domain"

type AuthToken interface {
	GenerateAuthToken(userID *domain.UUID) (*domain.JWT, *domain.JWT, error)
	ValidateToken(token *domain.JWT, isRefreshToken bool) error
	RegenerateAuthToken(refreshToken *domain.JWT) (*domain.JWT, error)
	RemoveAuthToken(token *domain.JWT) error
}
