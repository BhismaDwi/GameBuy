package game

import (
	"GameBuy/helpers/constant"
	"GameBuy/modules/category"
	"GameBuy/modules/platform"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll() (games []Game, err error)
	GetByID(id int) (games Game, err error)
	GetByTitle(title string) (games Game, err error)
	// GetAllGamesByID(id int) (games []Game, err error)
	Create(game Game) (err error)
	Update(game Game) (err error)
	Delete(game Game) (err error)
	CheckGameExists(id int) (exists bool, err error)
}

type gameRepository struct {
	db *sql.DB
}

func (p *gameRepository) CheckGameExists(id int) (exists bool, err error) {
	sqlStmt := "SELECT EXISTS(SELECT 1 FROM " + constant.GameTableName.String() + " WHERE id = $1)"

	err = p.db.QueryRow(sqlStmt, id).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// Update implements Repository.
func (p *gameRepository) Update(game Game) (err error) {
	sqlStmt := "UPDATE " + constant.GameTableName.String() + "\n" +
		"SET title = $1, harga = $2, category_id = $3, platform_id=$4, modified_at = $5, modified_by = $6 " + "\n" +
		"WHERE id = $7"

	params := []interface{}{
		game.Title,
		game.Harga,
		game.CategoryId,
		game.PlatformId,
		game.ModifiedAt,
		game.ModifiedBy,
		game.ID,
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
		return errors.New("update failed, game not found")
	}

	return nil
}

// Create implements Repository.
func (p *gameRepository) Create(game Game) (err error) {
	sqlStmt := "INSERT INTO " + constant.GameTableName.String() + "\n" +
		" (title, harga, category_id, platform_id, created_at, created_by, modified_at, modified_by)" + "\n" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	params := []interface{}{
		game.Title,
		game.Harga,
		game.CategoryId,
		game.PlatformId,
		game.CreatedAt,
		game.CreatedBy,
		game.ModifiedAt,
		game.ModifiedBy,
	}

	_, err = p.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

// Delete implements Repository.
func (p *gameRepository) Delete(game Game) (err error) {
	sqlStmt := "DELETE FROM " + constant.GameTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		game.ID,
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
		return errors.New("delete failed, game not found")
	}

	return nil
}

// GetAll implements Game.
func (p *gameRepository) GetAll() (games []Game, err error) {
	sqlStmt := "SELECT game.id, game.title, game.harga, game.category_id, game.platform_id, game.created_at, game.created_by, " +
		" game.modified_at, game.modified_by, platform.name, platform.id, category.name, category.id " +
		" FROM " + constant.GameTableName.String() + " AS game " +
		" JOIN " + constant.PlatformTableName.String() + " AS platform " +
		" ON " + constant.GameTableName.String() + ".platform_id = " + constant.PlatformTableName.String() + ".id " +
		" JOIN " + constant.CategoryTableName.String() + " AS category  " +
		" ON " + constant.GameTableName.String() + ".platform_id = " + constant.CategoryTableName.String() + ".id"

	rows, err := p.db.Query(sqlStmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var game Game
		var platform platform.Platform
		var category category.Category
		if err = rows.Scan(&game.ID, &game.Title, &game.Harga, &game.CategoryId, &game.PlatformId, &game.CreatedAt, &game.CreatedBy,
			&game.ModifiedAt, &game.ModifiedBy, &platform.Name, &platform.ID, &category.Name, &category.ID); err != nil {
			return nil, err
		}
		game.Platform = platform
		game.Category = category
		games = append(games, game)
	}

	return games, nil
}

// GetByID implements Repository.
func (p *gameRepository) GetByID(id int) (game Game, err error) {
	// sqlStmt := "SELECT id, title, harga, category_id, platform_id, created_at, created_by, modified_at, modified_by" + "\n" +
	// 	"FROM " + constant.GameTableName.String() + "\n" +
	// 	"WHERE id = $1"
	sqlStmt := "SELECT game.id, game.title, game.harga, game.category_id, game.platform_id, game.created_at, game.created_by, " +
		" game.modified_at, game.modified_by, platform.name, platform.id, category.name, category.id " +
		" FROM " + constant.GameTableName.String() + " AS game " +
		" JOIN " + constant.PlatformTableName.String() + " AS platform " +
		" ON " + constant.GameTableName.String() + ".platform_id = " + constant.PlatformTableName.String() + ".id " +
		" JOIN " + constant.CategoryTableName.String() + " AS category  " +
		" ON " + constant.GameTableName.String() + ".platform_id = " + constant.CategoryTableName.String() + ".id" +
		" WHERE game.id = $1"

	params := []interface{}{id}

	rows, err := p.db.Query(sqlStmt, params...)
	if err != nil {
		return game, err
	}
	defer rows.Close()

	for rows.Next() {
		var platform platform.Platform
		var category category.Category
		if err = rows.Scan(&game.ID, &game.Title, &game.Harga, &game.CategoryId, &game.PlatformId, &game.CreatedAt, &game.CreatedBy,
			&game.ModifiedAt, &game.ModifiedBy, &platform.Name, &platform.ID, &category.Name, &category.ID); err != nil {
			return game, err
		}
		game.Platform = platform
		game.Category = category
	}

	return game, nil
}

func (r *gameRepository) GetByTitle(title string) (game Game, err error) {
	sqlStmt := "SELECT id, title FROM " + constant.GameTableName.String() + "\n" +
		"WHERE title = $1"

	params := []interface{}{
		title,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return game, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&game.ID, &game.Title); err != nil {
			return game, err
		}
	}
	return game, nil
}

func NewRepository(database *sql.DB) Repository {
	return &gameRepository{
		db: database,
	}
}
