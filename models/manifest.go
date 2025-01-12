package models

import (
	"fmt"
	"time"

	"github.com/austinrgray/signalcodex/messages"

	pb "google.golang.org/protobuf/proto"
	timepb "google.golang.org/protobuf/types/known/timestamppb"
)

type ManifestStatus int

const (
	ManifestAssigned ManifestStatus = iota
	ManifestPendingPickup
	ManifestInTransit
	ManifestDelivered
	ManifestDamaged
	ManifestLost
	ManifestCancelled
	ManifestReturning
	ManifestReturned
	ManifestArchived
)

type ConsignmentStatus int

const (
	ConsignmentAssigned ConsignmentStatus = iota
	ConsignmentPendingPickup
	ConsignmentInTransit
	ConsignmentDelivered
	ConsignmentDamaged
	ConsignmentLost
	ConsignmentCancelled
	ConsignmentReturning
	ConsignmentReturned
	ConsignmentArchived
)

type Manifest struct {
	Id            string
	Status        ManifestStatus
	SoulsOnboard  uint32
	Provisions    int32
	VesselMass    float64
	Consignments  []Consignment
	TimeAssigned  time.Time
	TimeCompleted time.Time
}

type Consignment struct {
	Id            string
	Status        ConsignmentStatus
	PickupId      string
	DestinationId string
	Items         []Item
	TimeLoaded    time.Time
	TimeDelivered time.Time
}

type Item struct {
	Description string
	Quantity    uint32
	Amount      float64
	Mass        float64
}

func NewManifest(id string, status ManifestStatus, souls uint32, provisions int32, mass float64, tAssigned time.Time) Manifest {
	return Manifest{
		Id:           id,
		Status:       status,
		SoulsOnboard: souls,
		Provisions:   provisions,
		VesselMass:   mass,
		TimeAssigned: tAssigned,
	}
}

func NewConsignment(id string, status ConsignmentStatus, pickupId string, destinationId string, items []Item) Consignment {
	return Consignment{
		Id:            id,
		Status:        status,
		PickupId:      pickupId,
		DestinationId: destinationId,
		Items:         items,
	}
}

func NewItem(description string, quantity uint32, amount float64, mass float64) Item {
	return Item{
		Description: description,
		Quantity:    quantity,
		Amount:      amount,
		Mass:        mass,
	}
}

func (m *Manifest) ToProto() *messages.Manifest {
	consignments := make([]*messages.Consignment, len(m.Consignments))
	for i, c := range m.Consignments {
		consignments[i] = c.ToProto()
	}

	var timeAssigned *timepb.Timestamp
	if !m.TimeAssigned.IsZero() {
		timeAssigned = timepb.New(m.TimeAssigned)
	}

	var timeCompleted *timepb.Timestamp
	if !m.TimeCompleted.IsZero() {
		timeCompleted = timepb.New(m.TimeCompleted)
	}

	return &messages.Manifest{
		Timestamp:     timepb.Now(),
		TimeAssigned:  timeAssigned,
		TimeCompleted: timeCompleted,
		Id:            m.Id,
		Status:        messages.ManifestStatus(m.Status),
		SoulsOnboard:  int32(m.SoulsOnboard),
		Provisions:    int32(m.Provisions),
		VesselMass:    m.VesselMass,
		Consignments:  consignments,
	}
}

func (c *Consignment) ToProto() *messages.Consignment {
	items := make([]*messages.Item, len(c.Items))
	for i, item := range c.Items {
		items[i] = item.ToProto()
	}

	var timeLoaded *timepb.Timestamp
	if !c.TimeLoaded.IsZero() {
		timeLoaded = timepb.New(c.TimeLoaded)
	}

	var timeDelivered *timepb.Timestamp
	if !c.TimeDelivered.IsZero() {
		timeDelivered = timepb.New(c.TimeDelivered)
	}

	return &messages.Consignment{
		TimeLoaded:    timeLoaded,
		TimeDelivered: timeDelivered,
		Id:            c.Id,
		Status:        messages.ConsignmentStatus(c.Status),
		PickupId:      c.PickupId,
		DestinationId: c.DestinationId,
		Items:         items,
	}
}

func (i *Item) ToProto() *messages.Item {
	return &messages.Item{
		Description: i.Description,
		Quantity:    int32(i.Quantity),
		Amount:      i.Amount,
		Mass:        i.Mass,
	}
}

func ManifestFromProto(msg *messages.Manifest) Manifest {
	consignments := make([]Consignment, len(msg.Consignments))
	for i, c := range msg.GetConsignments() {
		consignments[i] = ConsignmentFromProto(c)
	}

	var souls uint32
	if msg.SoulsOnboard > 0 {
		souls = uint32(msg.SoulsOnboard)
	} else {
		souls = 0
	}

	return Manifest{
		TimeAssigned:  msg.TimeAssigned.AsTime(),
		TimeCompleted: msg.TimeCompleted.AsTime(),
		Id:            msg.Id,
		Status:        ManifestStatus(msg.Status),
		SoulsOnboard:  souls,
		Provisions:    msg.Provisions,
		VesselMass:    msg.VesselMass,
		Consignments:  consignments,
	}
}

func ConsignmentFromProto(msg *messages.Consignment) Consignment {
	items := make([]Item, len(msg.Items))
	for i, item := range msg.Items {
		items[i] = ItemFromProto(item)
	}

	return Consignment{
		TimeLoaded:    msg.TimeLoaded.AsTime(),
		TimeDelivered: msg.TimeDelivered.AsTime(),
		Id:            msg.Id,
		Status:        ConsignmentStatus(msg.Status),
		PickupId:      msg.PickupId,
		DestinationId: msg.DestinationId,
		Items:         items,
	}
}

func ItemFromProto(i *messages.Item) Item {
	var qty uint32
	if i.Quantity > 0 {
		qty = uint32(i.Quantity)
	} else {
		qty = 0
	}

	return Item{
		Description: i.Description,
		Quantity:    qty,
		Amount:      i.Amount,
		Mass:        i.Mass,
	}
}

func MarshalManifest(msg *messages.Manifest) ([]byte, error) {
	bytes, err := pb.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal manifest message: %+v", err)
	}
	return bytes, nil
}

func UnmarshalManifest(bytes []byte) (*messages.Manifest, error) {
	var msg messages.Manifest
	err := pb.Unmarshal(bytes, &msg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bytes to manifest message: %+v", err)
	}
	return &msg, nil
}
