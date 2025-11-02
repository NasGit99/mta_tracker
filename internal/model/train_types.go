package model

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
