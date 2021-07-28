package usecase

import (
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) GetTransaction(authUserID *domain.UUID, transactionID *domain.UUID) (*transaction.Transactions, error, domain.HTTPStatusCode) {

	var err error
	var repositoryErrorType domain.RepositoryErrorType
	var transactionData *transaction.Transactions
	var maxSellerStep, maxBuyerStep *int

	transactionData, err, repositoryErrorType = usecase.transactionsRepository.GetFullTransactionByID(transactionID, usecase.serverConfig.ObjectURLBase)

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, fmt.Errorf("Transaction Data Not Found"), 400
	}

	if err != nil {
		return nil, err, 500
	}

	if (transactionData.Seller != nil && transactionData.Seller.ID != authUserID) && (transactionData.Buyer != nil && transactionData.Buyer.ID != authUserID) {
		return nil, fmt.Errorf("Transaction Data Not Found"), 400
	}

	maxBuyerStep, maxSellerStep, err, _ = usecase.transactionStatusRepository.GetStepMax()

	if err != nil {
		return nil, err, 500
	}

	transactionData.Status.BuyerStepMax = maxBuyerStep
	transactionData.Status.SellerStepMax = maxSellerStep

	return transactionData, nil, 200

}
