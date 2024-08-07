package platform

import (
	"GameBuy/helpers/common"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllPlatformService(ctx *gin.Context) (result []Platform, err error)
	GetPlatformByIdService(ctx *gin.Context) (result Platform, err error)
	// GetAllGamesByPlatformService(ctx *gin.Context) (result []book.Book, err error)
	CreatePlatformService(ctx *gin.Context) (err error)
	DeletePlatformService(ctx *gin.Context) (err error)
	UpdatePlatformService(ctx *gin.Context) (err error)
}

type platformService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &platformService{repository}
}

func (service *platformService) GetAllPlatformService(ctx *gin.Context) (platform []Platform, err error) {
	return service.repository.GetAll()
}

func (service *platformService) GetPlatformByIdService(ctx *gin.Context) (platform Platform, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)
	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id category from param")
		return
	}
	return service.repository.GetByID(idInt)
}

func (service *platformService) CreatePlatformService(ctx *gin.Context) (err error) {
	var newPlatform Platform

	err = ctx.ShouldBind(&newPlatform)
	if err != nil {
		return err
	}

	var platforms []Platform
	platforms, err = service.repository.GetAll()
	if err != nil {
		return err
	}

	platform, err := service.repository.GetByName(newPlatform.Name)
	if err != nil {
		return err
	}

	if len(platforms) != 0 && platform.Name != "" {
		err = errors.New("platform already exists")
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newPlatform.CreatedAt = defaultField.CreatedAt
	newPlatform.CreatedBy = defaultField.CreatedBy
	newPlatform.ModifiedAt = defaultField.ModifiedAt
	newPlatform.ModifiedBy = defaultField.ModifiedBy

	err = service.repository.Create(newPlatform)
	if err != nil {
		return errors.New("failed to add new platform")
	}

	return
}

func (service *platformService) DeletePlatformService(ctx *gin.Context) (err error) {
	var (
		platform Platform
		id       = ctx.Param("id")
	)

	platform.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id platform from param")
		return
	}

	return service.repository.Delete(platform)
}

func (service *platformService) UpdatePlatformService(ctx *gin.Context) (err error) {
	var (
		platform Platform
		id       = ctx.Param("id")
	)

	err = ctx.ShouldBind(&platform)
	if err != nil {
		return
	}

	platform.ID, err = strconv.Atoi(id)

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	platform.ModifiedAt = defaultField.ModifiedAt
	platform.ModifiedBy = defaultField.ModifiedBy

	if err != nil {
		err = errors.New("failed to get id platform from param")
		return
	}
	return service.repository.Update(platform)
}
