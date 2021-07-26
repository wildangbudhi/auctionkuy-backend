package usecase

import (
	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
)

func (usecase *accountUsecase) Profile(authUserID *domain.UUID) (*account.Users, error, domain.HTTPStatusCode) {

	var err error
	var user *account.Users

	user, err, _ = usecase.usersRepository.GetUserByID(authUserID, usecase.serverConfig.ObjectURLBase)

	if err != nil {
		return nil, err, 500
	}

	if user.AvatarURL != nil {
		user.AvatarURL.SetPrefix(usecase.serverConfig.ObjectURLBase)
	}

	return user, nil, 200

}
