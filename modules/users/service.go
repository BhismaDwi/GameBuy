package users

import (
	"GameBuy/helpers/common"
	"GameBuy/middlewares"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	LoginService(ctx *gin.Context) (result LoginResponse, err error)
	SignUpService(ctx *gin.Context) (err error)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (service *userService) LoginService(ctx *gin.Context) (result LoginResponse, err error) {
	var userReq LoginRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return
	}

	err = userReq.ValidateLogin()
	if err != nil {
		return
	}

	user, err := service.repository.Login(userReq)
	if err != nil {
		return
	}

	if common.IsEmptyField(user.ID) {
		err = errors.New("invalid account")
		return
	}

	matches := common.CheckPassword(user.Password, userReq.Password)
	if !matches {
		err = errors.New("wrong username or password")
		return
	}

	jwtToken, err := middlewares.GenerateJwtToken()
	if err != nil {
		return
	}

	middlewares.DummyRedis[jwtToken] = middlewares.UserLoginRedis{
		UserId:    int64(user.ID),
		Username:  user.Username,
		Role:      user.Role.Name,
		LoginAt:   time.Now(),
		ExpiredAt: time.Now().Add(time.Hour * 1),
	}

	result.Token = jwtToken

	return
}

func (service *userService) SignUpService(ctx *gin.Context) (err error) {
	var userReq SignUpRequest

	err = ctx.ShouldBind(&userReq)
	if err != nil {
		return err
	}

	err = userReq.ValidateSignUp()
	if err != nil {
		return err
	}

	var users []User
	users, errUsers := service.repository.GetList()
	if errUsers != nil {
		return errUsers
	}

	user, err := service.repository.GetUserByUsername(userReq.Username)
	if err != nil {
		return err
	}

	if len(users) != 0 && user.Username != "" {
		err = errors.New("username already exist")
		return err
	}

	user, err = userReq.ConvertToModelForSignUp()
	if err != nil {
		return err
	}

	err = service.repository.SignUp(user)
	if err != nil {
		return err
	}

	return nil
}
