package exporter

import "github.com/prometheus/client_golang/prometheus"

var _ prometheus.Collector = &basicCollector{}

type basicCollector struct {
	ConnectionsActive *prometheus.Desc

	// connections (reading - writing - waiting)
	Connections *prometheus.Desc

	stats func() ([]NginxStats, error)
}

func NewBasicCollector(stats func() ([]NginxStats, error)) prometheus.Collector {
	return &basicCollector{
		ConnectionsActive: prometheus.NewDesc(
			"nginx_connections_active",
			"Number of active connections",
			[]string{},
			nil,
		),
		Connections: prometheus.NewDesc(
			"nginx_connections_total",
			"Number of connections (reading - waiting - writing)",
			[]string{"type"},
			nil,
		),
		stats: stats,
	}
}

// Collect implements [prometheus.Collector].
func (b *basicCollector) Collect(ch chan<- prometheus.Metric) {
	stats, err := b.stats()
	if err != nil {
		ch <- prometheus.NewInvalidMetric(b.ConnectionsActive, err)
		return
	}

	for _, s := range stats {
		ch <- prometheus.MustNewConstMetric(
			b.ConnectionsActive,
			prometheus.GaugeValue,
			s.ConnectionsActive,
		)

		for _, conn := range s.Connections {
			ch <- prometheus.MustNewConstMetric(
				b.Connections,
				prometheus.CounterValue,
				conn.Total,
				conn.Type,
			)
		}
	}
}

// Describe implements [prometheus.Collector].
func (b *basicCollector) Describe(ch chan<- *prometheus.Desc) {
	ds := []*prometheus.Desc{
		b.ConnectionsActive,
	}

	for _, d := range ds {
		ch <- d
	}
}
