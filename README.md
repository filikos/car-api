# Car-management-api
This Rest-API provides basic functionality to handle car related data. The API is written in Go and is using PostgreSQL within Docker for persistant data storage.

## Requirements
- Go v1.14.4
- Docker
- PostgreQSL Image

## Run Car-Management-API
The command `go run main.go` will start the Rest-API which with DB Connection. Ensure that the values set in `./config/dbConfig.env` are matching the PostgreSQL DB configuration. 

If you are using the "mockmode" no configuration is needed since all data will be stored temporary in memory.

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

#### Development Environment: 
For development purposes it is recommended to start the PostgreSQL container first and then the API since it depends on the DB connection. 
Note: 
- `cd` into the project path to ensure docker will find the dbConfig.env file.
- PostgreSQL persistant storage volume path: `$HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres`


1. `docker run --name dev --env-file ./config/dbConfig.env -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres`
2. `go run main.go`
3. *New Terminal* `docker container ls -a`
4. `docker exec -it *POSTGRES_CONTAINER_ID* psql -U postgres -W postgres`



## Ressources 
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
