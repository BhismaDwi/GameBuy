package platform

import (
	"GameBuy/databases/connection"
	"GameBuy/helpers/common"
	"GameBuy/middlewares"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.GET("/platforms", GetAllPlatformRouter)
		api.GET("/platforms/:id", GetPlatformByIdRouter)
		// api.GET("/platforms/:id/games", GetAllGamesByCategoryRouter)
		api.POST("/platforms", CreatePlatformRouter)
		api.PUT("/platforms/:id", UpdatePlatformRouter)
		api.DELETE("/platforms/:id", DeletePlatformRouter)
	}
}

func GetAllPlatformRouter(ctx *gin.Context) {
	var (
		platformRepo = NewRepository(connection.DBConnections)
		platformSrv  = NewService(platformRepo)
	)

	platforms, err := platformSrv.GetAllPlatformService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all platform data", int64(len(platforms)), platforms)
}

func GetPlatformByIdRouter(ctx *gin.Context) {
	var (
		platformRepo = NewRepository(connection.DBConnections)
		platformSrv  = NewService(platformRepo)
	)

	platform, err := platformSrv.GetPlatformByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully get platform data", platform)
}

// func GetAllGamesByCategoryRouter(ctx *gin.Context) {
// 	var (
// 		platformRepo = NewRepository(connection.DBConnections)
// 		platformSrv  = NewService(platformRepo)
// 	)

// books, err := platformSrv.GetAllBooksByCategoryService(ctx)
// if err != nil {
// 	common.GenerateErrorResponse(ctx, err.Error())
// 	return
// }

// common.GenerateSuccessResponseWithListData(ctx, "successfully get all book by category", int64(len(books)), books)
// }

func CreatePlatformRouter(ctx *gin.Context) {
	var (
		platformRepo = NewRepository(connection.DBConnections)
		platformSrv  = NewService(platformRepo)
	)

	err := platformSrv.CreatePlatformService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added platform data")
}

func DeletePlatformRouter(ctx *gin.Context) {
	var (
		platformRepo = NewRepository(connection.DBConnections)
		platformSrv  = NewService(platformRepo)
	)

	err := platformSrv.DeletePlatformService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete platform data")
}

func UpdatePlatformRouter(ctx *gin.Context) {
	var (
		platformRepo = NewRepository(connection.DBConnections)
		platformSrv  = NewService(platformRepo)
	)

	err := platformSrv.UpdatePlatformService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully update platform data")
}
