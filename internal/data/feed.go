package data

import (
	"io"
	"log"
	"net/http"

	"mta_tracker/internal/model"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	"google.golang.org/protobuf/proto"
)

func Train_data() ([]model.TrainData, error) {
	resp, err := http.Get("https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs")
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

	var trains []model.TrainData

	stopMap, err := LoadStops()
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		if entity.Vehicle != nil {
			tripData := entity.Vehicle.Trip
			vehicleData := entity.Vehicle.Vehicle
			positionData := entity.Vehicle.Position

			new_train := model.TrainData{
				Trip: model.TripInfo{
					ID:          tripData.GetTripId(),
					RouteID:     tripData.GetRouteId(),
					DirectionID: tripData.GetDirectionId(),
					StartTime:   tripData.GetStartTime(),
					StartDate:   tripData.GetStartDate(),
				},
				Vehicle: model.VehicleInfo{
					ID:             vehicleData.GetId(),
					Label:          vehicleData.GetLabel(),
					Occupancy:      entity.Vehicle.GetOccupancyPercentage(),
					Congestion:     entity.Vehicle.GetCongestionLevel().String(),
					CurrentStatus:  entity.Vehicle.GetCurrentStatus().String(),
					CurrentStopSeq: entity.Vehicle.GetCurrentStopSequence(),
				},
				Position: model.Position{
					Latitude:  float64(positionData.GetLatitude()),
					Longitude: float64(positionData.GetLongitude()),
					Bearing:   positionData.GetBearing(),
					Speed:     positionData.GetSpeed(),
					Timestamp: entity.Vehicle.GetTimestamp(),
				},
				StopID:   *entity.Vehicle.StopId,
				StopName: stopMap[*entity.Vehicle.StopId],
			}

			trains = append(trains, new_train)
		}
	}
	return trains, nil
}
