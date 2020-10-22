BINARY=basesvc
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} main.go

dev:
	@echo "> Run Development Env"
	@air -c .air.toml

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
	swag init -g infrastructure/rest/rest.go

docker:
	@echo "> Build Docker image"
	@docker build -t basesvc -f build/Dockerfile . 

run:
	@echo "> Run docker-compose"
	@docker-compose -f deployments/docker-compose.yml -f deployments/docker-compose.mysql.yml up --build -d

.PHONY: clean install unittest lint-prepare lint docs engine dev test dependencies