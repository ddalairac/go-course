# go-course

Repo for Udemy [Building Modern Web Applications with Go](https://www.udemy.com/course/building-modern-web-applications-with-go/) course

Basic commands

```go
//run a single file
go run main.go
// run app
go run .
```

To run the app a small config is needed first

```go
// init go module
go mod init go-course
// check that all dependencies proprly config
go mod tidy
```

## DB

recomended tools:

- [Postgres app](https://postgresapp.com/downloads.html)
- [DBeaver ](https://dbeaver.io/download/)

```go
//Install PostgreSQL driver
go get github.com/lib/pq
// Routing
go get github.com/gorilla/
```
