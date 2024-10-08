package users

import (
	"GameBuy/databases/connection"
	"GameBuy/helpers/common"

	"github.com/gin-gonic/gin"
)

func Initiator(router *gin.Engine) {
	api := router.Group("/api/users")
	{
		api.POST("/login", Login)
		api.POST("/signup", SignUp)
	}
}

func Login(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userSrv  = NewService(userRepo)
	)

	token, err := userSrv.LoginService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully login", token)
}

func SignUp(ctx *gin.Context) {
	var (
		userRepo = NewRepository(connection.DBConnections)
		userSrv  = NewService(userRepo)
	)

	err := userSrv.SignUpService(ctx)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "awesome, successfully create user")
}
