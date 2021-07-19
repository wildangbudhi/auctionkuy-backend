package domain

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	payload   jwt.MapClaims
	token     string
	secretKey []byte
}

func NewJWT(token string, secretKey []byte) (*JWT, error) {

	var err error

	var authToken JWT = JWT{
		token:     token,
		secretKey: secretKey,
	}

	err = authToken.decodeToken()

	if err != nil {
		return nil, err
	}

	return &authToken, nil

}

func NewJWTFromPayload(payload map[string]interface{}, secretKey []byte) (*JWT, error) {

	var err error

	var authToken JWT = JWT{
		payload:   payload,
		secretKey: secretKey,
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

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token Auth Invalid")
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

	aksesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, obj.payload)

	token, err := aksesToken.SignedString(obj.secretKey)

	if err != nil {
		return err
	}

	obj.token = token

	return nil
}
