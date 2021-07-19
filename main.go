package main

import (
	"log"

	"auctionkuy.wildangbudhi.com/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	server, err := utils.NewServer()

	if err != nil {
		log.Fatal(err)
	}

	server.Router.LoadHTMLGlob("templates/*")
	HealthCheckHandler(server)

	depedencyInjection(server)
	server.Router.Run(server.Config.ServerHost)
}

func depedencyInjection(server *utils.Server) {
}

func HealthCheckHandler(server *utils.Server) {
	server.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
