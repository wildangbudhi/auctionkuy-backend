package httprest

import (
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"github.com/gin-gonic/gin"
)

type transactionHTTPRestHandler struct {
	transactionUsecase transaction.TransactionUsecase
	authMiddleware     auth.AuthMiddlewareDelivery
}

func NewTransactionHTTPRestHandler(router *gin.RouterGroup, transactionUsecase transaction.TransactionUsecase, authMiddleware auth.AuthMiddlewareDelivery) {

	handler := transactionHTTPRestHandler{
		transactionUsecase: transactionUsecase,
		authMiddleware:     authMiddleware,
	}

	router.POST("/create", handler.authMiddleware.ValidateAccessToken, handler.AddTransaction)
	router.GET("/:id", handler.authMiddleware.ValidateAccessToken, handler.GetTransaction)
}
