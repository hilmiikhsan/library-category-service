package cmd

import (
	"net"

	"github.com/hilmiikhsan/library-category-service/cmd/proto/category"
	"github.com/hilmiikhsan/library-category-service/helpers"
	api "github.com/hilmiikhsan/library-category-service/internal/grpc"
	categoryRepository "github.com/hilmiikhsan/library-category-service/internal/repository/category"
	categoryServices "github.com/hilmiikhsan/library-category-service/internal/services/category"
	"github.com/hilmiikhsan/library-category-service/internal/validator"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	dependencyGrpc := dependencyGrpcInject()

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "6001"))
	if err != nil {
		helpers.Logger.Fatal("failed to listen grpc port: ", err)
	}

	server := grpc.NewServer()
	category.RegisterCategoryServiceServer(server, dependencyGrpc.CategoryAPI)

	helpers.Logger.Info("start listening grpc on port:" + helpers.GetEnv("GRPC_PORT", "6001"))
	if err := server.Serve(lis); err != nil {
		helpers.Logger.Fatal("failed to serve grpc port: ", err)
	}
}

type DependencyGrpc struct {
	Logger      *logrus.Logger
	CategoryAPI *api.CategoryAPI
}

func dependencyGrpcInject() *DependencyGrpc {
	categoryRepo := &categoryRepository.CategoryRepository{
		DB:     helpers.DB,
		Logger: helpers.Logger,
	}

	validator := validator.NewValidator()

	categorySvc := &categoryServices.CategoryService{
		CategoryRepo: categoryRepo,
		Logger:       helpers.Logger,
	}
	categoryAPI := &api.CategoryAPI{
		CategoryService: categorySvc,
		Validator:       validator,
	}

	return &DependencyGrpc{
		Logger:      helpers.Logger,
		CategoryAPI: categoryAPI,
	}
}
