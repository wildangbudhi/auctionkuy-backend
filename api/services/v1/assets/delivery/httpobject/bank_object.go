package httpobject

import (
	"net/http"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (handler *assetsHTTPObjectHandler) BankObject(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var object *minio.Object
	var objectStats *minio.ObjectInfo

	var objectName string = ctx.Param("object-name")

	if objectName == "" {
		ctx.String(http.StatusBadRequest, "Object name invalid")
		return
	}

	object, objectStats, err, statusCode = handler.assetsUsecase.BankObject(objectName)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	ctx.DataFromReader(int(statusCode), objectStats.Size, objectStats.ContentType, object, nil)

}
