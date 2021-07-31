package httpobject

import (
	"io/ioutil"
	"log"
	"mime/multipart"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/gin-gonic/gin"
)

type updateProfileAvatarFormBody struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

type updateProfileAvatarResponseBody struct {
	AvatarURL *domain.Image `json:"avatar_url"`
}

func (handler *accountHTTPObjectHandler) UpdateProfileAvatar(ctx *gin.Context) {

	var err error
	var statusCode domain.HTTPStatusCode
	var authUserID *domain.UUID

	var authHeaderInterface interface{}
	var isAuthHeaderExists bool = false

	authHeaderInterface, isAuthHeaderExists = ctx.Get("AUTH_HEADER")

	if !isAuthHeaderExists {
		log.Println("Auth header not found")
		ctx.String(int(statusCode), "%s", "Unauthorized")
		return
	}

	var isConversionOK bool = false

	authUserID, isConversionOK = authHeaderInterface.(*domain.UUID)

	if !isConversionOK {
		log.Println("Cannot convert interface{} to *domain.UUID")
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	requestForm := &updateProfileAvatarFormBody{}

	err = ctx.Bind(requestForm)

	if err != nil {
		ctx.String(400, "%s", err.Error())
		return
	}

	var avatarFileReader multipart.File

	avatarFileReader, err = requestForm.Avatar.Open()

	if err != nil {
		ctx.String(400, "%s", err.Error())
		return
	}

	var bytesAvatarFile []byte

	bytesAvatarFile, err = ioutil.ReadAll(avatarFileReader)

	if err != nil {
		ctx.String(400, "%s", err.Error())
		return
	}

	var avatarURL *domain.Image

	avatarURL, err, statusCode = handler.accountUsecase.UpdateProfileAvatar(
		authUserID,
		bytesAvatarFile,
		requestForm.Avatar.Header.Get("Content-Type"),
	)

	if err != nil {
		ctx.String(int(statusCode), "%s", err.Error())
		return
	}

	ctx.JSON(int(statusCode), updateProfileAvatarResponseBody{AvatarURL: avatarURL})

}
