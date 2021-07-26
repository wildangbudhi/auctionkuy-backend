package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
)

func (usecase *assetsUsecase) BankObject(objectName string) ([]byte, string, error, domain.HTTPStatusCode) {

	var err error
	var object []byte
	var objectContentType string

	if objectName == "" {
		return nil, "", fmt.Errorf("Object name invalid"), 400
	}

	objectName += ".png"

	log.Println(objectName)

	object, objectContentType, err = usecase.banksObjectRepository.GetBanksLogo(objectName)

	if err != nil {
		return nil, "", err, 500
	}

	return object, objectContentType, nil, 200

}
