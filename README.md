how to install golang-migrate:
```
go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

how to create migration:
```
migrate create -ext sql -dir db/migrations create_table_name
```

how to migrate:
```
migrate -path db/migrations -database "postgres://username:password@host:5432/db_name?sslmode=disable" up
```
