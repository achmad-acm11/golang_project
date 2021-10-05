package service

import (
	"context"
	"golang-restful-api/model/web"
)

type CategoryService interface {
	GetAllCategory(ctx context.Context) []web.CategoryResponse
	GetOneCategory(ctx context.Context, categoryId int) web.CategoryResponse
	CreateCategory(ctx context.Context, req web.CategoryCreateRequest) web.CategoryResponse
	UpdateCategory(ctx context.Context, req web.CategoryUpdateRequest) web.CategoryResponse
	DeleteCategory(ctx context.Context, categoryId int)
}
