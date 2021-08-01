package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) ConfimrPaymentTransaction(authUserID *domain.UUID, transactionID *domain.UUID, paymentMethodID *domain.UUID, paymentAccount *string) (error, domain.HTTPStatusCode) {

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

	log.Println(transactionData.BuyerID)

	if transactionData.BuyerID == nil {
		return fmt.Errorf("Transaction Not Available"), 400
	}

	if *transactionData.BuyerID != *authUserID {
		return fmt.Errorf("Transaction Not Available"), 400
	}

	var newTransactionStatusID int = 3

	transactionData.BuyerPaymentMethodID = paymentMethodID
	transactionData.BuyerPaymentAccount = paymentAccount
	transactionData.StatusID = &newTransactionStatusID

	err, _ = usecase.transactionsRepository.UpdateTransaction(transactionData)

	if err != nil {
		return err, 500
	}

	return nil, 200

}
