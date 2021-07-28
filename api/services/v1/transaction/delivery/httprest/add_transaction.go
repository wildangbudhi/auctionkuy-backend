package httprest

import (
	"log"
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"github.com/gin-gonic/gin"
)

type addTransactionRequesteBody struct {
	ItemName    *string  `json:"item_name"`
	ItemPrice   *float64 `json:"item_price"`
	Description *string  `json:"description"`
}

type addTransactionResponseBody struct {
	TransactionID *domain.UUID `json:"transaction_id"`
}

func (handler *transactionHTTPRestHandler) AddTransaction(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var authUserID *domain.UUID
	var transactionID *domain.UUID

	ctx.Header("Content-Type", "application/json")

	var authHeaderInterface interface{}
	var isAuthHeaderExists bool = false

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
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Unauthorized"})
		return
	}

	requestBodyData := &addTransactionRequesteBody{}

	err = ctx.BindJSON(requestBodyData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	var transactionData *transaction.Transactions = &transaction.Transactions{
		ItemName:    requestBodyData.ItemName,
		ItemPrice:   requestBodyData.ItemPrice,
		Description: requestBodyData.Description,
	}

	transactionID, err, statusCode = handler.transactionUsecase.AddTransaction(authUserID, transactionData)

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), addTransactionResponseBody{TransactionID: transactionID})

}
