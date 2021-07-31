package httpobject

import (
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"github.com/gin-gonic/gin"
)

type updateTransactionBodyFormBody struct {
	ItemPhoto           *multipart.FileHeader `form:"item_photo"`
	PackedItemImage     *multipart.FileHeader `form:"packed_item_image"`
	ReceivedItemImage   *multipart.FileHeader `form:"received_item_image"`
	PaymentReceiptImage *multipart.FileHeader `form:"payment_receipt_image"`
}

type updateTransactionResponseBody struct {
	*transaction.TransactionImages
}

func (handler *transactionHTTPObjectHandler) UpdateTransactionImages(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var authUserID *domain.UUID

	var authHeaderInterface interface{}
	var isAuthHeaderExists bool = false

	var transactionID string = ctx.Param("id")

	authHeaderInterface, isAuthHeaderExists = ctx.Get("AUTH_HEADER")

	if !isAuthHeaderExists {
		log.Println("Auth header not found")
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Unauthorized"})
		return
	}

	var isConversionOK bool = false

	authUserID, isConversionOK = authHeaderInterface.(*domain.UUID)

	if !isConversionOK {
		log.Println("Cannot convert interface{} to *domain.UUID")
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	requestForm := &updateTransactionBodyFormBody{}

	err = ctx.Bind(requestForm)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	var transactionUUID *domain.UUID

	transactionUUID, err = domain.NewUUIDFromString(transactionID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	var transactionImages *transaction.TransactionImages = &transaction.TransactionImages{}

	if requestForm.ItemPhoto != nil {

		transactionImages.ItemPhoto, err = extractFile(requestForm.ItemPhoto)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
			return
		}

	}

	if requestForm.PackedItemImage != nil {

		transactionImages.PackedItemImage, err = extractFile(requestForm.PackedItemImage)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
			return
		}

	}

	if requestForm.ReceivedItemImage != nil {

		transactionImages.ReceivedItemImage, err = extractFile(requestForm.ReceivedItemImage)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
			return
		}

	}

	if requestForm.PaymentReceiptImage != nil {

		transactionImages.PaymentReceiptImage, err = extractFile(requestForm.PaymentReceiptImage)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
			return
		}

	}

	var response *transaction.TransactionImages

	response, err, statusCode = handler.transactionUsecase.UpdateTransactionImages(authUserID, transactionUUID, transactionImages)

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), updateTransactionResponseBody{TransactionImages: response})

}

func extractFile(mulripartHeader *multipart.FileHeader) (*transaction.Files, error) {

	var err error
	var fileReader multipart.File

	fileReader, err = mulripartHeader.Open()

	if err != nil {
		return nil, err
	}

	var bytesFile []byte

	bytesFile, err = ioutil.ReadAll(fileReader)

	if err != nil {
		return nil, err
	}

	return &transaction.Files{Data: bytesFile, ContentType: mulripartHeader.Header.Get("Content-Type")}, nil
}
