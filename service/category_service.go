package service

import (
	"context"
	"customer-api/model/web"
)

type CategoryService interface {
	Create(context context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(context context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(context context.Context, categoryId int)
	FindById(context context.Context, categoryId int) web.CategoryResponse
	FindAll(context context.Context) []web.CategoryResponse
}
