package game

import (
	"GameBuy/databases/connection"
	"GameBuy/helpers/common"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	// api.Use(middlewares.JwtMiddleware())
	// api.Use(middlewares.Logging())
	{
		api.GET("/games", GetAllGameRouter)
		api.GET("/games/:id", GetGameByIdRouter)
		// api.GET("/games/:id/games", GetAllGamesByGameRouter)
		api.POST("/games", CreateGameRouter)
		api.PUT("/games/:id", UpdateGameRouter)
		api.DELETE("/games/:id", DeleteGameRouter)
	}
}

func GetAllGameRouter(ctx *gin.Context) {
	var (
		gameRepo = NewRepository(connection.DBConnections)
		gameSrv  = NewService(gameRepo)
	)

	games, err := gameSrv.GetAllGameService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all game data", int64(len(games)), games)
}

func GetGameByIdRouter(ctx *gin.Context) {
	var (
		gameRepo = NewRepository(connection.DBConnections)
		gameSrv  = NewService(gameRepo)
	)

	game, err := gameSrv.GetGameByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully get game data", game)
}

// func GetAllGamesByGameRouter(ctx *gin.Context) {
// 	var (
// 		gameRepo = NewRepository(connection.DBConnections)
// 		gameSrv  = NewService(gameRepo)
// 	)

// books, err := gameSrv.GetAllBooksByGameService(ctx)
// if err != nil {
// 	common.GenerateErrorResponse(ctx, err.Error())
// 	return
// }

// common.GenerateSuccessResponseWithListData(ctx, "successfully get all book by game", int64(len(books)), books)
// }

func CreateGameRouter(ctx *gin.Context) {
	var (
		gameRepo = NewRepository(connection.DBConnections)
		gameSrv  = NewService(gameRepo)
	)

	err := gameSrv.CreateGameService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added game data")
}

func DeleteGameRouter(ctx *gin.Context) {
	var (
		gameRepo = NewRepository(connection.DBConnections)
		gameSrv  = NewService(gameRepo)
	)

	err := gameSrv.DeleteGameService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete game data")
}

func UpdateGameRouter(ctx *gin.Context) {
	var (
		gameRepo = NewRepository(connection.DBConnections)
		gameSrv  = NewService(gameRepo)
	)

	err := gameSrv.UpdateGameService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update game data")
}
