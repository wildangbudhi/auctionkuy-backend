package transaction

import "auctionkuy.wildangbudhi.com/domain"

type Transactions struct {
	ID                     *domain.UUID       `json:"id"`
	StatusID               *int               `json:"status_id,omitempty"`
	Status                 *TransactionStatus `json:"status,omitempty"`
	ItemPhotoURL           *domain.Image      `json:"item_photo_url"`
	ItemName               *string            `json:"item_name"`
	Description            *string            `json:"description"`
	ItemPrice              *float64           `json:"item_price"`
	SellerID               *domain.UUID       `json:"seller_id,omitempty"`
	Seller                 *Users             `json:"seller"`
	BuyerID                *domain.UUID       `json:"buyer_id,omitempty"`
	Buyer                  *Users             `json:"buyer"`
	ShippingCourier        *string            `json:"shipping_courier"`
	ShippingReceiptID      *string            `json:"shipping_receipt_id"`
	PackedItemImageURL     *domain.Image      `json:"packed_item_image_url"`
	RecievedItemImageURL   *domain.Image      `json:"recieved_item_image_url"`
	BuyerPaymentMethodID   *domain.UUID       `json:"buyer_payment_method_id,omitempty"`
	BuyerPaymentMethod     *Banks             `json:"buyer_payment_method,omitempty"`
	BuyerPaymentAccount    *string            `json:"buyer_payment_account,omitempty"`
	PaymentReceiptImageURL *domain.Image      `json:"payment_receipt_image_url"`
	CreatedAt              *domain.Timestamp  `json:"created_at"`
	UpdatedAt              *domain.Timestamp  `json:"updated_at"`
}

type TransactionsThumbnail struct {
	ID           *domain.UUID       `json:"id"`
	Status       *TransactionStatus `json:"status,omitempty"`
	ItemPhotoURL *domain.Image      `json:"item_photo_url"`
	ItemName     *string            `json:"item_name"`
	SellerID     *domain.UUID       `json:"seller_id"`
	BuyerID      *domain.UUID       `json:"buyer_id"`
	CreatedAt    *domain.Timestamp  `json:"created_at"`
	UpdatedAt    *domain.Timestamp  `json:"updated_at"`
}

type TransactionsRepository interface {
	GetFullTransactionByID(id *domain.UUID, imagePrefix string) (*Transactions, error, domain.RepositoryErrorType)
	GetTransactionByID(id *domain.UUID, imagePrefix string) (*Transactions, error, domain.RepositoryErrorType)
	FetchTransactions(userID *domain.UUID, imagePrefix string) ([]TransactionsThumbnail, error, domain.RepositoryErrorType)
	CreateTransaction(transaction *Transactions) (error, domain.RepositoryErrorType)
	UpdateTransaction(transaction *Transactions) (error, domain.RepositoryErrorType)
}

type TransactionsObjectRepository interface {
	GetTransactionObject(objectID *domain.UUID) ([]byte, string, error, domain.RepositoryErrorType)
	PutTransactionObject(objectID *domain.UUID, data []byte, contentType string) (error, domain.RepositoryErrorType)
	RemoveUserTransactionObject(objectID *domain.UUID) (error, domain.RepositoryErrorType)
}
