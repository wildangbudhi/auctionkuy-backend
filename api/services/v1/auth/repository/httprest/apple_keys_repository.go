package httprest

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/auth"
)

type appleKeysRepository struct {
	httpClient *http.Client
}

func NewAppleKeysRepository() auth.AppleKeysRepository {
	return &appleKeysRepository{
		httpClient: new(http.Client),
	}
}

func (repo *appleKeysRepository) httpRequestCall(method string, url string, body io.Reader) (*http.Response, error) {

	var err error

	var request *http.Request
	request, err = http.NewRequest(method, url, body)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Services Unavailable")
	}

	var response *http.Response
	response, err = repo.httpClient.Do(request)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Services Unavailable")
	}

	if int(response.StatusCode/100) != 2 {
		var errResponseBody domain.HTTPErrorResponse = domain.HTTPErrorResponse{}

		err = json.NewDecoder(response.Body).Decode(&errResponseBody)

		if err != nil {

			var bodyBytes []byte
			bodyBytes, err = ioutil.ReadAll(response.Body)

			if err != nil {
				return nil, fmt.Errorf("Services Unavailable")
			}

			return nil, fmt.Errorf("%s", string(bodyBytes))

		}

		return nil, fmt.Errorf("%s", errResponseBody.Error)
	}

	return response, nil

}

func (repo *appleKeysRepository) GetAppleKeys() (*auth.AppleKeys, error) {

	var err error

	var url string = "https://appleid.apple.com/auth/keys"

	var response *http.Response

	response, err = repo.httpRequestCall("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var appleKeys auth.AppleKeys = auth.AppleKeys{}

	err = json.NewDecoder(response.Body).Decode(&appleKeys)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Services Unavailable")
	}

	return &appleKeys, nil

}
