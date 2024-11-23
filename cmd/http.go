package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-category-service/external"
	"github.com/hilmiikhsan/library-category-service/helpers"
	categoryAPI "github.com/hilmiikhsan/library-category-service/internal/api/category"
	healthCheckAPI "github.com/hilmiikhsan/library-category-service/internal/api/health_check"
	"github.com/hilmiikhsan/library-category-service/internal/interfaces"
	categoryRepository "github.com/hilmiikhsan/library-category-service/internal/repository/category"
	categoryServices "github.com/hilmiikhsan/library-category-service/internal/services/category"
	healthCheckServices "github.com/hilmiikhsan/library-category-service/internal/services/health_check"
	"github.com/hilmiikhsan/library-category-service/internal/validator"
	"github.com/sirupsen/logrus"
)

func ServeHTTP() {
	dependency := dependencyInject()

	router := gin.Default()

	router.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	categoryV1 := router.Group("/category/v1")
	categoryV1.POST("/create", dependency.MiddlewareValidateToken, dependency.CategoryAPI.CreateCategory)
	categoryV1.GET("/:id", dependency.MiddlewareValidateToken, dependency.CategoryAPI.GetDetailCategory)
	categoryV1.GET("/", dependency.MiddlewareValidateToken, dependency.CategoryAPI.GetListCategory)
	categoryV1.PUT("/update", dependency.MiddlewareValidateToken, dependency.CategoryAPI.UpdateCategory)
	categoryV1.DELETE("/:id", dependency.MiddlewareValidateToken, dependency.CategoryAPI.DeleteCategory)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		helpers.Logger.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	Logger             *logrus.Logger
	CategoryRepository interfaces.ICategoryRepository

	HealthcheckAPI interfaces.IHealthcheckHandler
	CategoryAPI    interfaces.ICategoryHandler
	External       interfaces.IExternal
}

func dependencyInject() Dependency {
	helpers.SetupLogger()

	healthcheckSvc := &healthCheckServices.Healthcheck{}
	healthcheckAPI := &healthCheckAPI.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	categoryRepo := &categoryRepository.CategoryRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	categorySvc := &categoryServices.CategoryService{
		CategoryRepo: categoryRepo,
		Logger:       helpers.Logger,
	}
	categoryAPI := &categoryAPI.CategoryHandler{
		CategoryService: categorySvc,
		Validator:       validator,
	}

	external := &external.External{
		Logger: helpers.Logger,
	}

	return Dependency{
		Logger:             helpers.Logger,
		CategoryRepository: categoryRepo,
		HealthcheckAPI:     healthcheckAPI,
		CategoryAPI:        categoryAPI,
		External:           external,
	}
}
