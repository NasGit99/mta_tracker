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

	fmt.Println("Trip ID | Route ID | StopID | StopName | DirectionID | TrainStatus")

	for _, t := range trains {
		fmt.Printf("%s | %s  | %s | %s | %d | %s\n",
			t.Trip.ID,
			t.Trip.RouteID,
			data.GetDirection(t.StopID),
			t.StopName,
			t.Trip.DirectionID,
			t.Vehicle.CurrentStatus,
		)
	}
}
