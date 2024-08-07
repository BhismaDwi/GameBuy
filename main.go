package main

import (
	"GameBuy/configs"
	"GameBuy/databases/connection"
	"GameBuy/databases/migration"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()

	migration.Initiator(connection.DBConnections)
}

func InitiateRouter() {
	router := gin.Default()

	// platform.Initiator(router)

	router.Run(":8080")
}
