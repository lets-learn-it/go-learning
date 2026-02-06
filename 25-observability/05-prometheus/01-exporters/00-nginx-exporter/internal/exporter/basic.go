package exporter

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type NginxStats struct {
	ConnectionsActive float64

	// Connections (Reading - Writing - Waiting)
	Connections []Connections
}

type Connections struct {
	Type  string
	Total float64
}

func ScanBasicStats(r io.Reader) ([]NginxStats, error) {
	s := bufio.NewScanner(r)

	var stats []NginxStats
	var connections []Connections
	var nginxStats NginxStats

	for s.Scan() {
		fields := strings.Fields(string(s.Bytes()))
		fmt.Println(fields)

		if len(fields) == 3 && fields[0] == "Active" {
			c, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float %s: %w", fields[2], err)
			}

			nginxStats.ConnectionsActive = c
		}

		if fields[0] == "Reading:" {
			r, err := strconv.ParseFloat(fields[1], 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float %s: %w", fields[2], err)
			}

			reading := Connections{Type: "Reading", Total: r}

			w, err := strconv.ParseFloat(fields[3], 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float %s: %w", fields[2], err)
			}

			writing := Connections{Type: "Writing", Total: w}

			wa, err := strconv.ParseFloat(fields[5], 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float %s: %w", fields[2], err)
			}

			waiting := Connections{Type: "Waiting", Total: wa}

			connections = append(connections, reading, writing, waiting)
			nginxStats.Connections = connections
		}
	}

	stats = append(stats, nginxStats)

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("failed to read metrics: %w", err)
	}

	return stats, nil
}
