package models

import "github.com/nats-io/nats.go"

type HeaderField string

const (
	TimeStampHeaderField  HeaderField = "timestamp"
	RetryCountHeaderField HeaderField = "retry-count"
	MsgVersionHeaderField HeaderField = "msg-version"
	MsgFormatHeaderField  HeaderField = "msg-format"
	MsgTypeHeaderField    HeaderField = "msg-type"
	MsgSubtypeHeaderField HeaderField = "msg-subtype"
	TraceIdHeaderField    HeaderField = "trace-id"
	SpanIdHeaderField     HeaderField = "span-id"
	RequestIdHeaderField  HeaderField = "request-id"
	VesselIdHeaderField   HeaderField = "vessel-id"
)

type MsgType string

const (
	CommandMsgType MsgType = "command"
)

type MsgSubType string

const (
	StatusReqCmd MsgSubType = "status-req"
)

func NewNatsMessage(subject string, reply string, header nats.Header, data []byte) *nats.Msg {
	return &nats.Msg{
		Subject: subject,
		Reply:   reply,
		Header:  header,
		Data:    data,
	}
}
