# Gin Otel

## Jaeger Setup

```
docker run --rm --name jaeger -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 4317:4317 -p 4318:4318 -p 14250:14250 -p 14268:14268 -p 14269:14269  -p 9411:9411 jaegertracing/all-in-one:1.53
```

## Environment Variables

```ps
$env:SERVICE_NAME="goApp"
$env:INSECURE_MODE="true"
$env:OTEL_EXPORTER_OTLP_ENDPOINT="localhost:4317"
```

## Reference

[[00] Complete guide to implementing OpenTelemetry in Go applications - signoz.io](https://signoz.io/opentelemetry/go/)
