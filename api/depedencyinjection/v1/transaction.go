package v1

import (
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/domain/v1/transaction"
	"auctionkuy.wildangbudhi.com/services/v1/transaction/delivery/httpobject"
	"auctionkuy.wildangbudhi.com/services/v1/transaction/delivery/httprest"
	"auctionkuy.wildangbudhi.com/services/v1/transaction/repository/minioobject"
	"auctionkuy.wildangbudhi.com/services/v1/transaction/repository/mysql"
	"auctionkuy.wildangbudhi.com/services/v1/transaction/usecase"
	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func TransactionHTTPDI(server *utils.Server) {

	var route *gin.RouterGroup = server.Router.Group("/v1/transaction")

	var transactionUsecase transaction.TransactionUsecase
	var authMiddleware auth.AuthMiddlewareDelivery

	var transactionRepository transaction.TransactionsRepository
	var transactionStatusRepository transaction.TransactionStatusRepository
	var transactionsObjectRepository transaction.TransactionsObjectRepository

	transactionRepository = mysql.NewTransactionsRepository(server.DB)
	transactionStatusRepository = mysql.NewTransactionStatusRepository(server.DB)
	transactionsObjectRepository = minioobject.NewTransactionsObjectRepository(server.ObjectStorage)

	transactionUsecase = usecase.NewTransactionUsecase(
		&server.Config,
		transactionRepository,
		transactionStatusRepository,
		transactionsObjectRepository,
	)
	authMiddleware = AuthMiddlewareDI(server)

	httprest.NewTransactionHTTPRestHandler(route, transactionUsecase, authMiddleware)
	httpobject.NewTransactionHTTPObjectHandler(route, transactionUsecase, authMiddleware)

}
