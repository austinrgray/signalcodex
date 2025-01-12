package models

import "time"

type OrbitalManeuver interface {
	Execute()
	Recover()
}

type ProgradeBurn struct {
	Id           string
	Start        time.Time
	BurnDuration float64
}

/*
A two-burn maneuver used to transfer between two coplanar circular orbits.
Use Case: Moving between two orbits of different radii, e.g., from Earth’s orbit to Mars’ orbit
*/
type HohmannTransfer struct {
	Id            string
	InitialOrbit  Orbit
	TargetOrbit   Orbit
	TransferOrbit Orbit
	PeriapsisBurn ProgradeBurn
	ApoapsisBurn  ProgradeBurn
}

/*
A three-burn maneuver, transferring through an intermediate orbit that extends further than the target orbit.
Use Case: More efficient than Hohmann transfer when the target orbit is significantly larger than the initial orbit.
*/
type BiEllipticTransfer struct {
}

/*
A single burn to adjust the inclination of an orbit.
Use Case: Aligning orbits with another celestial body or satellite, but costly in delta-v.
*/
type InclinationChange struct {
}

/*
A burn to adjust the orbital period, allowing synchronization with another object (e.g., docking or rendezvous).
Use Case: Timing an interception or meeting point.
*/
type PhasingManeuver struct {
}

/*
Adjusting an orbit to match another object's velocity and position.
Use Case: Docking with space stations or satellites.
*/
type CoOrbitalRendezvous struct {
}

/*
Final adjustments to relative velocity and position for docking.
Use Case: Precise approach in close proximity to a target.
*/
type TerminalPhaseManeuver struct {
}

/*
Using the gravitational pull of a celestial body to change trajectory and increase/decrease velocity.
Use Case: Boosting spacecraft to outer planets (e.g., Voyager missions) or redirecting to another body.
*/
type GravityAssist struct {
}

/*
A maneuver performed near a planet or celestial body where orbital velocity is highest, maximizing delta-v efficiency.
Use Case: Enhancing interplanetary or escape trajectories.
*/
type OberthEffect struct {
}

/*
A burn to leave a planet's sphere of influence and enter an interplanetary trajectory.
Use Case: Beginning missions to other planets or moons.
*/
type InterplanetaryInjection struct {
}

/*
Using a planet’s atmosphere to reduce velocity and achieve orbit or entry.
Use Case: Conserving delta-v by leveraging atmospheric drag.
*/
type AtmosphreicBraking struct {
}

/*
A repeated trajectory that cycles between planets, minimizing propellant use.
Use Case: Long-term missions with frequent transfers (e.g., Earth-Mars cyclers).
*/
type CyclerTrajectories struct {
}

/*
Small burns to maintain a spacecraft's position in a stable orbit.
Use Case: Ensuring satellite orbits remain geostationary or sun-synchronous.
*/
type StationKeeping struct {
}

/*
A burn to adjust an elliptical orbit into a circular one.
Use Case: Achieving stable orbits around celestial bodies.
*/
type OrbitCircularization struct {
}

/*
Burns to counteract atmospheric drag or gravitational perturbations.
Use Case: Extending satellite lifespans in low Earth orbit.
*/
type OrbitDecayCompensation struct {
}

/*
A burn to achieve escape velocity and leave a celestial body's sphere of influence.
Use Case: Exiting a planetary orbit or solar system.
*/
type EscapeTrajectory struct {
}

/*
Aligning the spacecraft for atmospheric entry, ensuring safe descent.
Use Case: Landing on Earth or another celestial body with an atmosphere.
*/
type ReEntryManeuver struct {
}

/*
Aligning the spacecraft for atmospheric entry, ensuring safe descent.
Use Case: Landing on Earth or another celestial body with an atmosphere.
*/
type PoweredDescent struct {
}

/*
A burn to establish orbit after reaching a target body.
Use Case: Inserting spacecraft into orbit around planets or moons.
*/
type OrbitInjection struct {
}

/*
Using low-thrust propulsion (e.g., ion engines) to spiral outward/inward to a target orbit.
Use Case: Efficient, long-term transfers.
*/
type LowThrustSpiralTransfer struct {
}

/*
A maneuver to position a spacecraft at one of the stable Lagrange points (e.g., L1, L2).
Use Case: Deploying space telescopes or observation satellites.
*/
type LagrangePointTransfer struct {
}

/*
Matching a Trojan object's orbit (leading or trailing 60° of a planet’s orbit).
Use Case: Studying Trojan asteroids (e.g., Lucy mission).
*/
type TrojanManeuver struct {
}

/*
A low-energy maneuver to slowly capture a spacecraft into orbit using gravity.
Use Case: Efficient orbital insertion around planets.
*/
type BallisticCapture struct {
}

/*
Adjusting trajectory to approach the Sun closely.
Use Case: Solar observation missions (e.g., Parker Solar Probe).
*/
type SunDiver struct {
}
