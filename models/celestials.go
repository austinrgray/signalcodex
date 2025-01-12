package models

import "time"

type CelestialBody interface {
	UpdatePosition() error
}

type Rotation struct {
	Axis          Vector3D
	SpinClockwise bool
	SpinSpeed     float64
	Period        float64
	LastUpdate    time.Time
}

type Planet struct {
	Id         string
	Name       string
	Position   Coordinate3D
	Rotation   Rotation
	Mass       float64
	Radius     float64
	Orbit      Orbit
	Atmosphere *Atmosphere
	Moons      []*Moon
}

type Moon struct {
	Id           string
	Name         string
	Position     Coordinate3D
	Rotation     Rotation
	Mass         float64
	Radius       float64
	Orbit        Orbit
	Atmosphere   *Atmosphere
	ParentPlanet *Planet
}

type Asteroid struct {
	Id       string
	Name     string
	Position Coordinate3D
	Rotation Rotation
	Mass     float64
	Radius   float64
	Orbit    Orbit
}

type Comet struct {
	Id                string
	Name              string
	Position          Coordinate3D
	Rotation          Rotation
	Mass              float64
	Radius            float64
	Orbit             Orbit
	Perihelion        float64  // Closest distance to the sun
	Apohelion         float64  // Farthest distance from the sun
	TailLength        float64  // The length of the comet's tail
	VolatileMaterials []string // Materials vaporized near perihelion
}

type SpaceStation struct {
	Id       string
	Name     string
	Position Coordinate3D
	Rotation Rotation
	Mass     float64
	Radius   float64
	Orbit    Orbit
}

type GravitationalBody interface {
	CelestialBody
	GetGravitationalInfluence() float64
}

type Atmosphere struct {
	Composition []string // e.g., Oxygen, Nitrogen, Carbon Dioxide, etc.
	Pressure    float64  // Atmospheric pressure in Pascals
	Temperature float64  // Average surface temperature in Kelvin
}
