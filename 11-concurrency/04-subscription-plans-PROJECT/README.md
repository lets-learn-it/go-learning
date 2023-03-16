# Subscription

## Run

```sh
make build

make run

make stop
```

## Testing

```sh
# inside cmd/web folder
go test -v .

# check coverage
go test -coverprofile=coverage.out

# coverage in webbrowser
go tool cover -html=coverage.out
```
