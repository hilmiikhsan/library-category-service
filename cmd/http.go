package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hilmiikhsan/library-category-service/helpers"
	"github.com/sirupsen/logrus"
)

func ServeHTTP() {
	// dependency := dependencyInject()

	router := gin.Default()

	// router.GET("/health", dependency.HealthcheckAPI.HealthcheckHandlerHTTP)

	err := router.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		helpers.Logger.Fatal("failed to run http server: ", err)
	}
}

type Dependency struct {
	Logger *logrus.Logger

	// HealthcheckAPI interfaces.IHealthcheckHandler
}

func dependencyInject() Dependency {
	helpers.SetupLogger()

	// healthcheckSvc := &healthCheckServices.Healthcheck{}
	// healthcheckAPI := &healthCheckAPI.Healthcheck{
	// 	HealthcheckServices: healthcheckSvc,
	// }

	// validator := validator.NewValidator()

	return Dependency{
		Logger: helpers.Logger,
	}
}
