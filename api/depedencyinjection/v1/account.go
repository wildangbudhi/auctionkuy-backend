package v1

import (
	"auctionkuy.wildangbudhi.com/domain/v1/account"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/services/v1/account/delivery/httprest"
	"auctionkuy.wildangbudhi.com/services/v1/account/repository/mysql"
	"auctionkuy.wildangbudhi.com/services/v1/account/usecase"
	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func AccountHTTPRestDI(server *utils.Server) {

	var route *gin.RouterGroup = server.Router.Group("/v1/account")

	var accountUsecase account.AccountUsecase
	var authMiddleware auth.AuthMiddlewareDelivery
	var usersRepository account.UsersRepository

	usersRepository = mysql.NewUsersRepository(server.DB)

	accountUsecase = usecase.NewAccountUsecase(&server.Config, usersRepository)
	authMiddleware = AuthMiddlewareDI(server)

	httprest.NewAccountHTTPRestHandler(route, accountUsecase, authMiddleware)

}