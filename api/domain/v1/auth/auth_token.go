package auth

import "auctionkuy.wildangbudhi.com/domain"

type AuthToken interface {
	GenerateAuthToken(userID *domain.UUID) (*domain.JWT, *domain.JWT, *domain.Timestamp, error)
	ValidateToken(token *domain.JWT, isRefreshToken bool) error
	RegenerateAuthToken(refreshToken *domain.JWT) (*domain.JWT, *domain.Timestamp, error)
	RemoveAuthToken(token *domain.JWT) error
}
