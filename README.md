# basesvc
Simple microservice template in Go with Clean Architecture


#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
# Move to directory
$ mkdir -p $GOPATH/src/github.com/iwanjunaid 
$ cd $GOPATH/src/github.com/iwanjunaid 

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/iwanjunaid/basesvc.git

# Move to project
$ cd basesvc

# Setup config and edit config 
$ make setup

# run the application in development env
$ make dev

# or run the application in production env
$ make prod

# check if the containers are running
$ docker ps

# Execute the call
$ curl http://localhost:8080/authors

# Stop
$ make stop
```


## Development
1. Create config file and configure
```bash
cp .env.dist .env
cp basesvc.config.json.dist basesvc.config.json
```
2. Run application
```bash
go run main.go api
