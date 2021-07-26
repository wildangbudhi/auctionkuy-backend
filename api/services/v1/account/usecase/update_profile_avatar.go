package usecase

import (
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
)

func (usecase *accountUsecase) UpdateProfileAvatar(authUserID *domain.UUID, data []byte, contentType string) (*domain.Image, error, domain.HTTPStatusCode) {

	log.Println(contentType)

	var err error

	var user *account.Users

	user, err, _ = usecase.usersRepository.GetUserByID(authUserID, usecase.serverConfig.ObjectURLBase)

	if err != nil {
		return nil, err, 500
	}

	if user.AvatarURL != nil {

		err, _ = usecase.userObjectRepository.RemoveUserAvatar(authUserID)

		if err != nil {
			return nil, err, 500
		}

	}

	err, _ = usecase.userObjectRepository.PutUserAvatar(authUserID, data, contentType)

	if err != nil {
		return nil, err, 500
	}

	var newAvatarURL *domain.Image

	newAvatarURL, err = domain.NewImage("account/profile/avater/"+user.ID.GetValue(), nil)

	if err != nil {
		return nil, err, 500
	}

	if user.AvatarURL == nil {
		user.AvatarURL = newAvatarURL

		err, _ = usecase.usersRepository.UpdateUser(user)

		if err != nil {
			return nil, err, 500
		}

	}

	newAvatarURL.SetPrefix(usecase.serverConfig.ObjectURLBase)

	return newAvatarURL, nil, 0

}
