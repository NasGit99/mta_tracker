package main

import (
	"fmt"
	"log"
	"mta_tracker/internal/data"
)

func main() {
	trains, err := data.Train_data()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d trips", len(trains))

	fmt.Println("Route Name | RouteID|  Trip Name | Direction | StopName | TrainStatus")

	routeMap, err := data.LoadRoutes()
	if err != nil {
		return
	}

	for _, t := range trains {
		route := data.ParseRoute(routeMap, t.Trip.RouteID)

		fmt.Printf("%s | %s | %s | %s | %s\n",
			route.RouteName,
			route.RouteID,
			data.GetDirection(t.StopID),
			t.StopName,
			t.Vehicle.CurrentStatus,
		)
	}
}
