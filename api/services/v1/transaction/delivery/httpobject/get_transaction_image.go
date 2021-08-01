package httpobject

import (
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

func (handler *transactionHTTPObjectHandler) GetTransactionImage(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var authUserID *domain.UUID
	var object []byte
	var objectContentType string

	var transactionID string = ctx.Param("id")
	var objectID string = ctx.Param("object-id")

	var authHeaderInterface interface{}
	var isAuthHeaderExists bool = false

	authHeaderInterface, isAuthHeaderExists = ctx.Get("AUTH_HEADER")

	if !isAuthHeaderExists {
		log.Println("Auth header not found")
		ctx.String(int(statusCode), "%s", "Unauthorized")
		return
	}

	var isConversionOK bool = false

	authUserID, isConversionOK = authHeaderInterface.(*domain.UUID)

	if !isConversionOK {
		log.Println("Cannot convert interface{} to *domain.UUID")
		ctx.String(int(statusCode), "%s", "Unauthorized")
		return
	}

	var transactionUUID, objectUUID *domain.UUID

	transactionUUID, err = domain.NewUUIDFromString(transactionID)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	objectUUID, err = domain.NewUUIDFromString(objectID)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	object, objectContentType, err, statusCode = handler.transactionUsecase.GetTransactionImage(
		authUserID,
		transactionUUID,
		objectUUID,
	)

	ctx.Data(int(statusCode), objectContentType, object)

}
