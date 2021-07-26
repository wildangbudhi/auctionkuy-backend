package main

import (
	"log"

	v1 "auctionkuy.wildangbudhi.com/depedencyinjection/v1"
	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	HealthCheckHandler(server)

	depedencyInjection(server)
	server.Router.Run(server.Config.ServerHost)
}

func depedencyInjection(server *utils.Server) {
	v1.AuthHTTPRestDI(server)
	v1.AccountHTTPDI(server)
	v1.AssetsHTTPDI(server)
}

func HealthCheckHandler(server *utils.Server) {
	server.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
