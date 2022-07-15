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
all: clean mod build

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

# run binary with hot-reloading
dev:
	go run github.com/cosmtrek/air@latest

# generate gofiles
gogen:
	go generate -x ./...

# NOTE: to override test flags, specify `f` such as:
# $ make test-all f="-v -race"
f := -v -cover -race

# check all tests
test-all: test-unit test-integration

# check unit tests
test-unit:
	go test ${f} $$(go list ./... | grep -v "integration")

# check integration tests
test-integration:
	go test ${f} ./integration/...

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
