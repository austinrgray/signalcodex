package models

import timepb "google.golang.org/protobuf/types/known/timestamppb"

type Manifest struct {
	Id           string
	Status       string
	Souls        uint8
	Provisions   int8
	VesselMass   float64
	Consignments []Consignment
}

type Consignment struct {
	Id                    string
	PickupWaypointId      string
	DestinationWaypointId string
	Items                 []Item
	PickedupTime          timepb.Timestamp
	DeliveredTime         timepb.Timestamp
}

type Item struct {
	Description string
	Qty         int
	Amount      float64
	Mass        float64
}
