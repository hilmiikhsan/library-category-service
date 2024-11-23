package main

import (
	"github.com/hilmiikhsan/library-category-service/cmd"
	"github.com/hilmiikhsan/library-category-service/helpers"
)

func main() {
	// load config
	helpers.SetupConfig()

	// load log
	helpers.SetupLogger()

	// load db
	helpers.SetupPostgres()

	// run grpc
	go cmd.ServeGRPC()

	// run http
	cmd.ServeHTTP()
}
