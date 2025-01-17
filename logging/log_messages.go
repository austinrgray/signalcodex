package logging

const (
	ReceivedTerminationSignalLogMsg      = "received termination signal, shutting down..."
	DrainingNATSConnectionLogMsg         = "draining NATS connection..."
	ServiceShutdownLogMsg                = "service shut down gracefully"
	NATSConnectSuccessLogMsg             = "connected to NATS server"
	NATSConnectFailureLogMsg             = "failed to connect to NATS server"
	NATSDrainingErrorLogMsg              = "error during NATS drain"
	JetstreamInitSuccessLogMsg           = "initialized Jetstream"
	JetstreamInitFailureLogMsg           = "failed to enable Jetstream"
	MicroServiceInitSuccessLogMsg        = "initialized micro service"
	MicroServiceInitFailureLogMsg        = "failed to initialize service"
	EndpointAddSuccessLogMsg             = "successfully added endpoint"
	EndpointAddFailureLogMsg             = "failed to add endpoint"
	EmitHandlerUnrecognizedMsgTypeLogMsg = "unrecognized msg type"
	EmitHandlerPublishFailedLogMsg       = "failed to publish message"
	EmitHandlerPublishSuccessLogMsg      = "message published successfully"
	EmitHandlerFallbackSpanIDLogMsg      = "error creating new span-id"
)
