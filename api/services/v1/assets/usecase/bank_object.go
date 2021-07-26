package usecase

import (
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
)

func (usecase *assetsUsecase) BankObject(objectName string) ([]byte, string, error, domain.HTTPStatusCode) {

	var err error
	var object []byte
	var objectContentType string
	var repositoryErrorType domain.RepositoryErrorType

	if objectName == "" {
		return nil, "", fmt.Errorf("Object name invalid"), 400
	}

	objectName += ".png"

	object, objectContentType, err, repositoryErrorType = usecase.banksObjectRepository.GetBanksLogo(objectName)

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, "", fmt.Errorf("Banks icon not found"), 400
	}

	if err != nil {
		return nil, "", err, 500
	}

	return object, objectContentType, nil, 200

}
