package transaction

import "auctionkuy.wildangbudhi.com/domain"

type TransactionUsecase interface {
	FetchTransaction(authUserID *domain.UUID) ([]TransactionsThumbnail, error, domain.HTTPStatusCode)
	GetTransaction(authUserID *domain.UUID, transactionID *domain.UUID) (*Transactions, error, domain.HTTPStatusCode)
	AddTransaction(authUserID *domain.UUID, transaction *Transactions) (*domain.UUID, error, domain.HTTPStatusCode)
	JoinTransaction(authUserID *domain.UUID, transactionID *domain.UUID) (*domain.UUID, error, domain.HTTPStatusCode)
	UpdateTransactionImages(authUserID *domain.UUID, transactionID *domain.UUID, images *TransactionImages) (*TransactionImages, error, domain.HTTPStatusCode)
}
