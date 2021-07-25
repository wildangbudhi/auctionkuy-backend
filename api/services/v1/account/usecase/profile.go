package usecase

import (
	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
)

func (usecase *accountUsecase) Profile(authUserID *domain.UUID) (*account.Users, error, domain.HTTPStatusCode) {

	var err error
	var user *account.Users

	user, err, _ = usecase.usersRepository.GetUserByID(authUserID)

	if err != nil {
		return nil, err, 500
	}

	return user, nil, 200

}
