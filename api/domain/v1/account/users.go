package account

import "auctionkuy.wildangbudhi.com/domain"

type Users struct {
	ID                   *domain.UUID        `json:"id"`
	Email                *domain.Email       `json:"email"`
	Name                 *string             `json:"name"`
	Phone                *domain.PhoneNumber `json:"phone"`
	NationalIDNumber     *string             `json:"national_id_number"`
	AvatarURL            *domain.Image       `json:"avatar_url"`
	FirstLogin           *bool               `json:"first_login"`
	Locale               *string             `json:"locale"`
	BankID               *domain.UUID        `json:"bank_id"`
	BankAccountID        *string             `json:"bank_account_id"`
	BankAccountOwnerName *string             `json:"bank_account_owner_name"`
}

type UsersRepository interface {
	GetUserByID(id *domain.UUID, imagePrefix string) (*Users, error, domain.RepositoryErrorType)
	UpdateUser(user *Users) (error, domain.RepositoryErrorType)
}

type UserObjectRepository interface {
	GetUserAvatar(userID *domain.UUID) ([]byte, string, error, domain.RepositoryErrorType)
	PutUserAvatar(userID *domain.UUID, data []byte, contentType string) (error, domain.RepositoryErrorType)
	RemoveUserAvatar(userID *domain.UUID) (error, domain.RepositoryErrorType)
}
