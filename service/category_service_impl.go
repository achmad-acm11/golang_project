package service

import (
	"context"
	"database/sql"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	repo     repository.CategoryRepository
	db       *sql.DB
	validate *validator.Validate
}

func NewCategoryService(repo repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		repo:     repo,
		db:       db,
		validate: validate,
	}
}
func (c *CategoryServiceImpl) GetAllCategory(ctx context.Context) []web.CategoryResponse {
	tx, err := c.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := c.repo.GetAll(ctx, tx)
	helper.PanicIfError(err)

	return helper.ToCategoriesResponse(categories)
}
func (c *CategoryServiceImpl) GetOneCategory(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.repo.GetOne(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}
func (c *CategoryServiceImpl) CreateCategory(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse {
	err := c.validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := c.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: req.Name,
	}
	category, err = c.repo.Create(ctx, tx, category)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
func (c *CategoryServiceImpl) UpdateCategory(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := c.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.repo.GetOne(ctx, tx, req.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = req.Name

	category, err = c.repo.Update(ctx, tx, category)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
func (c *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) {
	tx, err := c.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.repo.GetOne(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = c.repo.Delete(ctx, tx, category)
	helper.PanicIfError(err)
}
