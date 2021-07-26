package usecase

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"strings"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/dgrijalva/jwt-go"
)

func (usecase *authUsecase) Authenticate(appleToken, locale string) (string, string, int64, error, domain.HTTPStatusCode) {

	var err error
	var repositoryErrorType domain.RepositoryErrorType

	if appleToken == "" {
		return "", "", -1, fmt.Errorf("Apple Token cannot be empty"), 400
	}

	var appleKeys *auth.AppleKeys

	appleKeys, err = usecase.appleKeysRepository.GetAppleKeys()

	if err != nil {
		log.Println(err)
		return "", "", -1, fmt.Errorf("Apple Token Expired"), 401
	}

	var applePEM []byte

	applePEM, err = getApplePEM(appleToken, appleKeys)

	if err != nil {
		return "", "", -1, err, 401
	}

	var appleJWT *domain.JWT

	appleJWT, err = domain.NewJWT(appleToken, applePEM, jwt.SigningMethodRS256)

	if err != nil {
		log.Println(err)
		return "", "", -1, fmt.Errorf("Apple Token Invalid"), 401
	}

	if !validateAppleJWT(appleJWT) {
		return "", "", -1, fmt.Errorf("Apple Token Invalid"), 401
	}

	var userEmail *domain.Email

	userEmail, err = getEmailFromAppleJWT(appleJWT)

	if err != nil {
		return "", "", -1, err, 401
	}

	var isLocaleExists bool

	_, isLocaleExists = usecase.serverConfig.CountryData.PhoneNumberMaps[locale]

	if !isLocaleExists {
		return "", "", -1, fmt.Errorf("Locale invalid"), 400
	}

	var user *auth.Users

	user, err, repositoryErrorType = usecase.usersRepository.GetUserByEmail(userEmail)

	if repositoryErrorType == domain.RepositoryDataNotFound {

		var isFirstLogin bool = true

		user = &auth.Users{
			ID:         domain.NewUUID(),
			Email:      userEmail,
			Locale:     &locale,
			FirstLogin: &isFirstLogin,
		}
	} else if err == nil {
		user.Locale = &locale
	} else {
		return "", "", -1, err, 400
	}

	if repositoryErrorType == domain.RepositoryDataNotFound {
		err, _ = usecase.usersRepository.CreateUser(user)
	} else if err == nil {
		err, _ = usecase.usersRepository.UpdateUser(user)
	}

	if err != nil {
		return "", "", -1, err, 400
	}

	var authToken auth.AuthToken = NewAuthToken(usecase.serverConfig.SecretKey, usecase.sessionRepository)

	var token, refreshToken *domain.JWT
	var tokenExp *domain.Timestamp

	token, refreshToken, tokenExp, err = authToken.GenerateAuthToken(user.ID)

	if err != nil {
		return "", "", -1, err, 500
	}

	return token.GetToken(), refreshToken.GetToken(), tokenExp.GetValue().Unix(), nil, 200

}

func getEmailFromAppleJWT(appleJWT *domain.JWT) (*domain.Email, error) {

	var tokenPayload map[string]interface{} = appleJWT.GetPayload()

	var emailString string
	var isKeyExist bool

	emailString, isKeyExist = tokenPayload["email"].(string)

	if !isKeyExist {
		return nil, fmt.Errorf("Apple Token Invalid")
	}

	var email *domain.Email
	var err error

	email, err = domain.NewEmail(emailString)

	if err != nil {
		return nil, err
	}

	return email, nil

}

func validateAppleJWT(appleJWT *domain.JWT) bool {

	var tokenPayload map[string]interface{} = appleJWT.GetPayload()

	var iss, aud string
	var isKeyExist bool

	iss, isKeyExist = tokenPayload["iss"].(string)

	if !isKeyExist {
		return false
	}

	aud, isKeyExist = tokenPayload["aud"].(string)

	if !isKeyExist {
		return false
	}

	if iss != "https://appleid.apple.com" {
		return false
	}

	if aud != "com.AuctionKuy.Rekber" {
		return false
	}

	return true

}

func getApplePEM(appleToken string, appleKeys *auth.AppleKeys) ([]byte, error) {

	var err error

	var appleTokenHeaderB64Encoded string = strings.Split(appleToken, ".")[0]
	var appleTokenHeaderByte []byte

	appleTokenHeaderByte, err = base64.RawStdEncoding.DecodeString(appleTokenHeaderB64Encoded)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Apple Token Invalid")
	}

	var appleTokenHeader auth.AppleTokenHeader

	err = json.Unmarshal(appleTokenHeaderByte, &appleTokenHeader)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Apple Token Invalid")
	}

	var selectedApplePubKey auth.ApplePubKey

	for i := 0; i < len(appleKeys.Keys); i++ {

		if appleKeys.Keys[i].Alg == appleTokenHeader.ALG && appleKeys.Keys[i].KID == appleTokenHeader.KID {
			selectedApplePubKey = appleKeys.Keys[i]
			break
		}

	}

	var pubKey *rsa.PublicKey = &rsa.PublicKey{
		N: decodeBase64BigInt(selectedApplePubKey.N),
		E: int(decodeBase64BigInt(selectedApplePubKey.E).Int64()),
	}

	var pubKeyPEM []byte

	pubKeyPEM, err = ExportRsaPublicKeyAsPEM(pubKey)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Apple Token Invalid")
	}

	return pubKeyPEM, nil
}

func decodeBase64BigInt(s string) *big.Int {
	buffer, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	if err != nil {
		log.Fatalf("failed to decode base64: %v", err)
	}

	return big.NewInt(0).SetBytes(buffer)
}

func ExportRsaPublicKeyAsPEM(pubkey *rsa.PublicKey) ([]byte, error) {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return nil, err
	}
	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return pubkey_pem, nil
}
