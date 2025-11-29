# Notes API (Go + Chi + Clean Architecture)

A simple RESTful Notes API built using:
- Go 1.25+
- Chi Router
- Clean Architecture
- In-Memory Repository
- UUID-based IDs
---

## Running Locally

Run the API:
```sh
go run ./cmd/main
```
Server starts at:
```
http://localhost:8080
```
---
## Running with Docker
Build the Docker image:
```sh
docker build -t notes-api .
```
Run the container:
```sh
docker run -p 8080:8080 notes-api
```
---
## cURL to access the Api
Below are all cURL commands for testing Notes API.
---

## 1. Create a Note  
**POST /notes**
```sh
curl -X POST http://localhost:8080/notes   -H "Content-Type: application/json"   -d '{
    "title": "My First Note",
    "content": "Hello World"
  }'
```
---
## 2. Get All Notes  
**GET /notes**
```sh
curl http://localhost:8080/notes
```
---
## 3. Get Note by ID  
**GET /notes/{id}**
Example:
```sh
curl http://localhost:8080/notes/<uuid>
```
---
## 4. Update a Note  
**PUT /notes/{id}**
```sh
curl -X PUT http://localhost:8080/notes/<uuid>   -H "Content-Type: application/json"   -d '{
    "title": "Updated Title",
    "content": "Updated Content"
  }'
```
---

## 5. Delete a Note  
**DELETE /notes/{id}**
```sh
curl -X DELETE http://localhost:8080/notes/<uuid>
```
Returns:
```
204 No Content
```
---
## 6. Invalid UUID Example
```sh
curl http://localhost:8080/notes/123
```
Response:
```
400 Bad Request
invalid uuid
```
---
## 7. Note Not Found Example
```sh
curl http://localhost:8080/notes/00000000-0000-0000-0000-000000000000
```
Response:
```
404 Not Found
note not found
```