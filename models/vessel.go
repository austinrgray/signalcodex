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
	Mode   string
	Health string
	Errors []Error
}

type Comms struct {
	LocalAddr string
	Mode      string
	Health    string
	Conns     map[string]*Connection
	Errors    []Error
}

type Connection struct {
	RemoteAddr string
	Mode       string
	Health     string
	Conn       *nats.Conn
	In         chan nats.Msg
	Out        chan proto.Message
	Errors     []Error
}

type Engines struct {
	Mode            string
	Health          string
	AggregateThrust ThrustVector
	Thrusters       []*Thruster
	Errors          []Error
}

type Thruster struct {
	Mode         string
	Health       string
	ThrustVector ThrustVector
	Errors       []Error
}

type Habitat struct {
	Mode     string
	Health   string
	O2       float32
	CO2      float32
	KPA      float32
	Temp     float32
	Humidity float32
	AQI      float32
	Errors   []Error
}

type Nav struct {
	Mode             string
	Health           string
	CurrentTelemetry Telemetry
	CurrentDirective Directive
	NextDirective    Directive
	Errors           []Error
}

type RemoteSensors struct {
	Mode     string
	Health   string
	Readings []CelestialBody
	Errors   []Error
}
