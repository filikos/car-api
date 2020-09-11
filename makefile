.PHONY: up down test

up:
	docker-compose up

down:
	docker-compose down

test:
	docker run -d \
		--name dev \
		--env-file $${HOME}/go/src/workspace-go/coding-challenge/car-api/testdata/dbConfigTest.env \
		-p 5432:5432 \
		 postgres
	
	# runs all tests including integration tests.
	-go test ./... --tags=integration -failfast -v 

	# stop and removes container
	docker stop `docker ps -aqf "name=dev"`
	docker rm `docker ps -aqf "name=dev"`