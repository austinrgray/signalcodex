package models

import "time"

type Orbit struct {
	// Temporal and Positional Elements
	Epoch              time.Time
	Periapsis          float64
	Apoapsis           float64
	SemiMajorAxis      float64
	Eccentricity       float64
	Inclination        float64
	InclinationRate    float64
	LongAscendNode     float64
	ArgOfPeriapsis     float64
	ArgOfLatitude      float64
	MeanAnomaly        float64
	TrueAnomaly        float64
	MeanMotion         float64
	TimeSincePeriapsis float64

	// Energy and Momentum
	OrbitalEnergy       float64
	SpecificAngMomentum Vector3D

	// Perturbations
	J2Effect       float64
	SolarRadiation float64
	NBodyEffects   []*CelestialBody

	// Period
	OrbitalPeriod float64
}

type FlightParameters struct {
	// Trajectory and Maneuver Details
	FlightPathAngle          float64
	HyperbolicExcessVelocity Vector3D
	DeltaVRequired           float64
	DeltaVAvailable          float64
	ThrustMagnitude          float64
	BurnDuration             float64
	BurnDirection            Vector3D
	TargetCoordinates        Coordinate3D
	RelativeVelocity         Vector3D
	TrajectoryType           string
	PeriapsisDistance        float64
	ApoapsisDistance         float64
	OrbitalInclinationChange float64

	// Orbital Intersection and Targeting
	TargetBody           *CelestialBody
	IntersectionTime     float64
	IntersectionPosition Coordinate3D
	IntersectionVelocity Vector3D

	// Propulsion System Status
	//FuelRemaining            float64
	SpecificImpulse float64
	//ThrustEfficiency         float64

	// Environmental Factors
	SolarRadiationPressure    float64
	GravitationalPerturbation Vector3D
	AtmosphericDrag           float64

	// Timing
	ManeuverStartTime  time.Time
	ManeuverEndTime    time.Time
	CurrentMissionTime float64
}

type Telemetry struct {
	Attitude        Attitude3D
	Coordinate      Coordinate3D
	Velocity        Velocity3D
	Thrust          ThrustVector
	Mass            float64
	AngularMomentum AngularMomentum3D
	LinearMomentum  Momentum3D
	Orbit           Orbit
	Bodies          []*CelestialBody
}

type Attitude3D struct {
	Pitch float64
	Yaw   float64
	Rolls float64
}

type Coordinate3D struct {
	X float64
	Y float64
	Z float64
}

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

type ThrustVector struct {
	Force              float64
	Magnitude          float64
	Direction          Vector3D
	PointOfApplication Coordinate3D
}

type Torque struct {
	Force     Vector3D
	Axis      Vector3D
	Magnitude float64
}

type AngularVelocity3D struct {
	PitchRate float64
	YawRate   float64
	RollRate  float64
}

type Velocity3D struct {
	Linear  Vector3D
	Angular AngularVelocity3D
}

type AngularMomentum3D struct {
	LX float64
	LY float64
	LZ float64
}

type Momentum3D struct {
	PX float64
	PY float64
	PZ float64
}
