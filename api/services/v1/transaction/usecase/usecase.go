package usecase

import (
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"auctionkuy.wildangbudhi.com/utils"
)

type transactionUsecase struct {
	serverConfig                *utils.Config
	transactionsRepository      transaction.TransactionsRepository
	transactionStatusRepository transaction.TransactionStatusRepository
}

func NewTransactionUsecase(serverConfig *utils.Config, transactionsRepository transaction.TransactionsRepository, transactionStatusRepository transaction.TransactionStatusRepository) transaction.TransactionUsecase {
	return &transactionUsecase{
		serverConfig:                serverConfig,
		transactionsRepository:      transactionsRepository,
		transactionStatusRepository: transactionStatusRepository,
	}
}
