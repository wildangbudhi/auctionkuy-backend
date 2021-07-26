package usecase

import (
	"auctionkuy.wildangbudhi.com/domain/v1/account"
	"auctionkuy.wildangbudhi.com/utils"
)

type accountUsecase struct {
	serverConfig         *utils.Config
	usersRepository      account.UsersRepository
	userObjectRepository account.UserObjectRepository
}

func NewAccountUsecase(serverConfig *utils.Config, usersRepository account.UsersRepository, userObjectRepository account.UserObjectRepository) account.AccountUsecase {
	return &accountUsecase{
		serverConfig:         serverConfig,
		usersRepository:      usersRepository,
		userObjectRepository: userObjectRepository,
	}
}
