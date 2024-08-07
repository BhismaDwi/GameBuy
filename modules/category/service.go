package category

import (
	"GameBuy/helpers/common"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllCategoryService(ctx *gin.Context) (result []Category, err error)
	GetCategoryByIdService(ctx *gin.Context) (result Category, err error)
	// GetAllGamesByCategoryService(ctx *gin.Context) (result []book.Book, err error)
	CreateCategoryService(ctx *gin.Context) (err error)
	DeleteCategoryService(ctx *gin.Context) (err error)
	UpdateCategoryService(ctx *gin.Context) (err error)
}

type categoryService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &categoryService{repository}
}

func (service *categoryService) GetAllCategoryService(ctx *gin.Context) (category []Category, err error) {
	return service.repository.GetAll()
}

func (service *categoryService) GetCategoryByIdService(ctx *gin.Context) (category Category, err error) {
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

func (service *categoryService) CreateCategoryService(ctx *gin.Context) (err error) {
	var newCategory Category

	err = ctx.ShouldBind(&newCategory)
	if err != nil {
		return err
	}

	var categorys []Category
	categorys, err = service.repository.GetAll()
	if err != nil {
		return err
	}

	category, err := service.repository.GetByName(newCategory.Name)
	if err != nil {
		return err
	}

	if len(categorys) != 0 && category.Name != "" {
		err = errors.New("category already exists")
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newCategory.CreatedAt = defaultField.CreatedAt
	newCategory.CreatedBy = defaultField.CreatedBy
	newCategory.ModifiedAt = defaultField.ModifiedAt
	newCategory.ModifiedBy = defaultField.ModifiedBy

	err = service.repository.Create(newCategory)
	if err != nil {
		return errors.New("failed to add new category")
	}

	return
}

func (service *categoryService) DeleteCategoryService(ctx *gin.Context) (err error) {
	var (
		category Category
		id       = ctx.Param("id")
	)

	category.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id category from param")
		return
	}

	return service.repository.Delete(category)
}

func (service *categoryService) UpdateCategoryService(ctx *gin.Context) (err error) {
	var (
		category Category
		id       = ctx.Param("id")
	)

	err = ctx.ShouldBind(&category)
	if err != nil {
		return
	}

	category.ID, err = strconv.Atoi(id)

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	category.ModifiedAt = defaultField.ModifiedAt
	category.ModifiedBy = defaultField.ModifiedBy

	if err != nil {
		err = errors.New("failed to get id category from param")
		return
	}
	return service.repository.Update(category)
}
