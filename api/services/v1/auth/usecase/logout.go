package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/dgrijalva/jwt-go"
)

func (usecase *authUsecase) Logout(token string) (error, domain.HTTPStatusCode) {

	var err error
	var authToken auth.AuthToken = NewAuthToken(usecase.serverConfig.SecretKey, usecase.sessionRepository)

	var tokenJWT *domain.JWT

	tokenJWT, err = domain.NewJWT(token, usecase.serverConfig.SecretKey, jwt.SigningMethodHS512)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Unauthorized"), 401
	}

	err = authToken.RemoveAuthToken(tokenJWT)

	if err != nil {
		return err, 400
	}

	return nil, 200

}
