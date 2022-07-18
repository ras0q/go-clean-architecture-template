SHELL    := /bin/bash
APP_NAME := myapp

MYSQL_CONTAINER_NAME := ${APP_NAME}_mysql
MYSQL_USER := root
MYSQL_PASS := password
MYSQL_HOST := localhost
MYSQL_PORT := 3306
MYSQL_NAME := myapp

# all tasks run as phony
.PHONY: ${shell egrep -o ^[a-zA-Z_-]+: ./Makefile | sed 's/://'}

# default task
all: clean gogen mod build

# remove binary
clean:
	rm -f ${APP_NAME}
	go clean

# install dependencies
mod:
	go mod download

# compile binary
build:
	go build -o ${APP_NAME}

# run binary with live reload
dev:
	go run github.com/cosmtrek/air@latest

# generate gofiles
gogen:
	go generate -x ./...

# NOTE: to override test flags, specify `f` such as:
# $ make test-all f="-v -race"
f := -v -cover -coverpkg ./... -race

# check all tests
test-all: test-unit test-integration

# check unit tests
test-unit:
	go test ${f} $$(go list ./... | grep -v "integration")

# check integration tests
test-integration:
	go test ${f} ./integration/...

# measure coverage of all tests
cover-all: cover-unit cover-integration

# measure coverage of unit tests
cover-unit:
	go test -coverprofile ./dev/tmp/cover_unit.out $$(go list ./... | grep -v "integration")
	go tool cover -html ./dev/tmp/cover_unit.out -o ./dev/tmp/cover_unit.html

# measure coverage of integration tests
cover-integration:
	go test -coverprofile ./dev/tmp/cover_integration.out -coverpkg ./... ./integration/...
	go tool cover -html ./dev/tmp/cover_integration.out -o ./dev/tmp/cover_integration.html

# run linters
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run ./... --fix

db-up:
	docker run -d --rm -it \
		--name ${MYSQL_CONTAINER_NAME} \
		-p ${MYSQL_PORT}:3306 \
		-e MYSQL_ROOT_PASSWORD=${MYSQL_PASS} \
		-e MYSQL_DATABASE=${MYSQL_NAME} \
		mysql:latest

db-down:
	docker stop ${MYSQL_CONTAINER_NAME}
	docker rm ${MYSQL_CONTAINER_NAME}
