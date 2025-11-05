package data

import (
	"mta_tracker/internal/model"
	"os"
	"strings"
)

// Currently live feeds TripId might not always match so this function is not being
// used right now.
func LoadTrips() (map[string]model.TripInfo, error) {
	data, err := os.ReadFile("assets/trips.txt")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	tripMap := make(map[string]model.TripInfo)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) >= 4 {
			routeID := strings.TrimSpace(parts[0])
			tripID := strings.TrimSpace(parts[1])
			tripName := strings.TrimSpace(parts[3])

			tripMap[tripID] = model.TripInfo{
				RouteID:  routeID,
				TripName: tripName,
				ID:       tripID,
			}
		}
	}

	return tripMap, nil
}

func GetTrip(tripMap map[string]model.TripInfo, tripID string) model.TripInfo {
	tripID = strings.TrimSpace(tripID)
	if trip, ok := tripMap[tripID]; ok {
		return trip
	}
	return model.TripInfo{}
}
