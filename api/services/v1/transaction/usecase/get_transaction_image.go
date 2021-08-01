package usecase

import (
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) GetTransactionImage(authUserID *domain.UUID, transactionID *domain.UUID, objectID *domain.UUID) ([]byte, string, error, domain.HTTPStatusCode) {

	var err error
	var object []byte
	var objectContentType string
	var repositoryErrorType domain.RepositoryErrorType
	var transactionData *transaction.Transactions

	transactionData, err, repositoryErrorType = usecase.transactionsRepository.GetTransactionByID(transactionID, "")

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, "", fmt.Errorf("Transaction Not Found"), 400
	}

	if err != nil {
		return nil, "", err, 500
	}

	var isSeller, isBuyer bool = false, false

	if transactionData.SellerID != nil && *transactionData.SellerID == *authUserID {
		isSeller = true
	}

	if transactionData.BuyerID != nil && *transactionData.BuyerID == *authUserID {
		isBuyer = true
	}

	if !isSeller && !isBuyer {
		return nil, "", fmt.Errorf("Transaction Data Not Found"), 400
	}

	object, objectContentType, err, repositoryErrorType = usecase.transactionsObjectRepository.GetTransactionObject(objectID)

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, "", fmt.Errorf("Transaction image not found"), 400
	}

	if err != nil {
		return nil, "", err, 500
	}

	return object, objectContentType, nil, 200

}
