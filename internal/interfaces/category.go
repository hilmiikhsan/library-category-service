package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-category-service/internal/dto"
	"github.com/hilmiikhsan/library-category-service/internal/models"
)

type ICategoryRepository interface {
	InsertNewCategory(ctx context.Context, category *models.Category) error
	FindCategoryByName(ctx context.Context, name string) (*models.Category, error)
}

type ICategoryService interface {
	CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) error
}

type ICategoryHandler interface {
	CreateCategory(*gin.Context)
}
