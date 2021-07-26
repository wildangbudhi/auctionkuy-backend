package httpobject

import (
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

func (handler *assetsHTTPObjectHandler) BankObject(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var object []byte
	var objectContentType string

	var objectName string = ctx.Param("object-name")

	if objectName == "" {
		ctx.String(http.StatusBadRequest, "Object name invalid")
		return
	}

	object, objectContentType, err, statusCode = handler.assetsUsecase.BankObject(objectName)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	ctx.Data(int(statusCode), objectContentType, object)

}
