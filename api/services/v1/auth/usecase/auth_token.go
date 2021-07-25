package usecase

import (
	"fmt"
	"log"
	"time"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/dgrijalva/jwt-go"
)

type authToken struct {
	secretKey         []byte
	sessionRepository auth.SessionRepository
}

func NewAuthToken(secretKey []byte, sessionRepository auth.SessionRepository) auth.AuthToken {
	return &authToken{
		secretKey:         secretKey,
		sessionRepository: sessionRepository,
	}
}

func (obj *authToken) GenerateAuthToken(userID *domain.UUID) (*domain.JWT, *domain.JWT, *domain.Timestamp, error) {

	var err error

	var expTimestamp *domain.Timestamp

	expTimestamp, err = domain.NewNowTimestamp()

	if err != nil {
		return nil, nil, nil, err
	}

	expTimestamp.SetValue(expTimestamp.GetValue().Add(time.Hour * 24))

	var token, refreshToken *domain.JWT
	var tokenPayload, refreshTokenPayload jwt.MapClaims
	var tokenUUID *domain.UUID = domain.NewUUID()
	var refreshTokenUUID *domain.UUID = domain.NewUUID()

	tokenPayload = jwt.MapClaims{}
	tokenPayload["user_id"] = userID
	tokenPayload["uuid"] = tokenUUID.GetValue()
	tokenPayload["exp"] = expTimestamp.GetValue().Unix()

	token, err = domain.NewJWTFromPayload(tokenPayload, obj.secretKey, jwt.SigningMethodHS512)

	if err != nil {
		log.Println(err)
		return nil, nil, nil, fmt.Errorf("Failed to Generate Session")
	}

	refreshTokenPayload = jwt.MapClaims{}
	refreshTokenPayload["user_id"] = userID
	refreshTokenPayload["uuid"] = refreshTokenUUID.GetValue()

	refreshToken, err = domain.NewJWTFromPayload(refreshTokenPayload, obj.secretKey, jwt.SigningMethodHS512)

	if err != nil {
		log.Println(err)
		return nil, nil, nil, fmt.Errorf("Failed to Generate Session")
	}

	var sessionKey string = fmt.Sprintf("auth-token-%s", userID.GetValue())
	var sessionData auth.Session = auth.Session{
		AccessUUID:  tokenUUID.GetValue(),
		RefreshUUID: refreshTokenUUID.GetValue(),
	}

	err = obj.sessionRepository.SetSession(sessionKey, &sessionData, time.Hour*8760)

	if err != nil {
		return nil, nil, nil, err
	}

	return token, refreshToken, expTimestamp, nil

}

func (obj *authToken) ValidateToken(token *domain.JWT, isRefreshToken bool) error {

	var err error
	var keyExist bool

	var tokenPayload map[string]interface{} = token.GetPayload()
	var userID string
	var tokenUUID string

	userID, keyExist = tokenPayload["user_id"].(string)

	if !keyExist {
		return fmt.Errorf("Unauthorized")
	}

	tokenUUID, keyExist = tokenPayload["uuid"].(string)

	if !keyExist {
		return fmt.Errorf("Unauthorized")
	}

	var sessionKey string = fmt.Sprintf("auth-token-%s", userID)

	var sessionData *auth.Session

	sessionData, err = obj.sessionRepository.GetSession(sessionKey)

	if err != nil {
		return fmt.Errorf("Unauthorized")
	}

	var cacheTokenUUID string

	if isRefreshToken {
		cacheTokenUUID = sessionData.RefreshUUID
	} else {
		cacheTokenUUID = sessionData.AccessUUID
	}

	if tokenUUID != cacheTokenUUID {
		return fmt.Errorf("Unauthorized")
	}

	err = obj.sessionRepository.ExtendSessionExpiration(sessionKey, time.Hour*8760)

	if err != nil {
		return fmt.Errorf("Unauthorized")
	}

	return nil

}

func (obj *authToken) RegenerateAuthToken(refreshToken *domain.JWT) (*domain.JWT, *domain.Timestamp, error) {

	var err error

	err = obj.ValidateToken(refreshToken, true)

	if err != nil {
		return nil, nil, fmt.Errorf("Refresh Token Invalid")
	}

	var tokenPayload map[string]interface{} = refreshToken.GetPayload()
	var keyExist bool
	var userID string
	var tokenUUID string

	userID, keyExist = tokenPayload["user_id"].(string)

	if !keyExist {
		return nil, nil, fmt.Errorf("Refresh Token Invalid")
	}

	tokenUUID, keyExist = tokenPayload["uuid"].(string)

	if !keyExist {
		return nil, nil, fmt.Errorf("Refresh Token Invalid")
	}

	var expTimestamp *domain.Timestamp

	expTimestamp, err = domain.NewNowTimestamp()

	if err != nil {
		return nil, nil, err
	}

	expTimestamp.SetValue(expTimestamp.GetValue().Add(time.Hour * 24))

	var newAccessTokenUUID *domain.UUID = domain.NewUUID()

	var newAccessTokenPayload jwt.MapClaims = jwt.MapClaims{}
	newAccessTokenPayload["user_id"] = userID
	newAccessTokenPayload["uuid"] = newAccessTokenUUID.GetValue()
	newAccessTokenPayload["exp"] = expTimestamp.GetValue().Unix()

	var newAccessToken *domain.JWT
	newAccessToken, err = domain.NewJWTFromPayload(newAccessTokenPayload, obj.secretKey, jwt.SigningMethodHS512)

	if err != nil {
		log.Println(err)
		return nil, nil, fmt.Errorf("Failed to generate new session")
	}

	var sessionKey string = fmt.Sprintf("auth-token-%s", userID)
	var sessionData auth.Session = auth.Session{
		AccessUUID:  newAccessTokenUUID.GetValue(),
		RefreshUUID: tokenUUID,
	}

	err = obj.sessionRepository.SetSession(sessionKey, &sessionData, 0)

	if err != nil {
		return nil, nil, err
	}

	return newAccessToken, expTimestamp, nil

}

func (obj *authToken) RemoveAuthToken(token *domain.JWT) error {

	var err error
	var keyExist bool

	var tokenPayload map[string]interface{} = token.GetPayload()
	var userID string

	userID, keyExist = tokenPayload["user_id"].(string)

	if !keyExist {
		return fmt.Errorf("Access Token Invalid")
	}

	var sessionKey string = fmt.Sprintf("auth-token-%s", userID)

	err = obj.sessionRepository.RemoveSession(sessionKey)

	if err != nil {
		return err
	}

	return nil

}
