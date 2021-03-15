# Hello World Api


## Windows Setup
I've made use of Makefile.  You can either manually run the commands that are defined inthe Makefile, or install make via chocolatey:
 - open powershell as Administrator
 - install [chocolatey](https://chocolatey.org/install) with: `\Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))`
 - restart powershell
 - choco install make

## Run the API
This will update the swagger docs, and start the API.
```
make run
```

## Swagger
Swagger front-end is here http://localhost:8080/swagger/index.html

## Hit the API
For windows, install gitbash to get curl.  Or just use postman, or the swagger front-end.
```
curl -X GET localhost:8080/api/v1/hello
curl -X GET localhost:8080/api/v1/hello/mike
```

## DIY Project Setup
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

## See:
 - https://github.com/go-chi/chi
 - https://github.com/swaggo/swag
 - https://github.com/swaggo/swag#declarative-comments-format
 - https://github.com/alecthomas/template
 - https://github.com/joho/godotenv