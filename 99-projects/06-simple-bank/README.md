# Simple Bank

## migrate

- Install https://github.com/golang-migrate/migrate
- run command to create migration script
```
migrate create -ext sql -dir db/migration -seq init_schema
```
- apply migration script
```
migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## Generate Go code using sqlc

- Install sqlc `sudo snap install sqlc`
- Create queries
- `sqlc generate`

## mockgen

- Install mockgen `go install github.com/golang/mock/mockgen@v1.6.0`
- Create mock store `mockgen -destination db/mock/store.go -package mockdb github.com/simplebank/db/sqlc Store`

## References

