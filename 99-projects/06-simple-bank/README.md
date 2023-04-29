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

## References

