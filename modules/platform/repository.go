package platform

import (
	"GameBuy/helpers/constant"
	"database/sql"
)

type Repository interface {
	GetAll() (platforms []Platform, err error)
	GetByID(id int) (platforms Platform, err error)
	// GetAllGamesByID(id int) (platforms []Platform, err error)
	// Create(platform) (Platforms platforms, err error)
	// Update(id int, platform) err error
	// Delete(id int) err error
}

type platformRepository struct {
	db *sql.DB
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
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by " + "\n" +
		"FROM " + constant.PlatformTableName.String()

	params := []interface{}{id}

	rows, err := p.db.Query(sqlStmt, params...)

	if err != nil {
		return platform, err
	}

	defer rows.Close()
	for rows.Next() {
		var platform Platform
		if err = rows.Scan(&platform.ID, &platform.Name, &platform.CreatedAt, &platform.CreatedBy,
			&platform.ModifiedAt, &platform.ModifiedBy); err != nil {
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
