package account

import "auctionkuy.wildangbudhi.com/domain"

type AccountUsecase interface {
	Profile(authUserID *domain.UUID) (*Users, error, domain.HTTPStatusCode)
	UpdateProfile(authUserID *domain.UUID, name, phone, nationalIDNumber, locale, bankID, bankAccountID, bankAccountOwnerName *string) (*Users, error, domain.HTTPStatusCode)
	UpdateProfileAvatar(authUserID *domain.UUID, data []byte, contentType string) (*domain.Image, error, domain.HTTPStatusCode)
}
