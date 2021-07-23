package domain

import (
	"crypto/rsa"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	payload   jwt.MapClaims
	token     string
	alg       jwt.SigningMethod
	secretKey []byte
}

func NewJWT(token string, secretKey []byte, alg jwt.SigningMethod) (*JWT, error) {

	var err error

	var authToken JWT = JWT{
		token:     token,
		secretKey: secretKey,
		alg:       alg,
	}

	err = authToken.decodeToken()

	if err != nil {
		return nil, err
	}

	return &authToken, nil

}

func NewJWTFromPayload(payload map[string]interface{}, secretKey []byte, alg jwt.SigningMethod) (*JWT, error) {

	var err error

	var authToken JWT = JWT{
		payload:   payload,
		secretKey: secretKey,
		alg:       alg,
	}

	err = authToken.encodeToken()

	if err != nil {
		return nil, err
	}

	return &authToken, nil

}

func (obj *JWT) GetToken() string {
	return obj.token
}

func (obj *JWT) GetPayload() map[string]interface{} {
	return obj.payload
}

func (obj *JWT) decodeToken() error {
	token, err := jwt.Parse(obj.token, func(token *jwt.Token) (interface{}, error) {

		if strings.HasPrefix(obj.alg.Alg(), "RS") {

			var err error
			var key *rsa.PublicKey

			key, err = jwt.ParseRSAPublicKeyFromPEM(obj.secretKey)

			if err != nil {
				return "", fmt.Errorf("validate: parse key: %w", err)
			}

			return key, nil
		}

		return obj.secretKey, nil

	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return fmt.Errorf("Token Auth Invalid")
	}

	obj.payload = claims

	return nil
}

func (obj *JWT) encodeToken() error {

	aksesToken := jwt.NewWithClaims(obj.alg, obj.payload)

	token, err := aksesToken.SignedString(obj.secretKey)

	if err != nil {
		return err
	}

	obj.token = token

	return nil
}
