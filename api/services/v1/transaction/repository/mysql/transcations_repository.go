package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

type transactionsRepository struct {
	db *sql.DB
}

func NewTransactionsRepository(db *sql.DB) transaction.TransactionsRepository {
	return &transactionsRepository{
		db: db,
	}
}

func (repo *transactionsRepository) GetFullTransactionByID(id *domain.UUID, imagePrefix string) (*transaction.Transactions, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT 
		t.transcation_status_id AS status_id,
		ts.seller_command AS transaction_status_seller_command,
		ts.buyer_command AS transaction_status_buyer_command,
		ts.seller_step AS transaction_status_seller_step,
		ts.buyer_step AS transaction_status_buyer_step,
		t.item_photo_url,
		t.item_name,
		t.description,
		t.item_price,
		us.id AS user_seller_id,
		us.name AS user_seller_name,
		us.phone AS user_seller_phone,
		us.avatar_url AS user_seller_avatar_url,
		ub.id AS user_buyer_id,
		ub.name AS user_buyer_name,
		ub.phone AS user_buyer_phone,
		ub.avatar_url AS user_buyer_avatar_url,
		t.shipping_courier,
		t.shipping_receipt_id,
		t.packed_item_image_url,
		t.recieved_item_image_url,
		b.id AS buyer_payment_method_bank_id,
		b.name AS buyer_payment_method_bank_name,
		b.icon_url AS buyer_payment_method_bank_icon_url,
		t.buyer_payment_account,
		t.payment_receipt_image_url,
		t.created_at,
		t.updated_at
	FROM 
		transactions t 
	LEFT JOIN
		transaction_status ts 
		ON ts.id = t.transcation_status_id 
	LEFT JOIN 
		users us
		ON us.id = t.seller_id 
	LEFT JOIN 
		users ub
		ON ub.id = t.buyer_id 
	LEFT JOIN
		banks b 
		ON b.id = t.buyer_payment_method_id 
	WHERE 
		t.id = ?
	`

	var queryResult *sql.Row
	var transaction *transaction.Transactions = &transaction.Transactions{
		ID:                 id,
		Status:             &transaction.TransactionStatus{},
		Seller:             &transaction.Users{},
		Buyer:              &transaction.Users{},
		BuyerPaymentMethod: &transaction.Banks{},
	}

	queryResult = repo.db.QueryRow(queryString, id)

	err = queryResult.Scan(
		&transaction.Status.ID,
		&transaction.Status.SellerCommand,
		&transaction.Status.BuyerCommand,
		&transaction.Status.SellerStep,
		&transaction.Status.BuyerStep,
		&transaction.ItemPhotoURL,
		&transaction.ItemName,
		&transaction.Description,
		&transaction.ItemPrice,
		&transaction.Seller.ID,
		&transaction.Seller.Name,
		&transaction.Seller.Phone,
		&transaction.Seller.AvatarURL,
		&transaction.Buyer.ID,
		&transaction.Buyer.Name,
		&transaction.Buyer.Phone,
		&transaction.Buyer.AvatarURL,
		&transaction.ShippingCourier,
		&transaction.ShippingReceiptID,
		&transaction.PackedItemImageURL,
		&transaction.RecievedItemImageURL,
		&transaction.BuyerPaymentMethod.ID,
		&transaction.BuyerPaymentMethod.Name,
		&transaction.BuyerPaymentMethod.IconURL,
		&transaction.BuyerPaymentAccount,
		&transaction.PaymentReceiptImageURL,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err, domain.RepositoryDataNotFound
		}

		log.Println(err)
		return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	if transaction.ItemPhotoURL != nil {
		transaction.ItemPhotoURL.SetPrefix(imagePrefix)
	}

	if transaction.Seller.AvatarURL != nil {
		transaction.Seller.AvatarURL.SetPrefix(imagePrefix)
	}

	if transaction.Buyer.AvatarURL != nil {
		transaction.Buyer.AvatarURL.SetPrefix(imagePrefix)
	}

	if transaction.PackedItemImageURL != nil {
		transaction.PackedItemImageURL.SetPrefix(imagePrefix)
	}

	if transaction.RecievedItemImageURL != nil {
		transaction.RecievedItemImageURL.SetPrefix(imagePrefix)
	}

	if transaction.BuyerPaymentMethod.IconURL != nil {
		transaction.BuyerPaymentMethod.IconURL.SetPrefix(imagePrefix)
	}

	if transaction.PaymentReceiptImageURL != nil {
		transaction.PaymentReceiptImageURL.SetPrefix(imagePrefix)
	}

	if transaction.Status.ID == nil {
		transaction.Status = nil
	}

	if transaction.Seller.ID == nil {
		transaction.Seller = nil
	}

	if transaction.Buyer.ID == nil {
		transaction.Buyer = nil
	}

	if transaction.BuyerPaymentMethod.ID == nil {
		transaction.BuyerPaymentMethod = nil
	}

	return transaction, nil, 0

}

func (repo *transactionsRepository) GetTransactionByID(id *domain.UUID, imagePrefix string) (*transaction.Transactions, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT 
		t.transcation_status_id AS status_id,
		t.item_photo_url,
		t.item_name,
		t.description,
		t.item_price,
		t.seller_id,
		t.buyer_id,
		t.shipping_courier,
		t.shipping_receipt_id,
		t.packed_item_image_url,
		t.recieved_item_image_url,
		t.buyer_payment_method_id,
		t.buyer_payment_account,
		t.payment_receipt_image_url,
		t.created_at,
		t.updated_at
	FROM 
		transactions t 
	WHERE 
		t.id = ?
	`

	var queryResult *sql.Row
	var transaction *transaction.Transactions = &transaction.Transactions{
		ID:                 id,
		Status:             &transaction.TransactionStatus{},
		Seller:             &transaction.Users{},
		Buyer:              &transaction.Users{},
		BuyerPaymentMethod: &transaction.Banks{},
	}

	queryResult = repo.db.QueryRow(queryString, id)

	err = queryResult.Scan(
		&transaction.StatusID,
		&transaction.ItemPhotoURL,
		&transaction.ItemName,
		&transaction.Description,
		&transaction.ItemPrice,
		&transaction.SellerID,
		&transaction.BuyerID,
		&transaction.ShippingCourier,
		&transaction.ShippingReceiptID,
		&transaction.PackedItemImageURL,
		&transaction.RecievedItemImageURL,
		&transaction.BuyerPaymentMethodID,
		&transaction.BuyerPaymentAccount,
		&transaction.PaymentReceiptImageURL,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, err, domain.RepositoryDataNotFound
		}

		log.Println(err)
		return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	if transaction.ItemPhotoURL != nil {
		transaction.ItemPhotoURL.SetPrefix(imagePrefix)
	}

	if transaction.PackedItemImageURL != nil {
		transaction.PackedItemImageURL.SetPrefix(imagePrefix)
	}

	if transaction.RecievedItemImageURL != nil {
		transaction.RecievedItemImageURL.SetPrefix(imagePrefix)
	}

	if transaction.PaymentReceiptImageURL != nil {
		transaction.PaymentReceiptImageURL.SetPrefix(imagePrefix)
	}

	return transaction, nil, 0

}

func (repo *transactionsRepository) CreateTransaction(transaction *transaction.Transactions) (error, domain.RepositoryErrorType) {

	var err error
	var queryString string

	queryString = `
	INSERT INTO transactions (
		id,
		transcation_status_id,
		item_photo_url,
		item_name,
		description,
		item_price,
		seller_id,
		buyer_id,
		shipping_courier,
		shipping_receipt_id,
		packed_item_image_url,
		recieved_item_image_url,
		buyer_payment_method_id,
		buyer_payment_account,
		payment_receipt_image_url
	) 
	VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	var res sql.Result

	res, err = statement.Exec(
		transaction.ID,
		transaction.StatusID,
		transaction.ItemPhotoURL,
		transaction.ItemName,
		transaction.Description,
		transaction.ItemPrice,
		transaction.SellerID,
		transaction.BuyerID,
		transaction.ShippingCourier,
		transaction.ShippingReceiptID,
		transaction.PackedItemImageURL,
		transaction.RecievedItemImageURL,
		transaction.BuyerPaymentMethodID,
		transaction.BuyerPaymentAccount,
		transaction.PaymentReceiptImageURL,
	)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	if rowAffected == 0 {
		return fmt.Errorf("Failed to Save User Data"), domain.RepositoryCreateDataFailed
	}

	return nil, 0

}

func (repo *transactionsRepository) FetchTransactions(userID *domain.UUID, imagePrefix string) ([]transaction.TransactionsThumbnail, error, domain.RepositoryErrorType) {

	var err error
	var queryString string = `
	SELECT 
		t.id,
		t.transcation_status_id AS status_id,
		ts.seller_command AS transaction_status_seller_command,
		ts.buyer_command AS transaction_status_buyer_command,
		ts.seller_step AS transaction_status_seller_step,
		ts.buyer_step AS transaction_status_buyer_step,
		t.item_photo_url,
		t.item_name,
		t.seller_id,
		t.buyer_id,
		t.created_at,
		t.updated_at
	FROM 
		transactions t 
	LEFT JOIN
		transaction_status ts 
		ON ts.id = t.transcation_status_id 
	WHERE 
		t.seller_id = ?
		OR 
		t.buyer_id = ?
	`

	var queryResult *sql.Rows
	var transactionList []transaction.TransactionsThumbnail = make([]transaction.TransactionsThumbnail, 0)

	queryResult, err = repo.db.Query(queryString, userID, userID)

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
	}

	defer queryResult.Close()

	for queryResult.Next() {

		var transactionData transaction.TransactionsThumbnail = transaction.TransactionsThumbnail{
			Status: &transaction.TransactionStatus{},
		}

		err = queryResult.Scan(
			&transactionData.ID,
			&transactionData.Status.ID,
			&transactionData.Status.SellerCommand,
			&transactionData.Status.BuyerCommand,
			&transactionData.Status.SellerStep,
			&transactionData.Status.BuyerStep,
			&transactionData.ItemPhotoURL,
			&transactionData.ItemName,
			&transactionData.SellerID,
			&transactionData.BuyerID,
			&transactionData.CreatedAt,
			&transactionData.UpdatedAt,
		)

		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("Services Unavailable"), domain.RepositoryError
		}

		if transactionData.ItemPhotoURL != nil {
			transactionData.ItemPhotoURL.SetPrefix(imagePrefix)
		}

		transactionList = append(transactionList, transactionData)

	}

	return transactionList, nil, 0

}
