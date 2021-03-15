# Hello World Api

How I setup this project:
```
go mod init hello-world-go-api
touch main.go
touch hello.go
go get github.com/go-chi/chi
go get github.com/go-chi/chi/middleware
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/alecthomas/template
go get github.com/joho/godotenv
```

Whenever you make a code-change, regenerate the swagger docs.  Can we automate this on file save?
```
swag init -g main.go
```

Run the api:
```
go run main.go hello.go
```


```
(for mac/linux) open http://localhost:8080/swagger/index.html
(for windows) start http://localhost:8080/swagger/index.html
curl -X GET localhost:8080/api/v1/hello
curl -X GET localhost:8080/api/v1/hello/mike
```


see 
 - https://github.com/go-chi/chi
 - https://github.com/swaggo/swag
 - https://github.com/swaggo/swag#declarative-comments-format
 - https://github.com/alecthomas/template
 - https://github.com/joho/godotenv