package transaction

import (
	"encoding/json"

	"auctionkuy.wildangbudhi.com/domain"
)

type Users struct {
	ID        *domain.UUID        `json:"id"`
	Name      *string             `json:"name"`
	Phone     *domain.PhoneNumber `json:"phone"`
	AvatarURL *domain.Image       `json:"avatar_url"`
}

func (obj *Users) MarshalJSON() ([]byte, error) {

	if (Users{ID: obj.ID} == *obj) {
		return json.Marshal(obj.ID)
	}

	return json.Marshal(obj)

}
