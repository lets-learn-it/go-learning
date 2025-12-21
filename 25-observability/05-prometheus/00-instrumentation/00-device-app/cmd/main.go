package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pskp-95/device-app/internal/api"
)

type metrics struct {
	devices prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "deviceapp",
			Name:      "connected_devices",
			Help:      "number of currently connected devices.",
		}, []string{"version"}),
	}

	reg.MustRegister(m.devices)
	return m
}

func main() {
	// prometheus registry
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	// devices
	cDevices := api.NewConnectedDevices(&m.devices)

	dMux := http.NewServeMux()
	dMux.HandleFunc("GET /devices", cDevices.GetDevices)
	dMux.HandleFunc("POST /devices", cDevices.CreateDevice)

	pMux := http.NewServeMux()

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", dMux))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	select {}
}
