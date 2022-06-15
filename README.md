# go-server-example

HTTP server exmaple with Golang.

## Dependencies

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [mockery](https://github.com/vektra/mockery)
- [validator](https://github.com/go-playground/validator)
- [migrate](https://github.com/golang-migrate/migrate)

## Development

```sh
docker-compose up
```

## Migration

### Install

```sh
brew install golang-migrate
```

### Create migration files

```sh
migrate create -dir migrations -ext sql -seq create_todo_table 
```

### Migrate up

```sh
migrate -path migrations/ -database "mysql://user:password@tcp(127.0.0.1:3306)/intern" up  
```

## Testing

### Gernerate Mocks

#### Install Mockery

```sh
brew install mockery
```

#### Create Mocks

```sh
mockery --all  
```

### Unit Testing

```sh
go test -v ./...
```

### Load Testing

#### Create Todo

```sh
wrk -t 2 -c 10 -d 10s -s doc/wrk-script/create_todo.lua http://localhost:8080/api/todos
```

#### Get Todo

```sh
wrk -t 2 -c 10 -d 10s -s doc/wrk-script/get_todo.lua http://localhost:8080/api/todos/:id
```

#### Get todos

```sh
wrk -t 2 -c 10 -d 10s http://localhost:8080/api/todos
```
