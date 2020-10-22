# basesvc
Simple microservice template in Go with Clean Architecture


#### Run the Applications
Here is the steps to run it with `docker-compose`

```bash
#move to directory
$ mkdir -p $GOPATH/src/github.com/iwanjunaid 
$ cd $GOPATH/src/github.com/iwanjunaid 

# Clone into YOUR $GOPATH/src
$ git clone https://github.com/iwanjunaid/basesvc.git

#move to project
$ cd basesvc

# Build the docker image first
$ make docker

# Run the application
$ make run

# check if the containers are running
$ docker ps

# Execute the call
$ curl http://localhost:9090/authors

# Stop
$ make stop
```