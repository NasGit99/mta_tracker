package data

import (
	"os"
	"strings"
)

func LoadStops() (map[string]string, error) {
	data, err := os.ReadFile("assets/stops.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	stopMap := make(map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) >= 2 {
			id := strings.TrimSpace(parts[0])
			name := strings.TrimSpace(parts[1])
			stopMap[id] = name
		}
	}

	return stopMap, nil
}

func GetStopName(stopMap map[string]string, stopID string) string {
	stopID = strings.TrimSpace(stopID)
	if name, ok := stopMap[stopID]; ok {
		return name
	}
	return ""
}
