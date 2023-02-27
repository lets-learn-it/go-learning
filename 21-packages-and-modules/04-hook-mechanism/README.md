# Hook Mechanism


## build tags

**Old Method**

```go
// +build tag1
```

**New Method**

```go
//go:build tag1
```

While building, if you want to compile that file in executable

```sh
go build -tag tag1
```

