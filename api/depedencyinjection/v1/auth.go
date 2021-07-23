package v1

import (
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"auctionkuy.wildangbudhi.com/services/v1/auth/delivery/httprest"
	httprestrepo "auctionkuy.wildangbudhi.com/services/v1/auth/repository/httprest"
	"auctionkuy.wildangbudhi.com/services/v1/auth/repository/mysql"
	"auctionkuy.wildangbudhi.com/services/v1/auth/repository/redis"
	"auctionkuy.wildangbudhi.com/services/v1/auth/usecase"
	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func AuthDI(server *utils.Server) {

	var route *gin.RouterGroup = server.Router.Group("/v1/auth")

	var authUsecase auth.AuthUsecase
	var appleKeysRepository auth.AppleKeysRepository
	var usersRepository auth.UsersRepository
	var sessionRepository auth.SessionRepository

	appleKeysRepository = httprestrepo.NewAppleKeysRepository()
	usersRepository = mysql.NewUsersRepository(server.DB)
	sessionRepository = redis.NewSessionRepository(server.RedisDB)

	authUsecase = usecase.NewAuthUsecase(
		&server.Config,
		appleKeysRepository,
		usersRepository,
		sessionRepository,
	)

	httprest.NewAuthHTTPRestHandler(route, authUsecase)

}
