package usecase

import (
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
)

func (usecase *accountUsecase) GetProfileAvatar(userID *domain.UUID) ([]byte, string, error, domain.HTTPStatusCode) {

	var err error
	var object []byte
	var objectContentType string
	var repositoryErrorType domain.RepositoryErrorType

	object, objectContentType, err, repositoryErrorType = usecase.userObjectRepository.GetUserAvatar(userID)

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, "", fmt.Errorf("User avatar not found"), 400
	}

	if err != nil {
		return nil, "", err, 500
	}

	return object, objectContentType, nil, 200

}
