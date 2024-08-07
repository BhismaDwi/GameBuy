package category

import (
	"GameBuy/helpers/constant"
	"database/sql"
	"errors"
)

type Repository interface {
	GetAll() (categories []Category, err error)
	GetByID(id int) (categories Category, err error)
	GetByName(name string) (categories Category, err error)
	// GetAllGamesByID(id int) (categories []Category, err error)
	Create(category Category) (err error)
	Update(category Category) (err error)
	Delete(category Category) (err error)
}

type categoryRepository struct {
	db *sql.DB
}

// Update implements Repository.
func (p *categoryRepository) Update(category Category) (err error) {
	sqlStmt := "UPDATE " + constant.CategoryTableName.String() + "\n" +
		"SET name = $1, modified_at = $2, modified_by = $3 " + "\n" +
		"WHERE id = $4"

	params := []interface{}{
		category.Name,
		category.ModifiedAt,
		category.ModifiedBy,
		category.ID,
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
		return errors.New("update failed, category not found")
	}

	return nil
}

// Create implements Repository.
func (p *categoryRepository) Create(category Category) (err error) {
	sqlStmt := "INSERT INTO " + constant.CategoryTableName.String() + "\n" +
		" (name, created_at, created_by, modified_at, modified_by)" + "\n" +
		" VALUES ($1, $2, $3, $4, $5)"

	params := []interface{}{
		category.Name,
		category.CreatedAt,
		category.CreatedBy,
		category.ModifiedAt,
		category.ModifiedBy,
	}

	_, err = p.db.Exec(sqlStmt, params...)
	if err != nil {
		return
	}

	return nil
}

// Delete implements Repository.
func (p *categoryRepository) Delete(category Category) (err error) {
	sqlStmt := "DELETE FROM " + constant.CategoryTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{
		category.ID,
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

// GetAll implements Category.
func (p *categoryRepository) GetAll() (categories []Category, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by " + "\n" +
		"FROM " + constant.CategoryTableName.String()

	rows, err := p.db.Query(sqlStmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var category Category
		if err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy,
			&category.ModifiedAt, &category.ModifiedBy); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// GetByID implements Repository.
func (p *categoryRepository) GetByID(id int) (category Category, err error) {
	sqlStmt := "SELECT id, name, created_at, created_by, modified_at, modified_by" + "\n" +
		"FROM " + constant.CategoryTableName.String() + "\n" +
		"WHERE id = $1"

	params := []interface{}{id}

	rows, err := p.db.Query(sqlStmt, params...)
	if err != nil {
		return category, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy,
			&category.ModifiedAt, &category.ModifiedBy); err != nil {
			return category, err
		}
	}

	return category, nil
}

func (r *categoryRepository) GetByName(name string) (category Category, err error) {
	sqlStmt := "SELECT id, name FROM " + constant.CategoryTableName.String() + "\n" +
		"WHERE name = $1"

	params := []interface{}{
		name,
	}

	rows, err := r.db.Query(sqlStmt, params...)
	if err != nil {
		return category, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&category.ID, &category.Name); err != nil {
			return category, err
		}
	}
	return category, nil
}

func NewRepository(database *sql.DB) Repository {
	return &categoryRepository{
		db: database,
	}
}
