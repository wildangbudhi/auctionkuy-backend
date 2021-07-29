package usecase

import (
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) JoinTransaction(authUserID *domain.UUID, transactionID *domain.UUID) (*domain.UUID, error, domain.HTTPStatusCode) {

	var err error
	var repositoryErrorType domain.RepositoryErrorType
	var transactionData *transaction.Transactions

	transactionData, err, repositoryErrorType = usecase.transactionsRepository.GetTransactionByID(transactionID, usecase.serverConfig.ObjectURLBase)

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, fmt.Errorf("Transaction Not Found"), 400
	}

	if err != nil {
		return nil, err, 500
	}

	if transactionData.BuyerID != nil {
		return nil, fmt.Errorf("Transaction Not Available"), 400
	}

	var newTransactionStatusID int = 2

	transactionData.BuyerID = authUserID
	transactionData.StatusID = &newTransactionStatusID

	err, _ = usecase.transactionsRepository.UpdateTransaction(transactionData)

	if err != nil {
		return nil, err, 500
	}

	return transactionID, nil, 200

}
