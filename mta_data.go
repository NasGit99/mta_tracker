package main

import (
	"io"
	"log"
	"net/http"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

type TripInfo struct {
	ID          string
	RouteID     string
	DirectionID uint32
	StartTime   string
	StartDate   string
}

type VehicleInfo struct {
	ID             string
	Label          string
	Occupancy      uint32
	Congestion     string
	CurrentStatus  string
	CurrentStopSeq uint32
}

type Position struct {
	Latitude  float64
	Longitude float64
	Bearing   float32
	Speed     float32
	Timestamp uint64
}

type TrainData struct {
	Trip     TripInfo
	Vehicle  VehicleInfo
	Position Position
	StopID   string
	StopName string
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

	var trains []TrainData

	stopMap, err := LoadStops()
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		if entity.Vehicle != nil {
			tripData := entity.Vehicle.Trip
			vehicleData := entity.Vehicle.Vehicle
			positionData := entity.Vehicle.Position

			new_train := TrainData{
				Trip: TripInfo{
					ID:          tripData.GetTripId(),
					RouteID:     tripData.GetRouteId(),
					DirectionID: tripData.GetDirectionId(),
					StartTime:   tripData.GetStartTime(),
					StartDate:   tripData.GetStartDate(),
				},
				Vehicle: VehicleInfo{
					ID:             vehicleData.GetId(),
					Label:          vehicleData.GetLabel(),
					Occupancy:      entity.Vehicle.GetOccupancyPercentage(),
					Congestion:     entity.Vehicle.GetCongestionLevel().String(),
					CurrentStatus:  entity.Vehicle.GetCurrentStatus().String(),
					CurrentStopSeq: entity.Vehicle.GetCurrentStopSequence(),
				},
				Position: Position{
					Latitude:  float64(positionData.GetLatitude()),
					Longitude: float64(positionData.GetLongitude()),
					Bearing:   positionData.GetBearing(),
					Speed:     positionData.GetSpeed(),
					Timestamp: entity.Vehicle.GetTimestamp(),
				},
				StopID:   *entity.Vehicle.StopId,
				StopName: GetStopName(stopMap, *entity.Vehicle.StopId),
			}

			trains = append(trains, new_train)
		}
	}

	log.Printf("Loaded %d trips", len(trains))
	for _, t := range trains {
		log.Printf("%+v\n", t)
	}
}
