package middleware

import "auctionkuy.wildangbudhi.com/domain/v1/auth"

type authMiddlewareDelivery struct {
	authUsecase auth.AuthUsecase
}

func NewAuthMiddlewareDelivery(authUsecase auth.AuthUsecase) auth.AuthMiddlewareDelivery {
	return &authMiddlewareDelivery{
		authUsecase: authUsecase,
	}
}
