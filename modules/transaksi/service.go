package transaksi

import (
	"GameBuy/helpers/common"
	"GameBuy/modules/game"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllTransaksiService(ctx *gin.Context) (result []Transaksi, err error)
	GetTransaksiByIdService(ctx *gin.Context) (result Transaksi, err error)
	// GetAllTransaksisByTransaksiService(ctx *gin.Context) (result []book.Book, err error)
	CreateTransaksiService(ctx *gin.Context) (err error)
	DeleteTransaksiService(ctx *gin.Context) (err error)
}

type transaksiService struct {
	repository     Repository
	gameRepository game.Repository
}

func NewService(repository Repository) Service {
	return &transaksiService{repository: repository}
}
func NewServiceWithGame(repository Repository, gameRepository game.Repository) Service {
	return &transaksiService{repository: repository, gameRepository: gameRepository}
}

func (service *transaksiService) GetAllTransaksiService(ctx *gin.Context) (transaksi []Transaksi, err error) {
	return service.repository.GetAll()
}

func (service *transaksiService) GetTransaksiByIdService(ctx *gin.Context) (transaksi Transaksi, err error) {
	var (
		idInt int
		id    = ctx.Param("id")
	)
	idInt, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id transaksi from param")
		return
	}
	return service.repository.GetByID(idInt)
}

func (service *transaksiService) CreateTransaksiService(ctx *gin.Context) (err error) {
	var newTransaksi Transaksi

	err = ctx.ShouldBind(&newTransaksi)
	if err != nil {
		return err
	}
	// Check if platform exists
	for _, detail := range newTransaksi.Details {
		platformExists, err := service.gameRepository.CheckGameExists(detail.GameID)
		if err != nil {
			return err
		}
		if !platformExists {
			return errors.New("game does not exist")
		}
	}
	// userID, exists := ctx.Get("userID")
	// if !exists {
	// 	return errors.New("user ID not found in context")
	// }

	defaultField := common.DefaultFieldTable{}
	defaultField.SetDefaultField()

	newTransaksi.CreatedAt = defaultField.CreatedAt
	newTransaksi.CreatedBy = defaultField.CreatedBy
	newTransaksi.ModifiedAt = defaultField.ModifiedAt
	newTransaksi.ModifiedBy = defaultField.ModifiedBy
	// newTransaksi.UserID = userID.(int)

	err = service.repository.Create(newTransaksi)
	if err != nil {
		return errors.New("failed to add new transaksi")
	}

	return
}

func (service *transaksiService) DeleteTransaksiService(ctx *gin.Context) (err error) {
	var (
		transaksi Transaksi
		id        = ctx.Param("id")
	)

	transaksi.ID, err = strconv.Atoi(id)
	if err != nil {
		err = errors.New("failed to get id transaksi from param")
		return
	}

	return service.repository.Delete(transaksi)
}
