package transaction

import "auctionkuy.wildangbudhi.com/domain"

type Transaction struct {
	ID                     *domain.UUID       `json:"id"`
	Status                 *TransactionStatus `json:"status"`
	ItemPhotoURL           *domain.Image      `json:"item_photo_url"`
	ItemName               *string            `json:"item_name"`
	ItemPrice              *float64           `json:"item_price"`
	Seller                 *Users             `json:"seller"`
	Buyer                  *Users             `json:"buyer"`
	ShippingCourier        *string            `json:"shipping_courier"`
	ShippingReceiptID      *string            `json:"shipping_receipt_id"`
	PackedItemImageURL     *domain.Image      `json:"packed_item_image_url"`
	RecievedItemImageURL   *domain.Image      `json:"recieved_item_image_url"`
	PaymentReceiptImageURL *domain.Image      `json:"payment_receipt_image_url"`
	CreatedAt              *domain.Timestamp  `json:"created_at"`
	UpdatedAt              *domain.Timestamp  `json:"updated_at"`
}

type TransactionRepository interface {
	GetFullTransactionByID(id *domain.UUID)
}
