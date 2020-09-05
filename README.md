# Car-management-api
This Rest-API provides basic functionality to handle car related data. The API is written in Go and is using PostgreSQL within Docker for persistant data storage.

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