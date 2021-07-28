package usecase

import (
	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) FetchTransaction(authUserID *domain.UUID) ([]transaction.TransactionsThumbnail, error, domain.HTTPStatusCode) {

	var err error
	var transactionList []transaction.TransactionsThumbnail

	transactionList, err, _ = usecase.transactionsRepository.FetchTransactions(authUserID, usecase.serverConfig.ObjectURLBase)

	if err != nil {
		return nil, err, 500
	}

	return transactionList, nil, 200

}
