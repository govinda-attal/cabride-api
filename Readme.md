# Cab Ride API
Sample implementation if for a code test.
### Preface

Cab Data Researcher is a company that provides insights on the open data about NY cab trips. 
Cab trips in NY are public available as csv downloadable files. In order to make it more useful we want to wrap the data in a public API.

 
Data format is as follow:
medallion, hack_license, vendor_id, rate_code, store_and_fwd_flag, pickup_datetime, dropoff_datetime, passenger_count, trip_time_in_secs, trip_distance

1. The medallion is the cab identification.
2. API should provide a way to query how many trips a particular cab (medallion) has made given a particular pickup date ( using pickup_datetime and only considering the date part)
3. The API must receive one or more medallions and return how many trips each medallion has made.

 
Considering that the query creates a heavy load on the database, the results must be cached. The API must allow user to ask for fresh data, ignoring the cache. There must be also be a method to clear the cache.

### Implementation Specifics

A simple API with two endpoints.

The first endpoint, POST /v1/trips/fetch/2006-01-02 handles receipt of list of medallions in json format and queries rides information for given pickup date either in redis cache or MySQL database.
```
{
  "medallions": [
    "D7D598CD99978BD012A87A76A7C891B7"
  ]
}
```

The second endpoint DELETE /v1/cache/clear?pickup=2013-12-01 clears entries within cache for given date. If no date is passed it will clear all entries within the cache.

### Design
1. API is implementation as GRPC Microservice
2. With grpc-ecosystem GRPC Microservice is wrapped with light-weight OpenAPI HTTP/JSON Restful Gateway
3. MySQL Database holds data for the NY trips
4. Redis is used to Cache trips data

* Note: Preferrably protobug release 3.6.0 is installed on your machine on your OS: https://github.com/google/protobuf/releases 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

1. Linux OS - Ubuntu or Fedora is preferred.
2. git
3. Golang Setup locally on OS
4. Postman for REST API testing
5. Docker & Docker Compose

### Installing

This application can be setup locally in following ways:

#### Option A
```
go get github.com/govinda-attal/cabride-api
```

#### Option B (Preferred) :heavy_check_mark:
```
cd $GOPATH/src/github.com/
mkdir govinda-attal
cd govinda-attal
git clone https://github.com/govinda-attal/cabride-api.git
```

### Application Development Setup

'Makefile' will be used to setup cabride-api quickly on the development workstation.

```
cd $GOPATH/src/github.com/govinda-attal/cabride-api
make install # This will go get 'dep' and use it to install vendor dependencies.
```

## Running Tests

### Unit tests

Sample BDD Unit tests are implemented using ginkgo and gomega. These unit tests are written for HTTP handler only.
Three Unit tests are written for this code test.

```
cd $GOPATH/src/github.com/govinda-attal/cabride-api
make test # This will run execute unit tests with ginkgo -r command
```

### Integration tests

This microservice meets given requirements with Golang and MySQL Database as backend and Redis for cache. To keep this foot-print of this application minimum mysql db & cache will execute within a docker container. Where as following backend microservice can be hosted within a docker container or local OS.

#### Option A: Docker Compose - orchestrate DB and Microservice as docker containers (Preferred) :heavy_check_mark:
* Note: For some reason Microservice container port 9080 though nicely mapped to HOST OS - it is not accessible. So Use Option B.
```
cd $GOPATH/src/github.com/govinda-attal/cabride-api
docker-compose up -d # This will start MySQL DB, Cabride Microservice and Swagger-UI which will point to CabRide microservice swagger definition.
```

Docker compose will orchestrate containers and they can be accessed from Local OS as below:
1. MySQL DB on localhost:4406
2. Adminer on :earth_asia: http://localhost:8080
2. Microservice on :earth_asia: http://localhost:9080
3. Swagger-UI on :earth_asia: http://localhost:9090


#### Option B: MySQL, Redis, SwaggerUI & Adminer as Docker containers but Microservice is run locally on your OS

```
cd $GOPATH/src/github.com/govinda-attal/cabride-api
make local-providers-start # This will host MySQL DB, RediS, SwaggerUI & Adminer as docker containers.
make local-serve # This will build the Microservice and run it locally and the API is exposed on port 9080 
```

### Import NY Cab Rides Data in DB
Using Adminer UI import data in MYSQL DB from ./migrations/ny_cab_data_cab_trip_data_full.sql

### Postman Test Collection

Start Postman and import sample test collection at
``` 
$GOPATH/src/github.com/govinda-attal/cabride-api/test/fixtures/cabride-api.postman_collection.json
```

### Swagger UI to view and trial Microservice Open API

Post running command *docker-compose up -d* use browser to open :earth_asia: http://localhost:9090

## Cleanup

For containers orchestrated by Docker-compose
```
cd $GOPATH/src/github.com/govinda-attal/cabride-api
docker-compose down
docker image rm gattal/cabride-api:latest 
docker image rm mysql:5.7.22
doocker image rm redis:latest
docker image rm swaggerapi/swagger-ui:latest
docker image rm adminer:latest
```

In case when Microservice running locally

Press Ctrl+C on the terminal on which microservice was running.


To delete the source code for this microservice run :skull: cd $GOPATH/src/github.com/ && rm -rf  govinda-attal :skull: 

## Authors

* [Govinda Attal](https://github.com/govinda-attal)

## Acknowledgments

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
* [Organizing Database Access](https://www.alexedwards.net/blog/organising-database-access)
* [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
* [Dependecy Injection, Mocking & TDD/BDD in Go](https://www.youtube.com/watch?v=uFXfTXSSt4I)
* [Ginkgo - A Golang BDD Testing Framework](https://onsi.github.io/ginkgo/)
* [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
* [grpc-ecosystem/grpc-Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
* [grpc with Go](https://grpc.io/docs/quickstart/go.html)

