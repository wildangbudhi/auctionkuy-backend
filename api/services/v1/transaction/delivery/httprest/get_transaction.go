package httprest

import (
	"log"
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"github.com/gin-gonic/gin"
)

type getTransactionResponseBody struct {
	*transaction.Transactions
}

func (handler *transactionHTTPRestHandler) GetTransaction(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var authUserID *domain.UUID
	var transactionData *transaction.Transactions

	var transactionID string = ctx.Param("id")

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

	var transactionUUID *domain.UUID

	transactionUUID, err = domain.NewUUIDFromString(transactionID)

	if err != nil {
		ctx.JSON(400, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	transactionData, err, statusCode = handler.transactionUsecase.GetTransaction(authUserID, transactionUUID)

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), getTransactionResponseBody{Transactions: transactionData})

}
