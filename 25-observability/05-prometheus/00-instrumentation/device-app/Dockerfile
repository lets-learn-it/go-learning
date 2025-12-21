FROM golang:1.25-alpine as builder

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -o /device-app ./cmd/

FROM alpine:3.22

COPY --from=builder /device-app /device-app

ENTRYPOINT ["/device-app"]
