package usecase

import (
	"fmt"
	"log"
	"strings"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
)

func (usecase *transactionUsecase) UpdateTransactionImages(authUserID *domain.UUID, transactionID *domain.UUID, images *transaction.TransactionImages) (*transaction.TransactionImages, error, domain.HTTPStatusCode) {

	var err error
	var repositoryErrorType domain.RepositoryErrorType
	var transactionData *transaction.Transactions
	var response *transaction.TransactionImages = &transaction.TransactionImages{}

	transactionData, err, repositoryErrorType = usecase.transactionsRepository.GetTransactionByID(transactionID, usecase.serverConfig.ObjectURLBase)

	if repositoryErrorType == domain.RepositoryDataNotFound {
		return nil, fmt.Errorf("Transaction Data Not Found"), 400
	}

	if err != nil {
		return nil, err, 500
	}

	var isSeller, isBuyer bool = true, true

	if transactionData.Seller != nil && transactionData.Seller.ID != authUserID {
		isSeller = false
	}

	if transactionData.Buyer != nil && transactionData.Buyer.ID != authUserID {
		isBuyer = false
	}

	if isSeller || isBuyer {
		return nil, fmt.Errorf("Transaction Data Not Found"), 400
	}

	if images.ItemPhoto != nil {

		if transactionData.ItemPhotoURL != nil {

			var objectURLSplitted []string = strings.Split(transactionData.ItemPhotoURL.GetValue(), "/")
			var objectID string = objectURLSplitted[len(objectURLSplitted)-1]
			var objectUUID *domain.UUID

			objectUUID, err = domain.NewUUIDFromString(objectID)

			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("Object ID Invalid"), 500
			}

			err, _ = usecase.transactionsObjectRepository.RemoveUserTransactionObject(objectUUID)

			if err != nil {
				return nil, err, 500
			}

		}

		var newObjectUUID *domain.UUID = domain.NewUUID()

		err, _ = usecase.transactionsObjectRepository.PutTransactionObject(newObjectUUID, images.ItemPhoto.Data, images.ItemPhoto.ContentType)

		if err != nil {
			return nil, err, 500
		}

		var newImageURL *domain.Image

		newImageURL, err = domain.NewImage("transaction/image/"+newObjectUUID.GetValue(), nil)

		if err != nil {
			return nil, err, 400
		}

		transactionData.ItemPhotoURL = newImageURL

	}

	if images.PackedItemImage != nil {

		if transactionData.PackedItemImageURL != nil {

			var objectURLSplitted []string = strings.Split(transactionData.PackedItemImageURL.GetValue(), "/")
			var objectID string = objectURLSplitted[len(objectURLSplitted)-1]
			var objectUUID *domain.UUID

			objectUUID, err = domain.NewUUIDFromString(objectID)

			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("Object ID Invalid"), 500
			}

			err, _ = usecase.transactionsObjectRepository.RemoveUserTransactionObject(objectUUID)

			if err != nil {
				return nil, err, 500
			}

		}

		var newObjectUUID *domain.UUID = domain.NewUUID()

		err, _ = usecase.transactionsObjectRepository.PutTransactionObject(newObjectUUID, images.PackedItemImage.Data, images.PackedItemImage.ContentType)

		if err != nil {
			return nil, err, 500
		}

		var newImageURL *domain.Image

		newImageURL, err = domain.NewImage("transaction/image/"+newObjectUUID.GetValue(), nil)

		if err != nil {
			return nil, err, 400
		}

		transactionData.PackedItemImageURL = newImageURL

	}

	if images.ReceivedItemImage != nil {

		if transactionData.RecievedItemImageURL != nil {

			var objectURLSplitted []string = strings.Split(transactionData.RecievedItemImageURL.GetValue(), "/")
			var objectID string = objectURLSplitted[len(objectURLSplitted)-1]
			var objectUUID *domain.UUID

			objectUUID, err = domain.NewUUIDFromString(objectID)

			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("Object ID Invalid"), 500
			}

			err, _ = usecase.transactionsObjectRepository.RemoveUserTransactionObject(objectUUID)

			if err != nil {
				return nil, err, 500
			}

		}

		var newObjectUUID *domain.UUID = domain.NewUUID()

		err, _ = usecase.transactionsObjectRepository.PutTransactionObject(newObjectUUID, images.ReceivedItemImage.Data, images.ReceivedItemImage.ContentType)

		if err != nil {
			return nil, err, 500
		}

		var newImageURL *domain.Image

		newImageURL, err = domain.NewImage("transaction/image/"+newObjectUUID.GetValue(), nil)

		if err != nil {
			return nil, err, 400
		}

		transactionData.RecievedItemImageURL = newImageURL

	}

	if images.PaymentReceiptImage != nil {

		if transactionData.PaymentReceiptImageURL != nil {

			var objectURLSplitted []string = strings.Split(transactionData.PaymentReceiptImageURL.GetValue(), "/")
			var objectID string = objectURLSplitted[len(objectURLSplitted)-1]
			var objectUUID *domain.UUID

			objectUUID, err = domain.NewUUIDFromString(objectID)

			if err != nil {
				log.Println(err)
				return nil, fmt.Errorf("Object ID Invalid"), 500
			}

			err, _ = usecase.transactionsObjectRepository.RemoveUserTransactionObject(objectUUID)

			if err != nil {
				return nil, err, 500
			}

		}

		var newObjectUUID *domain.UUID = domain.NewUUID()

		err, _ = usecase.transactionsObjectRepository.PutTransactionObject(newObjectUUID, images.PaymentReceiptImage.Data, images.PaymentReceiptImage.ContentType)

		if err != nil {
			return nil, err, 500
		}

		var newImageURL *domain.Image

		newImageURL, err = domain.NewImage("transaction/image/"+newObjectUUID.GetValue(), nil)

		if err != nil {
			return nil, err, 400
		}

		transactionData.PaymentReceiptImageURL = newImageURL

	}

	err, _ = usecase.transactionsRepository.UpdateTransaction(transactionData)

	if err != nil {
		return nil, err, 500
	}

	response.ItemPhotoURL = transactionData.ItemPhotoURL

	if response.ItemPhotoURL != nil {
		response.ItemPhotoURL.SetPrefix(usecase.serverConfig.ObjectURLBase)
	}

	response.PackedItemImageURL = transactionData.PackedItemImageURL

	if response.PackedItemImageURL != nil {
		response.PackedItemImageURL.SetPrefix(usecase.serverConfig.ObjectURLBase)
	}

	response.ReceivedItemImageURL = transactionData.RecievedItemImageURL

	if response.ReceivedItemImageURL != nil {
		response.ReceivedItemImageURL.SetPrefix(usecase.serverConfig.ObjectURLBase)
	}

	response.PaymentReceiptImageURL = transactionData.PaymentReceiptImageURL

	if response.PaymentReceiptImageURL != nil {
		response.PaymentReceiptImageURL.SetPrefix(usecase.serverConfig.ObjectURLBase)
	}

	return response, nil, 200

}
