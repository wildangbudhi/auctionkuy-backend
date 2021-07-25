package account

import "auctionkuy.wildangbudhi.com/domain"

type Users struct {
	ID                   *domain.UUID        `json:"id"`
	Email                *domain.Email       `json:"email"`
	Name                 *string             `json:"name"`
	Phone                *domain.PhoneNumber `json:"phone"`
	NationalIDNumber     *string             `json:"national_id_number"`
	AvatarURL            *string             `json:"avatar_url"`
	FirstLogin           *bool               `json:"first_login"`
	Locale               *string             `json:"locale"`
	BankID               *string             `json:"bank_id"`
	BankAccountID        *string             `json:"bank_account_id"`
	BankAccountOwnerName *string             `json:"bank_account_owner_name"`
}

type UsersRepository interface {
	GetUserByID(id *domain.UUID) (*Users, error, domain.RepositoryErrorType)
}