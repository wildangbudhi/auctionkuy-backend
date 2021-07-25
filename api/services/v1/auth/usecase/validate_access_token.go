package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/dgrijalva/jwt-go"
)

func (usecase *authUsecase) ValidateAccessToken(token string) (string, error, domain.HTTPStatusCode) {

	var err error

	var authToken auth.AuthToken = NewAuthToken(usecase.serverConfig.SecretKey, usecase.sessionRepository)

	var tokenJWT *domain.JWT

	tokenJWT, err = domain.NewJWT(token, usecase.serverConfig.SecretKey, jwt.SigningMethodHS512)

	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("Unauthorized"), 401
	}

	var userID string
	var keyExist bool

	userID, keyExist = tokenJWT.GetPayload()["user_id"].(string)

	if !keyExist {
		log.Println("User ID Payload Not Found")
		return "", fmt.Errorf("Unauthorized"), 401
	}

	err = authToken.ValidateToken(tokenJWT, false)

	if err != nil {
		return "", fmt.Errorf("Unauthorized"), 401
	}

	return userID, nil, 200

}
