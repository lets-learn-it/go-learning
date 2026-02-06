FROM golang:1.25-alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -o /nginx-exporter ./cmd/exporter/

FROM alpine:3.22

COPY --from=builder /nginx-exporter /nginx-exporter

ENTRYPOINT ["/nginx-exporter"]
