# go-server-example

HTTP server exmaple with Golang.

---

## Dependencies

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [mockery](https://github.com/vektra/mockery)
- [validator](https://github.com/go-playground/validator)
- [migrate](https://github.com/golang-migrate/migrate)

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

```sh
mockery --all  
```

### Unit Testing

```sh
go test -v ./...
```

### Load Testing

```
wrk -t 10 -c 100 -d 30s http://localhost:8080/api/todos/1
```
