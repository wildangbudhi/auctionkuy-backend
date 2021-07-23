package usecase

import (
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/utils"
)

type authUsecase struct {
	serverConfig        *utils.Config
	appleKeysRepository auth.AppleKeysRepository
	usersRepository     auth.UsersRepository
	sessionRepository   auth.SessionRepository
}

func NewAuthUsecase(
	serverConfig *utils.Config,
	appleKeysRepository auth.AppleKeysRepository,
	usersRepository auth.UsersRepository,
	sessionRepository auth.SessionRepository,
) auth.AuthUsecase {
	return &authUsecase{
		serverConfig:        serverConfig,
		appleKeysRepository: appleKeysRepository,
		usersRepository:     usersRepository,
		sessionRepository:   sessionRepository,
	}
}
