BINARY=basesvc
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} main.go

dev: setup run

dependencies:
	@echo "> Installing the server dependencies ..."
	@go mod tidy -v
	@go install github.com/swaggo/swag/cmd/swag
	@go get -u github.com/cosmtrek/air

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...
 
docs:
	@echo "> Generate Swagger Docs"
	@if ! command -v swag &> /dev/null; then go install github.com/swaggo/swag/cmd/swag ; fi
	@swag init -g infrastructure/rest/rest.go

docker:
	@echo "> Build Docker image [PRODUCTION]"
	@docker build -t basesvc -f build/Dockerfile . 

run:
	@echo "> Run docker-compose"
	@docker-compose -f deployments/docker-compose.yml up --build -d

setup:
	@cp config/example/mysql.yml.example config/db/mysql.yml
	@cp config/example/rest.yml.example config/server/rest.yml
	@cp config/example/logger.yml.example config/logging/logger.yml


stop:
	@echo "> Stop docker-compose"
	@docker-compose -f deployments/docker-compose.yml down

.PHONY: clean install unittest lint-prepare lint docs engine dev test setup dependencies run stop