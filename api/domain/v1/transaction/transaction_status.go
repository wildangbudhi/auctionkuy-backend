package transaction

import (
	"auctionkuy.wildangbudhi.com/domain"
)

type TransactionStatusRepository interface {
	GetStepMax() (*int, *int, error, domain.RepositoryErrorType)
}

type TransactionStatus struct {
	ID            *int    `json:"id"`
	SellerCommand *string `json:"seller_command"`
	BuyerCommand  *string `json:"buyer_command"`
	SellerStep    *int    `json:"seller_step"`
	BuyerStep     *int    `json:"buyer_step"`
	SellerStepMax *int    `json:"seller_step_max"`
	BuyerStepMax  *int    `json:"buyer_step_max"`
}
