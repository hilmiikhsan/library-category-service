package grpc

import (
	"context"
	"strings"

	"github.com/hilmiikhsan/library-category-service/cmd/proto/category"
	"github.com/hilmiikhsan/library-category-service/constants"
	"github.com/hilmiikhsan/library-category-service/helpers"
	"github.com/hilmiikhsan/library-category-service/internal/dto"
	"github.com/hilmiikhsan/library-category-service/internal/interfaces"
	"github.com/hilmiikhsan/library-category-service/internal/validator"
)

type CategoryAPI struct {
	CategoryService interfaces.ICategoryService
	Validator       *validator.Validator
	category.UnimplementedCategoryServiceServer
}

func (api *CategoryAPI) GetDetailCategory(ctx context.Context, req *category.CategoryRequest) (*category.CategoryResponse, error) {
	internalReq := dto.GetDetailCategoryRequest{
		ID: req.Id,
	}

	if err := api.Validator.Validate(internalReq); err != nil {
		helpers.Logger.Error("api::GetDetailCategory - Failed to validate request : ", err)
		return &category.CategoryResponse{
			Message: "Failed to validate request",
			Data:    nil,
		}, nil
	}

	res, err := api.CategoryService.GetDetailCategory(ctx, internalReq.ID)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrCategoryNotFound) {
			helpers.Logger.Error("api::GetDetailCategory - Category not found")
			return &category.CategoryResponse{
				Message: constants.ErrCategoryNotFound,
				Data:    nil,
			}, nil
		}

		helpers.Logger.Error("api::GetDetailCategory - Failed to get detail category : ", err)
		return &category.CategoryResponse{
			Message: "Failed to get detail category",
			Data:    nil,
		}, nil
	}

	return &category.CategoryResponse{
		Message: constants.SuccessMessage,
		Data: &category.CategoryData{
			Id:   res.ID,
			Name: res.Name,
		},
	}, nil
}
