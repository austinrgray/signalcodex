package models

import (
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type Vessel struct {
	Info          VesselInfo
	Manifest      Manifest
	Mission       *Mission
	Bridge        *Bridge
	Comms         *Comms
	Engines       *Engines
	Habitat       *Habitat
	Nav           *Nav
	RemoteSensors *RemoteSensors
}

type VesselInfo struct {
	Id      string
	Name    string
	Version string
	ModelId string
	Class   string
}

type Bridge struct {
	Mode         string
	Health       string
	SystemAlerts []SystemAlert
}

type Comms struct {
	LocalAddr    string
	Mode         string
	Health       string
	Conns        map[string]*Connection
	SystemAlerts []SystemAlert
}

type Connection struct {
	RemoteAddr   string
	Mode         string
	Health       string
	Conn         *nats.Conn
	In           chan nats.Msg
	Out          chan proto.Message
	SystemAlerts []SystemAlert
}

type Engines struct {
	Mode            string
	Health          string
	AggregateThrust ThrustVector
	Thrusters       []*Thruster
	SystemAlerts    []SystemAlert
}

type Thruster struct {
	Mode         string
	Health       string
	ThrustVector ThrustVector
	SystemAlerts []SystemAlert
}

type Habitat struct {
	Mode         string
	Health       string
	O2           float32
	CO2          float32
	KPA          float32
	Temp         float32
	Humidity     float32
	AQI          float32
	SystemAlerts []SystemAlert
}

type Nav struct {
	Mode             string
	Health           string
	CurrentTelemetry Telemetry
	CurrentDirective Directive
	NextDirective    Directive
	SystemAlerts     []SystemAlert
}

type RemoteSensors struct {
	Mode         string
	Health       string
	Readings     []CelestialBody
	SystemAlerts []SystemAlert
}
