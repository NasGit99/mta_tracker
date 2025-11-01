package main

import (
	"io"
	"log"
	"net/http"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

type Trip struct {
	DirectionID uint32
	RouteID     string
	StartTime   string
	StartDate   string
	StopID      string
	StopName    string
	VehicleID   string
}

func stopReader{
	
}

func main() {
	resp, err := http.Get("https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-ace")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}

	var trains []Trip

	for _, entity := range feed.Entity {
		if entity.Vehicle != nil {
			tr := entity.Vehicle.Trip
			newTrip := Trip{
				DirectionID: tr.GetDirectionId(),
				RouteID:     tr.GetRouteId(),
				StartTime:   tr.GetStartTime(),
				StartDate:   tr.GetStartDate(),
				StopID:      entity.Vehicle.GetStopId(),
				StopName:    "test_until_figure_out",
				VehicleID:   entity.Vehicle.Vehicle.GetId(),
			}
			trains = append(trains, newTrip)
		}
	}

	log.Printf("Loaded %d trips", len(trains))
	for _, t := range trains {
		log.Printf("%+v\n", t)
	}
}
