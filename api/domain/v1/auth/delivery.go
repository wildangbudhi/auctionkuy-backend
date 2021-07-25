package auth

import "github.com/gin-gonic/gin"

type AuthMiddlewareDelivery interface {
	ValidateAccessToken(ctx *gin.Context)
}
