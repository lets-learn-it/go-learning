package api

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

type ConnectedDevices struct {
	devices  []Device
	guageVec *prometheus.GaugeVec
}

func NewConnectedDevices(gv *prometheus.GaugeVec) *ConnectedDevices {
	cds := &ConnectedDevices{
		devices: []Device{
			{1, "5F-28-CC-D6-F5-34", "2.1.6"},
			{2, "5F-29-CC-D6-F6-36", "2.1.5"},
			{3, "5F-34-DE-D6-FF-44", "2.1.6"},
		},
		guageVec: gv,
	}

	cMap := make(map[string]int)

	for _, device := range cds.devices {
		cMap[device.Firmware] += 1
	}

	for version, count := range cMap {
		cds.guageVec.With(prometheus.Labels{"version": version}).Set(float64(count))
	}

	return cds
}

func (cd *ConnectedDevices) GetDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(cd.devices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (cd *ConnectedDevices) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var dv Device

	err := json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cd.devices = append(cd.devices, dv)
	cd.guageVec.With(prometheus.Labels{"version": dv.Firmware}).Inc()

	w.WriteHeader(http.StatusCreated)
}

func (cd *ConnectedDevices) Count() map[string]int {
	cMap := make(map[string]int)

	for _, device := range cd.devices {
		cMap[device.Firmware] += 1
	}

	return cMap
}
