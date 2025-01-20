package logging

type LogKey string
type LogMsg string
type LogFmt string
type LogErrMsg string
type LogErrFmt string

// Logger Log Messages
const (
	LogErrFmtDebugLoggerInit LogErrFmt = "error opening debug log file: %w"
	LogErrFmtErrorLoggerInit LogErrFmt = "error opening error log file: %w"
	LogErrMsgLoggerInit      LogErrMsg = "error initializing logger, using default instead"
)

// Load Config Log Messages
const (
	LogErrMsgConfigInit LogErrMsg = "error loading config, using default instead"
)

// Service Log Messages
const (
	LogErrMsgServiceOSTermSignal       LogErrMsg = "received termination signal, shutting down..."
	LogErrMsgServiceUnrecoverableError LogErrMsg = "returned unrecoverable error, shutting down..."
	LogErrMsgServiceInitRetryLimit     LogErrMsg = "initialization retry limit exceeded, shutting down..."
	LogMsgServiceShutdown              LogMsg    = "service shut down gracefully"
	LogMsgServicePaused                LogMsg    = "service paused..."
	LogMsgServiceStart                 LogMsg    = "starting service..."
)

// NATS Log Messages
const (
	LogErrMsgNATSInit                LogErrMsg = "failed to initialize connection assets"
	LogErrMsgNATSConnect             LogErrMsg = "failed to connect to NATS"
	LogErrMsgNATSDrainingConnection  LogErrMsg = "error during NATS drain"
	LogErrMsgNATSSubscription        LogErrMsg = "subscription error"
	LogErrMsgNATSConnection          LogErrMsg = "error with NATS connection"
	LogMsgNATSInit                   LogMsg    = "successfully initialized connection assets"
	LogMsgNATSConnect                LogMsg    = "successfully connected to NATS"
	LogMsgNATSDrainingConnection     LogMsg    = "draining NATS connection..."
	LogMsgNATSClosingConnection      LogMsg    = "closing connection to NATS"
	LogMsgNATSClosedServer           LogMsg    = "server closed"
	LogMsgNATSDisconnectedConnection LogMsg    = "disconnected from server"
	LogMsgNATSLameDuckServer         LogMsg    = "lame duck server"
)

// Jetstream Log Messages
const (
	LogErrMsgJetstreamInit LogErrMsg = "failed to enable Jetstream"
	LogMsgJetstreamInit    LogMsg    = "successfully initialized Jetstream"
)

// Micro Log Messages
const (
	LogErrMsgMicroServiceInit LogErrMsg = "failed to initialize service"
	LogErrMsgMicroAddEp       LogErrMsg = "failed to add endpoint"
	LogMsgMicroInit           LogMsg    = "initialized micro service"
	LogMsgMiscroAddEp         LogMsg    = "successfully added endpoint"
)

// Handler Log Messages
const (
	LogErrMsgHandlerMsgTypeParse LogErrMsg = "error parsing message type from header"
	LogErrMsgHandlerReadHeader   LogErrMsg = "error reading from header"
	LogErrMsgHandlerCreateHeader LogErrMsg = "error creating header"
	LogErrMsgHandlerPublish      LogErrMsg = "error publishing message"
	LogErrMsgHandlerMarshal      LogErrMsg = "error marshaling message"
	LogErrMsgHandlerUnmarshal    LogErrMsg = "error unmarshaling message"
	LogErrMsgResponse            LogErrMsg = "err"
	LogMsgHandlerPublish         LogMsg    = "published message"
	LogMsgHandlerMarshal         LogMsg    = "marshaled message"
	LogMsgHandlerUnmarshal       LogMsg    = "unmarshaled message"
	LogMsgResponse               LogMsg    = "ok"
)

// Log Keys
const (
	LogKeySvcName       LogKey = "service"
	LogKeySubName       LogKey = "subject"
	LogKeyStreamName    LogKey = "stream"
	LogKeyConsumerName  LogKey = "consumer"
	LogKeyMicroName     LogKey = "micro"
	LogKeyHandlerName   LogKey = "handler"
	LogKeyServerAddress LogKey = "server-address"
	LogKeyClientAddress LogKey = "client-address"
	LogKeyHeaderMsg     LogKey = "header"
	LogKeyPayloadMsg    LogKey = "payload"
	LogKeyAckSequence   LogKey = "ack-sequence"
	LogKeyCount         LogKey = "count"
	LogKeyTimeStamp     LogKey = "timestamp"
	LogKeySalience      LogKey = "salience"
)
