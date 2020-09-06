# Car-management-api
This Rest-API provides basic functionality to handle car related data. The API is written in Go and is using PostgreSQL within Docker for persistant data storage. 

---

## Endpoints
<details><summary>POST /createCar</summary>
<p>

#### Description:
Creates new car.

#### Parameters:
Content-Type: **application/json**
body *required

##### Model:
```json
{
    "model" : "string",
    "make": "string",
    "variant": "string"
}
```

##### Example Body:
```json
{
    "model" : "A45",
    "make": "Mercedes",
    "variant": "AMG"
}
```

##### Example Response Body:
```json
{
	"ID" : "0e03dda8-2c9a-4b19-958d-a96382587aee",  
    "model" : "A45",
    "make": "Mercedes",
    "variant": "AMG"
}
```
##### Responses:
200 OK
400 Bad Request
</p>
</details>




<details><summary>GET /cars</summary>
<p>

#### Description:
List all cars.

#### Parameters:
Content-Type: **application/json**

##### Example Response Body:
```json
[{
	"ID" : "0e03dda8-2c9a-4b19-958d-a96382587aee",  
    "model" : "A45",
    "make": "Mercedes",
    "variant": "AMG"
}
{
	"ID" : "63581320-83c4-4cdf-bee0-407b3579cb71",  
    "model" : "Model S",
    "make": "Tesla",
    "variant": "Sport"
}]
```
##### Responses:
200 OK
</p>
</details>




<details><summary>GET /cars/{carID}</summary>
<p>

#### Description:
Return one specific car.

#### Parameters:
Content-Type: **application/json**
path: {carID} *required
Note: carID has to be [RFC4122](https://tools.ietf.org/html/rfc4122) compliant.

##### Example Request:

`curl --location --request GET 'http://localhost:8080/cars/8ac3fc27-26a2-4fab-95d7-b6add7dcbfbf'`
##### Example Response:
```json
{
	"ID" : "63581320-83c4-4cdf-bee0-407b3579cb71",  
    "model" : "Model S",
    "make": "Tesla",
    "variant": "Sport"
}
```
##### Responses:
200 OK
400 Bad Request
404 Not Found
</p>
</details>




<details><summary>DELETE /cars/{carID}</summary>
<p>

#### Description:
Delete one specific car.

#### Parameters:
Content-Type: **application/json**
path: {carID} *required
Note: carID has to be [RFC4122](https://tools.ietf.org/html/rfc4122) compliant.

##### Example Request:

`curl --location --request DELETE 'http://localhost:8080/cars/8ac3fc27-26a2-4fab-95d7-b6add7dcbfbf'`

##### Responses:
204 No Content
400 Bad Request
404 Not Found
</p>
</details>




<details><summary>GET /search/{make}</summary>
<p>

#### Description:
Return all cars matching the given make.

#### Parameters:
Content-Type: **application/json**
path: {make} *required

##### Example Response Body:
```json
[{
	"ID" : "0e03dda8-2c9a-4b19-958d-a96382587aee",  
    "model" : "A45",
    "make": "Mercedes",
    "variant": "AMG"
}
{
	"ID" : "63581320-83c4-4cdf-bee0-407b3579cb71",  
    "model" : "B-Class",
    "make": "Mercedes",
    "variant": "Comfort"
}]
```
##### Responses:
200 OK
</p>
</details>

---
### Requirements
- Go v1.14.4
- Docker v19.03.12
- Docker-Compose v1.26.2
- PostgreQSL Docker-Image

---

### Run the car-management-api
To use persistant data store the PostgreSQL files location is needed. The path is working on MacOS & Linux, for Windows systems you may need to change the Postgres volume path within the docker-compose.yml. 

```
mdkir $HOME/docker/volumes/postgres
```

Copy PostgreSQL files into the directory above or create a new database. Use the migration.sql script to create the nessessary tables. 

Starting the API & PostgreSQL in Docker containers:
```
docker-compose up
```

Stop both containers:
```
docker-compose down
```

---

## Development Environment: 
For development purposes it is recommended to start the PostgreSQL Database manually using docker.

```
cd $GOPATH/src/car-api

docker run --name dev --env-file ./config/dbConfig env -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres 
```


### Access Database within Docker container: 
1. Open new bash window: `docker container ls`
2. `docker exec -it *POSTGRES_CONTAINER_ID* psql -U postgres -W postgres`
3. Run SQL Querys: Example `SELECT * FROM cars;`


The command `go run main.go` will start the Rest-API listening to port http://localhost:8080. The API will try to connect with the, make sure the Database is already running. 

Activating the "mockmode" will not connect to any database, all information will be stored in memory. The port the API will listen to can be set manually. Verbose logging is also supported, for more information and usage run:

 `go run main.go -h`

```
NAME:
   Car-Management-API - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   v0.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value   Port the Rest-API will listen on. (default: 8080)
   --mockmode     Set 'true' to use mocked mode. (default: API will use DB connection)
   --verbose      Set 'true' to enable verbose DEBUG-level logging. (default: Logging on WARN-level)
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
 ```