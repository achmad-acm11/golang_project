package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}
func (c *CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error) {
	SQL := "SELECT * FROM categories"

	data, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer data.Close()

	var categories []domain.Category
	for data.Next() {
		category := domain.Category{}

		data.Scan(&category.Id, &category.Name)

		categories = append(categories, category)
	}
	return categories, nil
}
func (c *CategoryRepositoryImpl) GetOne(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT * FROM categories WHERE id = ?"

	data, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer data.Close()

	category := domain.Category{}
	if data.Next() {
		data.Scan(&category.Id, &category.Name)
		return category, nil
	} else {
		return category, errors.New("not found")
	}
}
func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "INSERT INTO categories (name) VALUES (?)"

	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	categoryId, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(categoryId)
	return category, nil
}
func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "UPDATE categories SET name=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category, nil
}
func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) error {
	SQL := "DELETE FROM categories WHERE id = ?"

	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)

	return nil
}
