package data

import (
	"mta_tracker/internal/model"
	"os"
	"strings"
)

func LoadRoutes() (map[string]model.Routes, error) {
	data, err := os.ReadFile("assets/routes.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	routeMap := make(map[string]model.Routes)

	trips, err := LoadTrips()
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) >= 4 {
			id := strings.TrimSpace(parts[1])
			route := strings.TrimSpace(parts[3])
			direction := GetDirection(id)

			tripName := ""
			if trip, ok := trips[id]; ok {
				tripName = trip.TripName
			}

			routeMap[id] = model.Routes{
				RouteName: route,
				Direction: direction,
				TripName:  tripName,
				TripID:    id,
			}
		}
	}

	return routeMap, nil
}
