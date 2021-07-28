package transaction

import (
	"auctionkuy.wildangbudhi.com/domain"
)

type Users struct {
	ID        *domain.UUID        `json:"id"`
	Name      *string             `json:"name"`
	Phone     *domain.PhoneNumber `json:"phone"`
	AvatarURL *domain.Image       `json:"avatar_url"`
}
