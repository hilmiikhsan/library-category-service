package category

import (
	"net/http"
	"strconv"
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

func (api *CategoryHandler) GetDetailCategory(ctx *gin.Context) {
	var (
		id = ctx.Param("id")
	)

	if id == "" {
		helpers.Logger.Error("handler::GetDetailCategory - Missing required parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error("missing required parameter: id"))
		return
	}

	if !helpers.IsValidUUID(id) {
		helpers.Logger.Error("handler::GetDetailCategory - Invalid UUID format for parameter: id")
		ctx.JSON(http.StatusBadRequest, helpers.Error(constants.ErrParamIdIsRequired))
		return
	}

	res, err := api.CategoryService.GetDetailCategory(ctx.Request.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrCategoryNotFound) {
			helpers.Logger.Error("handler::GetDetailCategory - Category not found")
			ctx.JSON(http.StatusNotFound, helpers.Error(constants.ErrCategoryNotFound))
			return
		}

		helpers.Logger.Error("handler::GetDetailCategory - Failed to get category detail : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}

func (api *CategoryHandler) GetListCategory(ctx *gin.Context) {
	pageIndexStr := ctx.Query("page")
	pageSizeStr := ctx.Query("limit")

	pageIndex, _ := strconv.Atoi(pageIndexStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if pageIndex <= 0 {
		pageIndex = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	res, err := api.CategoryService.GetListCategory(ctx.Request.Context(), pageSize, pageIndex)
	if err != nil {
		helpers.Logger.Error("handler::GetListCategory - Failed to get list category : ", err)
		ctx.JSON(http.StatusInternalServerError, helpers.Error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(res, ""))
}
