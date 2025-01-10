package models

type Telemetry struct {
	Coordinate      Coordinate3D
	Attitude        Attitude3D
	Velocity        Velocity3D
	Thrust          ThrustVector
	LinearMomentum  Momentum3D
	AngularMomentum AngularMomentum3D
	Forces          []*GravitationalBody
}

type GravitationalBody struct {
	Name                  string
	Coordinate            Coordinate3D
	Attitude              Attitude3D
	Mass                  float64
	GravitationalConstant float64
	OrbitalParameters     OrbitalParameters
	ForceVector           Vector3D
	Velocity              Velocity3D
}

type OrbitalParameters struct {
	SemiMajorAxis            float64
	Eccentricity             float64
	Inclination              float64
	LongitudeOfAscendingNode float64
	ArgumentOfPeriapsis      float64
	TrueAnomoly              float64
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
	Magnitude float64
	Direction Vector3D
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
