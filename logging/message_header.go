package logging

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func ZapMessageHeader(header nats.Header) []zap.Field {
	return []zap.Field{}
}
