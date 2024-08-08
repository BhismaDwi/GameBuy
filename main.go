package main

import (
	"GameBuy/configs"
	"GameBuy/databases/connection"
	"GameBuy/modules/category"
	"GameBuy/modules/game"
	"GameBuy/modules/platform"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	// migration.Initiator(connection.DBConnections)

	InitiateRouter()
}

func InitiateRouter() {
	router := gin.Default()

	platform.Initiator(router)
	category.Initiator(router)
	game.Initiator(router)

	router.Run(":8081")
}
