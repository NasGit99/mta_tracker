package data

import (
	"strings"
)

func GetDirection(routeID string) string {

	r := routeID

	if strings.Contains(r, "N") {
		return "North"
	}

	if strings.Contains(r, "S") {
		return "South"
	}
	return "Unknown"
}
