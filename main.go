package main

import (
	"GameBuy/configs"
	"GameBuy/databases/connection"
)

func main() {
	configs.Initiator()

	connection.Initiator()
	defer connection.DBConnections.Close()
}
