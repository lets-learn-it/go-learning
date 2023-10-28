# Simple GRPC App

## Setup

### Install Protocol Buffer compiler

https://github.com/protocolbuffers/protobuf/releases

### Generate Code

```
protoc --go_out=. --go-grpc_out=. ./proto/greet.proto
```

## GRPC

![grpc](./images/grpc.png)

## References

[[00] FULL PROJECT - GO + GRPC - youtube.com](https://www.youtube.com/watch?v=a6G5-LUlFO4)
