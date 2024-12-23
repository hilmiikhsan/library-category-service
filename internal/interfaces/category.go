package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-category-service/cmd/proto/category"
	"github.com/hilmiikhsan/library-category-service/internal/dto"
	"github.com/hilmiikhsan/library-category-service/internal/models"
)

type ICategoryRepository interface {
	InsertNewCategory(ctx context.Context, category *models.Category) error
	FindCategoryByName(ctx context.Context, name string) (*models.Category, error)
	FindCategoryByID(ctx context.Context, id string) (*models.Category, error)
	FindAllCategory(ctx context.Context, limit, offset int) ([]models.Category, error)
	UpdateNewCategory(ctx context.Context, category *models.Category) error
	DeleteCategoryByID(ctx context.Context, id string) error
}

type ICategoryService interface {
	CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) error
	GetDetailCategory(ctx context.Context, id string) (*dto.GetDetailCategoryResponse, error)
	GetListCategory(ctx context.Context, limit, offset int) (*dto.GetListCategoryResponse, error)
	UpdateCategory(ctx context.Context, req *dto.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, id string) error
}

type ICategoryHandler interface {
	CreateCategory(*gin.Context)
	GetDetailCategory(*gin.Context)
	GetListCategory(*gin.Context)
	UpdateCategory(*gin.Context)
	DeleteCategory(*gin.Context)
}

type ICategoryAPI interface {
	GetDetailCategory(ctx context.Context, req *category.CategoryRequest) (*category.CategoryResponse, error)
}
