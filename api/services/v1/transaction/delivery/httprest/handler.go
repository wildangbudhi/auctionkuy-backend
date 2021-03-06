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

	router.GET("/", handler.authMiddleware.ValidateAccessToken, handler.FetchTransaction)
	router.GET("/:id", handler.authMiddleware.ValidateAccessToken, handler.GetTransaction)
	router.POST("/create", handler.authMiddleware.ValidateAccessToken, handler.AddTransaction)
	router.GET("/:id/join", handler.authMiddleware.ValidateAccessToken, handler.JoinTransaction)
	router.POST("/:id/confirm/payment", handler.authMiddleware.ValidateAccessToken, handler.ConfirmPaymentTransaction)
	router.POST("/:id/confirm/shipping", handler.authMiddleware.ValidateAccessToken, handler.ConfirmShippingTransaction)
	router.GET("/:id/confirm/arrival", handler.authMiddleware.ValidateAccessToken, handler.ConfirmArrivalTransaction)
	router.GET("/:id/confirm/withdrawal", handler.authMiddleware.ValidateAccessToken, handler.ConfirmWithdrawalTransaction)
}
