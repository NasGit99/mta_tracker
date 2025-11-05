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

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) >= 4 {
			routeID := strings.TrimSpace(parts[1])
			route := strings.TrimSpace(parts[3])

			routeMap[routeID] = model.Routes{
				RouteName: route,
				RouteID:   routeID,
			}
		}
	}

	return routeMap, nil
}

func ParseRoute(routeMap map[string]model.Routes, routeID string) model.Routes {
	routeID = strings.TrimSpace(routeID)
	if route, ok := routeMap[routeID]; ok {
		return route
	}
	return model.Routes{}
}
