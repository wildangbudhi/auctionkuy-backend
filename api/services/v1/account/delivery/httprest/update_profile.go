package httprest

import (
	"log"
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
	"github.com/gin-gonic/gin"
)

type updateProfileRequesteBody struct {
	Name                 *string `json:"name"`
	Phone                *string `json:"phone"`
	NationalIDNumber     *string `json:"national_id_number"`
	Locale               *string `json:"locale"`
	BankID               *string `json:"bank_id"`
	BankAccountID        *string `json:"bank_account_id"`
	BankAccountOwnerName *string `json:"bank_account_owner_name"`
}

type updateProfileResponseBody struct {
	*account.Users
}

func (handler *accountHTTPRestHandler) UpdateProfile(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var authUserID *domain.UUID
	var user *account.Users

	ctx.Header("Content-Type", "application/json")

	var authHeaderInterface interface{}
	var isAuthHeaderExists bool = false

	authHeaderInterface, isAuthHeaderExists = ctx.Get("AUTH_HEADER")

	if !isAuthHeaderExists {
		log.Println("Auth header not found")
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Unauthorized"})
		return
	}

	var isConversionOK bool = false

	authUserID, isConversionOK = authHeaderInterface.(*domain.UUID)

	if !isConversionOK {
		log.Println("Cannot convert interface{} to *domain.UUID")
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Unauthorized"})
		return
	}

	requestBodyData := &updateProfileRequesteBody{}

	err = ctx.BindJSON(requestBodyData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	user, err, statusCode = handler.accountUsecase.UpdateProfile(
		authUserID,
		requestBodyData.Name,
		requestBodyData.Phone,
		requestBodyData.NationalIDNumber,
		requestBodyData.Locale,
		requestBodyData.BankID,
		requestBodyData.BankAccountID,
		requestBodyData.BankAccountOwnerName,
	)

	if err != nil {
		ctx.JSON(int(statusCode), domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(int(statusCode), updateProfileResponseBody{Users: user})

}
