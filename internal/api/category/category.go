package category

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-category-service/constants"
	"github.com/hilmiikhsan/library-category-service/helpers"
	"github.com/hilmiikhsan/library-category-service/internal/dto"
	"github.com/hilmiikhsan/library-category-service/internal/interfaces"
	"github.com/hilmiikhsan/library-category-service/internal/validator"
)

type CategoryHandler struct {
	CategoryService interfaces.ICategoryService
	Validator       *validator.Validator
}

func (api *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var (
		req = new(dto.CreateCategoryRequest)
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.Logger.Error("handler::CreateCategory - Failed to bind request : ", err)
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrFailedBadRequest))
		return
	}

	if err := api.Validator.Validate(req); err != nil {
		helpers.Logger.Error("handler::CreateCategory - Failed to validate request : ", err)
		code, errs := helpers.Errors(err, req)
		ctx.JSON(code, helpers.Error(errs))
		return
	}

	err := api.CategoryService.CreateCategory(ctx.Request.Context(), req)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrCategoryAlreadyExist) {
			helpers.Logger.Error("handler::CreateCategory - Category already exist")
			ctx.JSON(http.StatusConflict, helpers.Error(constants.ErrCategoryAlreadyExist))
			return
		}

		helpers.Logger.Error("handler::CreateCategory - Failed to create category : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Success(nil, ""))
}
