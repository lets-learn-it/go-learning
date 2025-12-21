package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pskp-95/nginx-exporter/internal/exporter"
)

func main() {
	var (
		targetHost = flag.String("target.host", "localhost", "nginx address with basic_status page")
		targetPort = flag.Int("target.port", 8080, "nginx port with basic_status page")
		targetPath = flag.String("target.path", "/status", "nginx status path")
		promPort   = flag.Int("prom.port", 9150, "port to expose prometheus metrics")
		logPath    = flag.String("target.log", "./nginx/access.log", "path to access.log file")
	)
	flag.Parse()

	log.Println(*targetHost)
	log.Println(*logPath)

	uri := fmt.Sprintf("http://%s:%d%s", *targetHost, *targetPort, *targetPath)

	basicStats := func() ([]exporter.NginxStats, error) {
		var netClient = &http.Client{
			Timeout: 10 * time.Second,
		}

		resp, err := netClient.Get(uri)
		if err != nil {
			log.Fatalf("netClinet.Get failed %s: %s", uri, err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("io.ReadAll failed: %s", err)
		}

		r := bytes.NewReader(body)

		return exporter.ScanBasicStats(r)
	}

	bc := exporter.NewBasicCollector(basicStats)

	reg := prometheus.NewRegistry()
	reg.MustRegister(bc)

	m := exporter.NewMetrics(reg)
	go m.TailAccessLogs(*logPath)

	mux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promHandler)

	port := fmt.Sprintf(":%d", *promPort)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("failed to start nginx exporter: %v", err)
	}
}
