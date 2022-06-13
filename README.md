# go-server-example

HTTP server exmaple with Golang.

---

## Dependencies

- [Echo](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [mockery](https://github.com/vektra/mockery)

## Migration

### Create migration files

```sh
migrate create -dir migrations -ext sql -seq create_todo_table 
```

### Migrate up

```sh
migrate -path migrations/ -database "mysql://user:password@tcp(127.0.0.1:3306)/intern" up  
```

## Testing

### Gernerate mocks

```sh
mockery --all  
```
