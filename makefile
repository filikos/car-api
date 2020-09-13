.PHONY: help up down tests

help: 
	@echo "Available targets:"
	@echo "up		info: starts Rest-API & docker containers"
	@echo "down		info: stops and removes Rest-API & docker containers"
	@echo "test		info: runs all existing tests"
up:
	docker-compose up

down:
	docker-compose down

tests:
	@echo "Starting PostgreSQL docker container"
	docker run -d \
		--name dev \
		--env-file $${HOME}/go/src/workspace-go/coding-challenge/car-api/testdata/dbConfigTest.env \
		-p 5432:5432 \
		 postgres
	
	
	@echo "Runs all tests including integration tests."
	-go test ./... --tags=integration -failfast -v 

	@echo "Stop and remove PostgreSQL docker container"
	docker stop `docker ps -aqf "name=dev"`
	docker rm `docker ps -aqf "name=dev"`