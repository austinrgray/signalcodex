package models

import (
	"time"
)

type Mission struct {
	Id                      string
	Status                  string
	Currenttelemetry        Telemetry
	Waypoint                []Waypoint
	Directives              []Directive
	Consignments            []Consignment
	TimeAssigned            time.Time
	ScheduledStartTime      time.Time
	ActualStartTime         time.Time
	ScheduledCompletionTime time.Time
	ActualCompletionTime    time.Time
}

type Directive struct {
	Id              string
	DeltaT          uint64
	Status          string
	Action          string
	TargetTelemetry Telemetry
	Waypoint        *Waypoint
}

type Waypoint struct {
	ArrivalWindow   time.Time
	DepartureWindow time.Time
}
