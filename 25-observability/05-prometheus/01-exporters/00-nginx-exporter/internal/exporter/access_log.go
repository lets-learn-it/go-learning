package exporter

import (
	"log"
	"regexp"
	"strconv"

	"github.com/hpcloud/tail"
	"github.com/prometheus/client_golang/prometheus"
)

var exp = regexp.MustCompile(`^(?P<ip>\S+) \S+ \S+ \[(?P<time>[^\]]+)\] "(?P<method>\S+) (?P<path>\S+) (?P<protocol>[^"]+)" (?P<status_code>\d{3}) (?P<body_bytes>\d+) "(?P<referer>[^"]*)" "(?P<user_agent>[^"]*)" (?P<request_time>[\d\.]+)$`)

type Metrics struct {
	size     prometheus.Counter
	duration *prometheus.HistogramVec
	requests *prometheus.CounterVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		size: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "nginx",
			Name:      "size_bytes_total",
			Help:      "Total bytes sent to clients.",
		}),
		requests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "nginx",
			Name:      "http_requests_total",
			Help:      "Total number of requests.",
		}, []string{"status_code", "method", "path"}),
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "nginx",
			Name:      "http_requests_duration_seconds",
			Help:      "Duration of the requests.",
			// optionally configure buckets
			// Buckets: prometheus.LinearBuckets(0.01, 0.05, 20),
			Buckets: prometheus.DefBuckets,
		}, []string{"status_code", "method", "path"}),
	}

	reg.MustRegister(m.size, m.duration, m.requests)

	return m
}

func (m *Metrics) TailAccessLogs(path string) {
	t, err := tail.TailFile(path, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		log.Fatalf("tail.TailFile failed: %s", err)
	}

	for line := range t.Lines {
		match := exp.FindStringSubmatch(line.Text)
		result := make(map[string]string)

		for i, name := range exp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}

		s, err := strconv.ParseFloat(result["body_bytes"], 64)
		if err != nil {
			continue
		}

		m.size.Add(s)

		m.requests.With(prometheus.Labels{
			"method":      result["method"],
			"status_code": result["status_code"],
			"path":        result["path"],
		}).Add(1)

		u, err := strconv.ParseFloat(result["request_time"], 64)
		if err != nil {
			continue
		}

		m.duration.With(prometheus.Labels{
			"method":      result["method"],
			"status_code": result["status_code"],
			"path":        result["path"],
		}).Observe(u)
	}
}
