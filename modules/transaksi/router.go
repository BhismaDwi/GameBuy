package transaksi

import (
	"GameBuy/databases/connection"
	"GameBuy/helpers/common"
	"GameBuy/middlewares"
	"GameBuy/modules/game"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api")
	api.Use(middlewares.JwtMiddleware())
	api.Use(middlewares.Logging())
	{
		api.GET("/transaksi", GetAllTransaksiRouter)
		api.GET("/transaksi/:id", GetTransaksiByIdRouter)
		// api.GET("/transaksis/:id/transaksis", GetAllTransaksisByTransaksiRouter)
		api.POST("/transaksi", CreateTransaksiRouter)
		api.DELETE("/transaksi/:id", DeleteTransaksiRouter)
	}
}

func GetAllTransaksiRouter(ctx *gin.Context) {
	var (
		transaksiRepo = NewRepository(connection.DBConnections)
		transaksiSrv  = NewService(transaksiRepo)
	)

	transaksis, err := transaksiSrv.GetAllTransaksiService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all transaksi data", int64(len(transaksis)), transaksis)
}

func GetTransaksiByIdRouter(ctx *gin.Context) {
	var (
		transaksiRepo = NewRepository(connection.DBConnections)
		transaksiSrv  = NewService(transaksiRepo)
	)

	transaksi, err := transaksiSrv.GetTransaksiByIdService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}
	common.GenerateSuccessResponseWithData(ctx, "successfully get transaksi data", transaksi)
}

// func GetAllTransaksisByTransaksiRouter(ctx *gin.Context) {
// 	var (
// 		transaksiRepo = NewRepository(connection.DBConnections)
// 		transaksiSrv  = NewService(transaksiRepo)
// 	)

// books, err := transaksiSrv.GetAllBooksByTransaksiService(ctx)
// if err != nil {
// 	common.GenerateErrorResponse(ctx, err.Error())
// 	return
// }

// common.GenerateSuccessResponseWithListData(ctx, "successfully get all book by transaksi", int64(len(books)), books)
// }

func CreateTransaksiRouter(ctx *gin.Context) {
	var (
		transaksiRepo = NewRepository(connection.DBConnections)
		gameRepo      = game.NewRepository(connection.DBConnections)
		transaksiSrv  = NewServiceWithGame(transaksiRepo, gameRepo)
	)

	err := transaksiSrv.CreateTransaksiService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully added transaksi data")
}

func DeleteTransaksiRouter(ctx *gin.Context) {
	var (
		transaksiRepo = NewRepository(connection.DBConnections)
		transaksiSrv  = NewService(transaksiRepo)
	)

	err := transaksiSrv.DeleteTransaksiService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully delete transaksi data")
}
