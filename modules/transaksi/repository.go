package transaksi

import (
	"GameBuy/helpers/constant"
	"GameBuy/modules/category"
	"GameBuy/modules/game"
	"GameBuy/modules/platform"
	"GameBuy/modules/transaksidetail"
	"GameBuy/modules/users"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll() (transaksies []Transaksi, err error)
	GetByID(id int) (transaksies Transaksi, err error)
	// GetAllTransaksisByID(id int) (transaksies []Transaksi, err error)
	Create(transaksi Transaksi) (err error)
	Delete(transaksi Transaksi) (err error)
}

type transaksiRepository struct {
	db *sql.DB
}

// Create implements Repository.
func (p *transaksiRepository) Create(transaksi Transaksi) (err error) {
	// Start a transaction
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Calculate total harga
	var totalHarga int
	for _, detail := range transaksi.Details {
		gameStmt := "SELECT harga FROM " + constant.GameTableName.String() + " WHERE id = $1"

		var price int
		err = tx.QueryRow(gameStmt, detail.GameID).Scan(&price)

		if err != nil {
			return err
		}
		totalHarga += price
	}

	sqlStmt := "INSERT INTO " + constant.TransaksiTableName.String() + "\n" +
		" (tgl_transaksi, user_id, total_harga, created_at, created_by, modified_at, modified_by)" + "\n" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	params := []interface{}{
		transaksi.TglTransakksi,
		transaksi.UserID,
		totalHarga,
		transaksi.CreatedAt,
		transaksi.CreatedBy,
		transaksi.ModifiedAt,
		transaksi.ModifiedBy,
	}

	var transaksiID int
	err = tx.QueryRow(sqlStmt, params...).Scan(&transaksiID)
	if err != nil {

		return err
	}

	// Insert transaksi details
	detailStmt := "INSERT INTO " + constant.TransaksiDetailTableName.String() + "\n" +
		" (transaksi_id, game_id, created_at, created_by, modified_at, modified_by)" + "\n" +
		" VALUES ($1, $2, $3, $4, $5, $6)"

	for _, detail := range transaksi.Details {
		_, err = tx.Exec(detailStmt, transaksiID, detail.GameID, transaksi.CreatedAt, transaksi.CreatedBy, transaksi.ModifiedAt, transaksi.ModifiedBy)
		if err != nil {
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Delete implements Repository.
func (p *transaksiRepository) Delete(transaksi Transaksi) (err error) {
	sqlStmt := "DELETE FROM " + constant.TransaksiTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		transaksi.ID,
	}

	result, err := p.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("delete failed, transaksi not found")
	}

	return nil
}

// GetAll implements Transaksi.
func (p *transaksiRepository) GetAll() (transaksies []Transaksi, err error) {
	sqlStmt := "SELECT t.id, t.tgl_transaksi, t.user_id, t.total_harga, t.created_at, t.created_by, t.modified_at, t.modified_by, " +
		" td.id, td.transaksi_id, td.game_id, g.id, g.title, g.harga, g.category_id, g.platform_id, p.name, p.id, c.name, c.id, u.id, u.username, u.role_id" +
		" FROM " + constant.TransaksiTableName.String() + " AS t " +
		" JOIN " + constant.TransaksiDetailTableName.String() + " AS td " +
		" ON t.id = td.transaksi_id " +
		" JOIN " + constant.GameTableName.String() +
		" AS g ON g.id = td.game_id " +
		" JOIN " + constant.PlatformTableName.String() + " AS p " +
		" ON g.platform_id = p.id " +
		" JOIN " + constant.CategoryTableName.String() + " AS c  " +
		" ON g.platform_id = c.id " +
		" JOIN " + constant.UsersTableName.String() + " AS u " +
		" ON t.user_id = u.id " +
		" ORDER BY t.id ASC"

	rows, err := p.db.Query(sqlStmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var transaksi Transaksi
		var transaksiDetail transaksidetail.TransaksiDetail
		var game game.Game
		var platform platform.Platform
		var category category.Category
		var user users.User

		if err = rows.Scan(&transaksi.ID, &transaksi.TglTransakksi, &transaksi.UserID, &transaksi.TotalHarga, &transaksi.CreatedAt, &transaksi.CreatedBy, &transaksi.ModifiedAt, &transaksi.ModifiedBy, &transaksiDetail.ID, &transaksiDetail.TransaksiID, &transaksiDetail.GameID, &game.ID, &game.Title, &game.Harga, &game.CategoryId, &game.PlatformId, &platform.Name, &platform.ID, &category.Name, &category.ID, &user.ID, &user.Username, &user.RoleId); err != nil {
			return nil, err
		}
		transaksi.User = user
		game.Category = category
		game.Platform = platform
		transaksiDetail.Game = game
		transaksi.Details = append(transaksi.Details, transaksiDetail)
		transaksies = append(transaksies, transaksi)
	}

	return transaksies, nil
}

// GetByID implements Repository.
func (p *transaksiRepository) GetByID(id int) (transaksi Transaksi, err error) {
	// sqlStmt := "SELECT id, tgl_transaksi, user_id, total_harga, created_at, created_by, modified_at, modified_by " + "\n" +
	// 	"FROM " + constant.TransaksiTableName.String() + "\n" +
	// 	"WHERE id = $1"

	sqlStmt := "SELECT t.id, t.tgl_transaksi, t.user_id, t.total_harga, t.created_at, t.created_by, t.modified_at, t.modified_by, " +
		" td.id, td.transaksi_id, td.game_id, g.id, g.title, g.harga, g.category_id, g.platform_id, p.name, p.id, c.name, c.id, u.id, u.username, u.role_id" +
		" FROM " + constant.TransaksiTableName.String() + " AS t " +
		" JOIN " + constant.TransaksiDetailTableName.String() + " AS td " +
		" ON t.id = td.transaksi_id " +
		" JOIN " + constant.GameTableName.String() +
		" AS g ON g.id = td.game_id " +
		" JOIN " + constant.PlatformTableName.String() + " AS p " +
		" ON g.platform_id = p.id " +
		" JOIN " + constant.CategoryTableName.String() + " AS c  " +
		" ON g.platform_id = c.id " +
		" JOIN " + constant.UsersTableName.String() + " AS u " +
		" ON t.user_id = u.id " +
		" WHERE t.id = $1" +
		" ORDER BY t.id ASC"

	params := []interface{}{id}

	rows, err := p.db.Query(sqlStmt, params...)
	if err != nil {
		return transaksi, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaksiDetail transaksidetail.TransaksiDetail
		var game game.Game
		var platform platform.Platform
		var category category.Category
		var user users.User

		if err = rows.Scan(&transaksi.ID, &transaksi.TglTransakksi, &transaksi.UserID, &transaksi.TotalHarga, &transaksi.CreatedAt, &transaksi.CreatedBy, &transaksi.ModifiedAt, &transaksi.ModifiedBy, &transaksiDetail.ID, &transaksiDetail.TransaksiID, &transaksiDetail.GameID, &game.ID, &game.Title, &game.Harga, &game.CategoryId, &game.PlatformId, &platform.Name, &platform.ID, &category.Name, &category.ID, &user.ID, &user.Username, &user.RoleId); err != nil {
			return transaksi, err
		}

		transaksi.User = user
		game.Category = category
		game.Platform = platform
		transaksiDetail.Game = game
		transaksi.Details = append(transaksi.Details, transaksiDetail)
	}

	return transaksi, nil
}

func NewRepository(database *sql.DB) Repository {
	return &transaksiRepository{
		db: database,
	}
}
