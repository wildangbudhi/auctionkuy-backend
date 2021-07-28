package usecase

import (
	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) AddTransaction(authUserID *domain.UUID, transaction *transaction.Transactions) (*domain.UUID, error, domain.HTTPStatusCode) {

	var err error

	if transaction.Description != nil && len(*transaction.Description) > 300 {
		return nil, err, 400
	}

	var transactionID *domain.UUID = domain.NewUUID()
	var transactionStatusUD int = 1

	transaction.ID = transactionID
	transaction.SellerID = authUserID
	transaction.StatusID = &transactionStatusUD

	err, _ = usecase.transactionsRepository.CreateTransaction(transaction)

	if err != nil {
		return nil, err, 500
	}

	return transactionID, nil, 200

}
