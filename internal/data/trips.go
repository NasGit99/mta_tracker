package data

import (
	"mta_tracker/internal/model"
	"os"
	"strings"
)

func LoadTrips() (map[string]model.TripRoute, error) {
	data, err := os.ReadFile("assets/trips.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	tripMap := make(map[string]model.TripRoute)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) >= 6 {
			routeID := strings.TrimSpace(parts[0])
			tripID := strings.TrimSpace(parts[1])
			name := strings.TrimSpace(parts[3])

			tripMap[routeID] = model.TripRoute{
				RouteID:  routeID,
				TripName: name,
				TripID:   tripID,
			}
		}
	}

	return tripMap, nil
}

func GetTripName(tripMap map[string]model.TripRoute, tripID string) string {
	tripID = strings.TrimSpace(tripID)
	if trip, ok := tripMap[tripID]; ok {
		return trip.TripName
	}
	return ""
}
