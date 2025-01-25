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
	SubRelayEmitterAll                 SubName = SubName(SvcRelayEmitter + ".>")
	SubRelayEmitterFleet               SubName = SubName(SvcRelayEmitter + ".*")
	SubRelayEmitterFleetCommand        SubName = SubName(SvcRelayEmitter + ".*.command")
	SubRelayEmitterVesselFmtCommand    SubName = SubName(SvcRelayEmitter + ".%s.command")
	SubRelayReceiverAll                SubName = SubName(SvcRelayReceiver + ".>")
	SubRelayReceiverFleet              SubName = SubName(SvcRelayReceiver + ".*")
	SubRelayReceiverFleetCommand       SubName = SubName(SvcRelayReceiver + ".*.status")
	SubRelayReceiverVesselFmtCommand   SubName = SubName(SvcRelayReceiver + ".%s.status")
	SubMissionControlAll               SubName = SubName(SvcMissionControl + ".>")
	SubMissionControlSentryFleetStatus SubName = SubName(SvcMissionControl + ".sentry.fleet.status")
)

type StreamName string

const (
	StreamRelayEmitter   StreamName = StreamName(SvcRelayEmitter + "-stream")
	StreamRelayReceiver  StreamName = StreamName(SvcRelayReceiver + "-stream")
	StreamMissionControl StreamName = StreamName(SvcMissionControl + "-stream")
)

type MicroName string

const (
	MicroRelayEmitter   MicroName = MicroName(SvcRelayEmitter + "-service")
	MicroMissionControl MicroName = MicroName(SvcMissionControl + "-service")
)

type EndPointName string

const (
	EPRelayEmitCommand        EndPointName = EndPointName("queue-emit-command")
	EPRelayEmitAddConsumer    EndPointName = EndPointName("add-consumer")
	EPRelayEmitRemoveConsumer EndPointName = EndPointName("remove-consumer")
)
