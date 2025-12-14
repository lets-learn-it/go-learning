# Benchmark Testing

## Execute micro benchmarks with memory allocations

```shell
go test -bench=. ./...  -cpu=1,2,4,8 -benchmem
```

## CPU profiling

```shell
go test  -bench=. ./utils -cpuprofile=<profile_file>  -cpu=1,2,4,8 -benchmem
```

## Memory profiling

```shell
go test  -bench=. ./utils -memprofile=<profile_file>  -cpu=1,2,4,8 -benchmem
```

## Analysing profile data

```shell
go tool pprof <profile_file>
```

### pprof commands
- top10
- list <function_name>
- web
