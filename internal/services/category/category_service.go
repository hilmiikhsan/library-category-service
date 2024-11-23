package category

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/library-category-service/constants"
	"github.com/hilmiikhsan/library-category-service/internal/dto"
	"github.com/hilmiikhsan/library-category-service/internal/interfaces"
	"github.com/hilmiikhsan/library-category-service/internal/models"
	"github.com/sirupsen/logrus"
)

type CategoryService struct {
	CategoryRepo interfaces.ICategoryRepository
	Logger       *logrus.Logger
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *dto.CreateCategoryRequest) error {
	categoryData, err := s.CategoryRepo.FindCategoryByName(ctx, req.Name)
	if err != nil {
		s.Logger.Error("category::CreateCategory - failed to find category by name: ", err)
		return err
	}

	if len(categoryData.Name) > 0 {
		s.Logger.Error("category::CreateCategory - category already exists")
		return errors.New(constants.ErrCategoryAlreadyExist)
	}

	err = s.CategoryRepo.InsertNewCategory(ctx, &models.Category{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		s.Logger.Error("category::CreateCategory - failed to insert new category: ", err)
		return err
	}

	return nil
}
