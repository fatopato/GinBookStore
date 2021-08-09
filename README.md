# GinBookStore
Simple Rest API with Golang and Gin Framework.

The API uses sqlite3 so there will be a sqlite3 file created named as 'data.db' after the server started

You can get imports by running 
```
$ go get .
```
Then start the server by
```
$ go run .
```
The server starts on 8080 port and you can reach the the API via http://localhost:8080/api/v1/
You can do CRUD operations by interacting with the API

# Sample Requsts

Get All Books:
```
$ curl -i http://localhost:8080/api/v1/books
```

Get All Books:
```
$ curl -i http://localhost:8080/api/v1/books
```

Save Book:
```
$ curl -i -X POST -H "Content-Type: application/json" -d "{ \"title\": \"How to be worse at programming\", \"author\": \"Fatopato\" }" http://localhost:8080/api/v1/books
```

Get Book By Id (id: 1):

```
curl -i http://localhost:8080/api/v1/books/1
```

Update Book by Id (id:1)

```
curl -i -X PUT -H "Content-Type: application/json" -d "{ \"title\": \"It's getting even worse\", \"author\": \"Fatopato\" }" http://localhost:8080/api/v1/books/1
```

Delete Book by Id (id:1)
```
curl -i -X DELETE http://localhost:8080/api/v1/books/1
```
