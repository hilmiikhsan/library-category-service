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

func (s *CategoryService) GetDetailCategory(ctx context.Context, id string) (*dto.GetDetailCategoryResponse, error) {
	categoryData, err := s.CategoryRepo.FindCategoryByID(ctx, id)
	if err != nil {
		s.Logger.Error("category::GetCategoryDetail - failed to find category by id: ", err)
		return nil, err
	}

	return &dto.GetDetailCategoryResponse{
		ID:          categoryData.ID.String(),
		Name:        categoryData.Name,
		Description: categoryData.Description,
	}, nil
}

func (s *CategoryService) GetListCategory(ctx context.Context, limit, offset int) (*dto.GetListCategoryResponse, error) {
	pageSize := limit
	pageIndex := (offset - 1) * limit

	categoryData, err := s.CategoryRepo.FindAllCategory(ctx, pageSize, pageIndex)
	if err != nil {
		s.Logger.Error("category::GetListCategory - failed to find all category: ", err)
		return nil, err
	}

	categories := make([]dto.Category, 0)
	for _, category := range categoryData {
		categories = append(categories, dto.Category{
			ID:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
		})
	}

	pagination := dto.Pagination{
		Page:  offset,
		Limit: limit,
	}

	response := &dto.GetListCategoryResponse{
		CategoryList: categories,
		Pagination:   pagination,
	}

	return response, nil
}
