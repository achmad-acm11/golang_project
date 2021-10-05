package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/domain"
)

type CategoryRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) ([]domain.Category, error)
	GetOne(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) error
}
