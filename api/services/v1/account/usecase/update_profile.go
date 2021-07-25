package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
)

func (usecase *accountUsecase) UpdateProfile(authUserID *domain.UUID, name, phone, nationalIDNumber, locale, bankID, bankAccountID, bankAccountOwnerName *string) (*account.Users, error, domain.HTTPStatusCode) {

	var err error
	var user *account.Users

	user, err, _ = usecase.usersRepository.GetUserByID(authUserID)

	if err != nil {
		return nil, err, 500
	}

	if name != nil {
		user.Name = name
	}

	if nationalIDNumber != nil {
		user.NationalIDNumber = nationalIDNumber
	}

	if bankID != nil {

		var bankUUID *domain.UUID

		bankUUID, err = domain.NewUUIDFromString(*bankID)

		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("Bank ID Format Invalid"), 400
		}

		user.BankID = bankUUID
	}

	if bankAccountID != nil {
		user.BankAccountID = bankAccountID
	}

	if bankAccountOwnerName != nil {
		user.BankAccountOwnerName = bankAccountOwnerName
	}

	if locale != nil {

		var isLocaleExists bool

		_, isLocaleExists = usecase.serverConfig.CountryData.PhoneNumberMaps[*locale]

		if !isLocaleExists {
			return nil, fmt.Errorf("Locale invalid"), 400
		}

		user.Locale = locale

	}

	if phone != nil {

		var phoneNumber *domain.PhoneNumber

		phoneNumber, err = domain.NewPhoneNumber(*phone)

		if err != nil {

			err = phoneNumber.SanitizePhoneNumber(&usecase.serverConfig.CountryData, *user.Locale)

			if err != nil {
				return nil, err, 400
			}

		}

		user.Phone = phoneNumber
	}

	var firstTimeLogin bool = false
	user.FirstLogin = &firstTimeLogin

	err, _ = usecase.usersRepository.UpdateUser(user)

	if err != nil {
		return nil, err, 500
	}

	user, err, _ = usecase.usersRepository.GetUserByID(authUserID)

	if err != nil {
		return nil, err, 500
	}

	return user, nil, 200

}
