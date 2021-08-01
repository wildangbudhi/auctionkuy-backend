package usecase

import (
	"fmt"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) ConfirmArrivalTransaction(authUserID *domain.UUID, transactionID *domain.UUID) (error, domain.HTTPStatusCode) {

	var err error
	var repositoryErrorType domain.RepositoryErrorType
	var transactionData *transaction.Transactions

	transactionData, err, repositoryErrorType = usecase.transactionsRepository.GetTransactionByID(transactionID, "")

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return fmt.Errorf("Transaction Not Found"), 400
	}

	if err != nil {
		return err, 500
	}

	if transactionData.BuyerID == nil {
		return fmt.Errorf("Transaction Not Available"), 400
	}

	if *transactionData.BuyerID != *authUserID {
		return fmt.Errorf("Transaction Not Available"), 400
	}

	var newTransactionStatusID int = 5
	transactionData.StatusID = &newTransactionStatusID

	err, _ = usecase.transactionsRepository.UpdateTransaction(transactionData)

	if err != nil {
		return err, 500
	}

	return nil, 200

}
