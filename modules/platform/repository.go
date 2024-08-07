package platform

import (
	"GameBuy/helpers/constant"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll() (platforms []Platform, err error)
	GetByID(id int) (platforms Platform, err error)
	GetByName(name string) (platforms Platform, err error)
	// GetAllGamesByID(id int) (platforms []Platform, err error)
	Create(platform Platform) (err error)
	Update(platform Platform) (err error)
	Delete(platform Platform) (err error)
}

type platformRepository struct {
	db *sql.DB
}

// Update implements Repository.
func (p *platformRepository) Update(platform Platform) (err error) {
	sqlStmt := "UPDATE " + constant.PlatformTableName.String() + "\n" +
		"SET name = $1, modified_at = $2, modified_by = $3 " + "\n" +
		"WHERE id = $4"

	params := []interface{}{
		platform.Name,
		platform.ModifiedAt,
		platform.ModifiedBy,
		platform.ID,
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
		return errors.New("update failed, platform not found")
	}

	return nil
}

// Create implements Repository.
func (p *platformRepository) Create(platform Platform) (err error) {
	sqlStmt := "INSERT INTO " + constant.PlatformTableName.String() + "\n" +
		" (name, created_at, created_by, modified_at, modified_by)" + "\n" +
		" VALUES ($1, $2, $3, $4, $5)"

	params := []interface{}{
		platform.Name,
		platform.CreatedAt,
		platform.CreatedBy,
		platform.ModifiedAt,
		platform.ModifiedBy,
	}

	_, err = p.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

// Delete implements Repository.
func (p *platformRepository) Delete(platform Platform) (err error) {
	sqlStmt := "DELETE FROM " + constant.PlatformTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		platform.ID,
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
		return errors.New("delete failed, category not found")
	}

	return nil
}

// GetAll implements Platform.
func (p *platformRepository) GetAll() (platforms []Platform, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by " + "\n" +
		"FROM " + constant.PlatformTableName.String()

	rows, err := p.db.Query(sqlStmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var platform Platform
		if err = rows.Scan(&platform.ID, &platform.Name, &platform.CreatedAt, &platform.CreatedBy,
			&platform.ModifiedAt, &platform.ModifiedBy); err != nil {
			return nil, err
		}
		platforms = append(platforms, platform)
	}

	return platforms, nil
}

// GetByID implements Repository.
func (p *platformRepository) GetByID(id int) (platform Platform, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by" + "\n" +
		"FROM " + constant.PlatformTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{id}

	rows, err := p.db.Query(sqlStmt, params...)
	if err != nil {
		return platform, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&platform.ID, &platform.Name, &platform.CreatedAt, &platform.CreatedBy,
			&platform.ModifiedAt, &platform.ModifiedBy); err != nil {
			return platform, err
		}
	}

	return platform, nil
}

func (r *platformRepository) GetByName(name string) (platform Platform, err error) {
	sqlStmt := "SELECT id, name FROM " + constant.PlatformTableName.String() + "\n" +
		"WHERE name = $1"

	params := []interface{}{
		name,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return platform, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&platform.ID, &platform.Name); err != nil {
			return platform, err
		}
	}
	return platform, nil
}

func NewRepository(database *sql.DB) Repository {
	return &platformRepository{
		db: database,
	}
}
