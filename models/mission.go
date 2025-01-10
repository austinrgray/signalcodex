package models

import timepb "google.golang.org/protobuf/types/known/timestamppb"

type Mission struct {
	Id                      string
	Status                  string
	Currenttelemetry        Telemetry
	Waypoint                []Waypoint
	Directives              []Directive
	Consignments            []Consignment
	TimeAssigned            timepb.Timestamp
	ScheduledStartTime      timepb.Timestamp
	ActualStartTime         timepb.Timestamp
	ScheduledCompletionTime timepb.Timestamp
	ActualCompletionTime    timepb.Timestamp
}

type Directive struct {
	Id              string
	DeltaT          uint64
	Status          string
	Action          string
	TargetTelemetry Telemetry
	Waypoint        *Waypoint
	NextDirective   *Directive
}

type Waypoint struct {
	Id                string
	Name              string
	Type              string
	Coordinate        Coordinate3D
	Attitude          Attitude3D
	Velocity          Velocity3D
	OrbitalParameters OrbitalParameters
	Forces            []GravitationalBody
	ArrivalWindow     timepb.Timestamp
	DepartureWindow   timepb.Timestamp
}
