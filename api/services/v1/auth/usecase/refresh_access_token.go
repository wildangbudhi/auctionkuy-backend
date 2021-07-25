package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/dgrijalva/jwt-go"
)

func (usecase *authUsecase) RefreshAccessToken(refreshToken string) (string, int64, error, domain.HTTPStatusCode) {

	var err error
	var authToken auth.AuthToken = NewAuthToken(usecase.serverConfig.SecretKey, usecase.sessionRepository)

	var tokenJWT, newAccessToken *domain.JWT
	var tokenExp *domain.Timestamp

	tokenJWT, err = domain.NewJWT(refreshToken, usecase.serverConfig.SecretKey, jwt.SigningMethodHS512)

	if err != nil {
		log.Println(err)
		return "", -1, fmt.Errorf("Unauthorized"), 401
	}

	newAccessToken, tokenExp, err = authToken.RegenerateAuthToken(tokenJWT)

	if err != nil {
		return "", -1, err, 401
	}

	return newAccessToken.GetToken(), tokenExp.GetValue().Unix(), nil, 200

}
