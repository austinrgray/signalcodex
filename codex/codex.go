package codex

type SvcName string

const (
	SvcCartographer   SvcName = "cartographer"
	SvcDispatchWeb    SvcName = "dispatch-terminal-webserver"
	SvcMissionControl SvcName = "mission-control"
	SvcRelayEmitter   SvcName = "relay-emitter"
	SvcRelayReceiver  SvcName = "relay-receiver"
)

type SubName string

const (
	SubRelayEmitterAll              SubName = SubName(SvcRelayEmitter + ".>")
	SubRelayEmitterFleet            SubName = SubName(SvcRelayEmitter + ".*")
	SubRelayEmitterFleetCommand     SubName = SubName(SvcRelayEmitter + ".*.command")
	SubRelayEmitterVesselFmtCommand SubName = SubName(SvcRelayEmitter + ".%s.command")
)

type StreamName string

const (
	StreamRelayEmitter StreamName = StreamName(SvcRelayEmitter + "-stream")
)

type MicroName string

const (
	MicroRelayEmitter MicroName = MicroName(SvcRelayEmitter + "-service")
)

type EndPointName string

const (
	EPRelayEmitCommand EndPointName = EndPointName("queue-emit-command")
)
