# Port forwarder

- Start port forwarder `go run ./cmd/`. By default, it forwards from 8080 to 8081
- create tcp server `socat tcp4-listen:8081 -`
- create connection with port forwarder `socat tcp4-connect:127.0.0.1:8080 -`
