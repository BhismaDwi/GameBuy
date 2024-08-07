package game

import (
	"GameBuy/helpers/common"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllGameService(ctx *gin.Context) (result []Game, err error)
	GetGameByIdService(ctx *gin.Context) (result Game, err error)
	// GetAllGamesByGameService(ctx *gin.Context) (result []book.Book, err error)
	CreateGameService(ctx *gin.Context) (err error)
	DeleteGameService(ctx *gin.Context) (err error)
	UpdateGameService(ctx *gin.Context) (err error)
}

type gameService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &gameService{repository}
}

func (service *gameService) GetAllGameService(ctx *gin.Context) (game []Game, err error) {
	return service.repository.GetAll()
}

func (service *gameService) GetGameByIdService(ctx *gin.Context) (game Game, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)
	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id game from param")
		return
	}
	return service.repository.GetByID(idInt)
}

func (service *gameService) CreateGameService(ctx *gin.Context) (err error) {
	var newGame Game

	err = ctx.ShouldBind(&newGame)
	if err != nil {
		return err
	}

	var games []Game
	games, err = service.repository.GetAll()
	if err != nil {
		return err
	}

	game, err := service.repository.GetByTitle(newGame.Title)
	if err != nil {
		return err
	}

	if len(games) != 0 && game.Title != "" {
		err = errors.New("game already exists")
		return err
	}

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newGame.CreatedAt = defaultField.CreatedAt
	newGame.CreatedBy = defaultField.CreatedBy
	newGame.ModifiedAt = defaultField.ModifiedAt
	newGame.ModifiedBy = defaultField.ModifiedBy

	err = service.repository.Create(newGame)
	if err != nil {
		return errors.New("failed to add new game")
	}

	return
}

func (service *gameService) DeleteGameService(ctx *gin.Context) (err error) {
	var (
		game Game
		id   = ctx.Param("id")
	)

	game.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id game from param")
		return
	}

	return service.repository.Delete(game)
}

func (service *gameService) UpdateGameService(ctx *gin.Context) (err error) {
	var (
		game Game
		id   = ctx.Param("id")
	)

	err = ctx.ShouldBind(&game)
	if err != nil {
		return
	}

	game.ID, err = strconv.Atoi(id)

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	game.ModifiedAt = defaultField.ModifiedAt
	game.ModifiedBy = defaultField.ModifiedBy

	if err != nil {
		err = errors.New("failed to get id game from param")
		return
	}
	return service.repository.Update(game)
}
