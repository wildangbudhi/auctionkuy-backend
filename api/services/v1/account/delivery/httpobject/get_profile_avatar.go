package httpobject

import (
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

func (handler *accountHTTPObjectHandler) GetProfileAvatar(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var object []byte
	var objectContentType string

	var userID string = ctx.Param("userid")

	var userUUID *domain.UUID

	userUUID, err = domain.NewUUIDFromString(userID)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	object, objectContentType, err, statusCode = handler.accountUsecase.GetProfileAvatar(userUUID)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	log.Println(objectContentType)

	ctx.Data(int(statusCode), objectContentType, object)

}
