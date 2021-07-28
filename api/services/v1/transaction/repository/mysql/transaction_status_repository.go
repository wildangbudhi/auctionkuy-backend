package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

type transactionStatusRepository struct {
	db *sql.DB
}

func NewTransactionStatusRepository(db *sql.DB) transaction.TransactionStatusRepository {
	return &transactionStatusRepository{
		db: db,
	}
}

func (repo *transactionStatusRepository) GetStepMax() (*int, *int, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT 
		MAX(ts.buyer_step) AS buyer_max_step,
		MAX(ts.seller_step) AS seller_max_step
	FROM 
		transaction_status ts 
	`

	var queryResult *sql.Row
	var maxBuyerStep, maxSellerStep int

	queryResult = repo.db.QueryRow(queryString)

	err = queryResult.Scan(
		&maxBuyerStep,
		&maxSellerStep,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, nil, err, domain.RepositoryDataNotFound
		}

		log.Println(err)
		return nil, nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	return &maxBuyerStep, &maxSellerStep, nil, 0

}
