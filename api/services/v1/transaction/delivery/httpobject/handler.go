package httpobject

import (
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"github.com/gin-gonic/gin"
)

type transactionHTTPObjectHandler struct {
	transactionUsecase transaction.TransactionUsecase
	authMiddleware     auth.AuthMiddlewareDelivery
}

func NewTransactionHTTPObjectHandler(router *gin.RouterGroup, transactionUsecase transaction.TransactionUsecase, authMiddleware auth.AuthMiddlewareDelivery) {

	handler := transactionHTTPObjectHandler{
		transactionUsecase: transactionUsecase,
		authMiddleware:     authMiddleware,
	}

	router.PATCH("/:id/image", handler.authMiddleware.ValidateAccessToken, handler.UpdateTransactionImages)
}
