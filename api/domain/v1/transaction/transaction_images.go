package transaction

import "auctionkuy.wildangbudhi.com/domain"

type TransactionImages struct {
	ItemPhoto              *Files        `json:"item_photo,omitempty"`
	ItemPhotoURL           *domain.Image `json:"item_photo_url"`
	PackedItemImage        *Files        `json:"packed_item_image,omitempty"`
	PackedItemImageURL     *domain.Image `json:"packed_item_image_url"`
	ReceivedItemImage      *Files        `json:"received_item_image,omitempty"`
	ReceivedItemImageURL   *domain.Image `json:"received_item_image_url"`
	PaymentReceiptImage    *Files        `json:"payment_receipt_image,omitempty"`
	PaymentReceiptImageURL *domain.Image `json:"payment_receipt_image_url"`
}
