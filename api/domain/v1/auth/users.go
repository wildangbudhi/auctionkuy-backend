package auth

import "auctionkuy.wildangbudhi.com/domain"

type Users struct {
	ID         *domain.UUID
	Email      *domain.Email
	Locale     *string
	FirstLogin *bool
}

type UsersRepository interface {
	GetUserByEmail(email *domain.Email) (*Users, error, domain.RepositoryErrorType)
	CreateUser(user *Users) (error, domain.RepositoryErrorType)
	UpdateUser(user *Users) (error, domain.RepositoryErrorType)
}
