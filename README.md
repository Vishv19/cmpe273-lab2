*To Start Server*

go run webService.go

Server would start at "0.0.0.0:8080" or "http://localhost:8080/hello"

To POST request

POST http://localhost:8080/hello with following body data:

{
   “name”: “foo”
}

To GET resonse "GET http://localhost:8080/hello"

Response:

{
   “greeting” : “Hello, foo!”
}