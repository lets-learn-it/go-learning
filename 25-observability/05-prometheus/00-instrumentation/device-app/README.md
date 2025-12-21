# Device app

## Http actions

```sh
# create device
curl -X POST http://127.0.0.1:8080/devices -H 'Content-Type: application/json' -d '{"id": 4, "mac": "5F-28-CC-D7-F5-34", "firmware": "2.1.7"}'

# get devices
curl http://127.0.0.1:8080/devices
```

## References

[[00] How to Monitor/Instrument Golang with Prometheus (Counter - Gauge - Histogram - Summary) by Anton Putra on youtube.com](https://www.youtube.com/watch?v=WUBjlJzI2a0)
[[01] Instrumenting a Go application for Prometheus - prometheus.io](https://prometheus.io/docs/guides/go-application/)
