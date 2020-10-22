BINARY=basesvc
test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} main.go

dependencies:
	@echo "> Installing the server dependencies ..."
	@go mod tidy -v
	@go install github.com/swaggo/swag/cmd/swag

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
	swag init -g infrastructure/router/router.go

.PHONY: clean install unittest lint-prepare lint docs engine test dependencies