test:
	go test -v ./... -cover

run:
	go run main.go serve-http

hot:
	@echo " >> Installing gin if not installed"
	@go install github.com/codegangsta/gin@latest
	@gin -i -p 9003 -a 9091 serve-http

goose-create:
# example : make goose-create name=create_users_table
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
ifndef name
	$(error Usage: make goose-create name=<table_name>)
else
	@goose -dir scripts/migrations/sql create $(name) sql
endif

goose-up:
# example : make goose-up
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir scripts/migrations/sql postgres "host=localhost user=postgres password=21012123op dbname=library_category sslmode=disable" up

goose-down:
# example : make goose-down
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir scripts/migrations/sql postgres "host=localhost user=postgres password=21012123op dbname=library_category sslmode=disable" down

goose-status:
# example : make goose-status
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir scripts/migrations/sql postgres "host=localhost user=postgres password=21012123op dbname=library_category sslmode=disable" status

PROTO_SRC_DIR_EXTERNAL := ./external/proto/tokenvalidation
PROTO_OUT_DIR_EXTERNAL := ./external/proto/tokenvalidation
PROTO_FILE_EXTERNAL := token_validation.proto

generate-proto-external:
	protoc --proto_path=$(PROTO_SRC_DIR_EXTERNAL) \
		--go_out=$(PROTO_OUT_DIR_EXTERNAL) --go-grpc_out=$(PROTO_OUT_DIR_EXTERNAL) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_SRC_DIR_EXTERNAL)/$(PROTO_FILE_EXTERNAL)

PROTO_SRC_DIR_INTERNAL := ./cmd/proto/category
PROTO_OUT_DIR_INTERNAL := ./cmd/proto/category
PROTO_FILE_INTERNAL := category.proto

generate-proto-internal:
	protoc --proto_path=$(PROTO_SRC_DIR_INTERNAL) \
		--go_out=$(PROTO_OUT_DIR_INTERNAL) --go-grpc_out=$(PROTO_OUT_DIR_INTERNAL) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_SRC_DIR_INTERNAL)/$(PROTO_FILE_INTERNAL)